package cmd

import (
	"context"
	"fmt"
	"path/filepath"
	strings "strings"
	"sync"
	"time"
	"unicode"

	// badger "github.com/dgraph-io/badger/v2"
	"github.com/DataDog/zstd"
	badger "github.com/dgraph-io/badger/v3"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/sap/gorfc/gorfc"
	log "github.com/sirupsen/logrus"
	"github.com/ulranh/saphistory/internal"
)

// sap server
type server struct {
	name string
	conn *gorfc.Connection
}

// sap system data
type systemData struct {
	sid     string
	sapInfo internal.D3StringList
}

// sap tcode data
type tcodeData struct {
	tcode     string
	tcodeInfo *internal.D2StringList
}

// transaction data
type transaction struct {
	name string
	head []string
	// headKey    string
	fumo       string
	params     map[string]interface{}
	filter     map[string][]interface{} // only exclude (strings starting with !) or include
	table      string
	fields     []string
	allServers bool
	dataFunc   func(transaction, server) []*internal.D1StringList
}

// save - save transaction-data of all systems to the badger db
func (env envInfo) save() error {
	systems, err := env.getSystemInfo()
	if err != nil {
		return errors.Wrap(err, " Save - GetSystemInfo")
	}

	// return nothing if there are no systems
	if 0 == len(systems) {
		log.Println("No systems available")
		return nil
	}

	// select the existing secret info
	secret, err := env.getSecret()
	if err != nil {
		return errors.Wrap(err, " :save - getSecret")
	}

	// get correct password for every system
	for _, system := range systems {
		system.Password, err = getPw(secret, system.Sid)
		if err != nil {
			log.WithFields(log.Fields{
				"system": system.Sid,
				"error":  err,
			}).Error("Cannot get secret information for system")
			continue
		}
	}

	err = env.collectSystemsTransactions(env.getTransactions(), systems)
	if err != nil {
		return errors.Wrap(err, " :save - collectSystemsTransactions")
	}
	return nil
}

// collectSystemsTransactions - collecting all transactions for all systems
func (env envInfo) collectSystemsTransactions(transactions []transaction, systems []*internal.SystemInfo) error {

	var wg sync.WaitGroup
	sDataC := make(chan systemData, len(systems))

	for _, s := range systems {

		wg.Add(1)
		go func(s *internal.SystemInfo) {
			defer wg.Done()
			sDataC <- systemData{
				sid:     s.Sid,
				sapInfo: env.collectSystemTransactions(transactions, s),
			}
		}(s)
	}

	go func() {
		wg.Wait()
		close(sDataC)
	}()

	for sysData := range sDataC {

		err := env.saveBadger(sysData)
		if err != nil {
			return errors.Wrap(err, " collectSystemsTransactions - saveBadger")
		}
	}
	return nil
}

// collectSystemTransaction - collecting all transactions for one system
func (env envInfo) collectSystemTransactions(transactions []transaction, s *internal.SystemInfo) internal.D3StringList {
	var wg sync.WaitGroup
	tsDataC := make(chan tcodeData, len(transactions))

	for _, t := range transactions {

		wg.Add(1)
		go func(t transaction) {
			defer wg.Done()
			tsDataC <- tcodeData{
				tcode:     t.name,
				tcodeInfo: env.collectSystemTransaction(t, s),
			}
		}(t)
	}

	go func() {
		wg.Wait()
		close(tsDataC)
	}()

	var tcData []tcodeData
	for v := range tsDataC {
		if v.tcodeInfo != nil {
			tcData = append(tcData, v)
		}
	}

	// sort in correct transaction order
	var tData []*internal.D2StringList
	for _, t1 := range transactions {
		for _, t2 := range tcData {
			if t1.name == t2.tcode {
				tData = append(tData, t2.tcodeInfo)
			}
		}
	}

	return internal.D3StringList{
		Data3: tData,
	}
}

// collectSystemTransaction - collecting one transaction for one system
func (env envInfo) collectSystemTransaction(t transaction, s *internal.SystemInfo) *internal.D2StringList {

	// get connection strings for all sap application servers
	servers := getServerConnections(s, t)
	if 0 == len(servers) {
		return nil
	}

	// prepare timeout
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(*env.timeout)*time.Second))
	defer cancel()
	tDataC := make(chan []*internal.D1StringList, len(servers))

	for _, srv := range servers {
		go func(srv server) {
			defer srv.conn.Close()

			// call fumo depending on transaction
			tDataC <- t.dataFunc(t, srv)
		}(srv)

	}

	var sList []*internal.D1StringList
readChannel:
	for i := 0; i < len(servers); i++ {
		select {
		case tc := <-tDataC:
			if tc != nil {
				sList = append(sList, tc...)
			}
		case <-ctx.Done():
			fmt.Println("timeout")
			break readChannel
		}
	}

	return &internal.D2StringList{
		Data2: sList,
	}
}

// collectStandard -  collecting transaction data for one system with normal fumos
func collectStandard(t transaction, srv server) []*internal.D1StringList {

	var lines []*internal.D1StringList

	res, err := srv.conn.Call(t.fumo, t.params)
	if err != nil {
		log.WithFields(log.Fields{
			"server":      srv.name,
			"transaction": t.name,
			"error":       err,
		}).Error("Can't call function module")
		return nil
	}

	if res[t.table] == nil {
		return nil
	}
	for _, data := range res[t.table].([]interface{}) {
		line := data.(map[string]interface{})

		if len(t.filter) == 0 || inFilter(line, t.filter) {

			var stringList internal.D1StringList
			for _, v := range t.fields {
				rawStr := i2String(line[v])

				// only printable characters
				clean := strings.Map(func(r rune) rune {
					if unicode.IsGraphic(r) {
						return r
					}
					return -1
				}, rawStr)

				stringList.Data1 = append(stringList.Data1, clean)
			}

			// when allServers -> first column is servername
			if t.allServers {
				// stringList.Data1[0] = srv.name
				stringList.Data1 = append([]string{srv.name}, stringList.Data1...)
			}
			lines = append(lines, &stringList)
		}
	}
	return lines
}

// collectTable - collecting transaction data for one system with rfc_read_table
func collectTable(t transaction, srv server) []*internal.D1StringList {

	var lines []*internal.D1StringList

	res, err := srv.conn.Call(t.fumo, t.params)
	if err != nil {
		log.WithFields(log.Fields{
			"server":      srv.name,
			"transaction": t.name,
			"error":       err,
		}).Error("Can't call function module")
		return lines
	}

	if res[t.table] == nil {
		return nil
	}
	for _, data := range res[t.table].([]interface{}) {
		line := data.(map[string]interface{})
		if len(t.filter) == 0 || inFilter(line, t.filter) {
			var stringList internal.D1StringList
			stringList.Data1 = strings.Split(line[t.fields[0]].(string), "|")
			for i := range stringList.Data1 {
				stringList.Data1[i] = strings.TrimSpace(stringList.Data1[i])
			}

			lines = append(lines, &stringList)
		}
	}
	return lines
}

// collecting transaction data for one system with help of xmb fumos
// func collectXmb(t transaction, srv server) []*internal.StringList {

// 	var lines []*internal.StringList

// 	params := map[string]interface{}{
// 		"EXTCOMPANY": "World",
// 		"EXTPRODUCT": "Syslog",
// 		"VERSION":    "0.1",
// 		"INTERFACE":  "XMB",
// 	}
// 	_, err := srv.conn.Call("SXMI_LOGON", params)
// 	if err != nil {
// 		return lines
// 	}

// 	// !!!!!!! kann eventuell durch inhalt von collectTable ersetzt werden
// 	if 1 == len(t.fields) {
// 		lines = collectTable(t, srv)
// 	} else {
// 		lines = collectStandard(t, srv)
// 	}

// 	params = map[string]interface{}{
// 		"INTERFACE": "XMB",
// 	}
// 	_, err = srv.conn.Call("SXMI_LOGOFF", params)
// 	if err != nil {
// 		return lines
// 		// log.Fatal("SXMI_LOGOFF: ", err)
// 	}
// 	return lines
// }

// connect -  connect to sap system
func connect(s *internal.SystemInfo) (*gorfc.Connection, error) {
	c, err := gorfc.ConnectionFromParams(
		gorfc.ConnectionParameters{
			"Dest":   s.Sid,
			"User":   s.Username,
			"Passwd": s.Password,
			"Client": s.Client,
			"Lang":   "en",
			"Ashost": s.Hostname,
			"Sysnr":  s.Sysnr,
			// Saprouter: "/H/203.13.155.17/S/3299/W/xjkb3d/H/172.19.137.194/H/",
		},
	)
	if err != nil {
		log.WithFields(log.Fields{
			"system": s.Sid,
			"server": s.Hostname,
			"error":  err,
		}).Error("Can't connect to system with user/password")
		return nil, err
	}

	return c, nil
}

// getServerConnections - get rfc connection for all corresponding servers
func getServerConnections(s *internal.SystemInfo, t transaction) []server {
	var servers []server

	// retrieve connection of first (or only) application server
	defaultSrv, err := connect(s)
	if err != nil {
		log.WithFields(log.Fields{
			"system": s.Sid,
			"server": s.Hostname,
			"error":  err,
		}).Error("can't connect to server")
		return nil
	}

	if !t.allServers {
		return []server{server{"", defaultSrv}}
	}

	// defaultSrv is used to get all application server connections and must be closed
	defer defaultSrv.Close()

	params := map[string]interface{}{}
	r, err := defaultSrv.Call("TH_SERVER_LIST", params)
	if err != nil {
		return nil
	}
	if r["LIST"] == nil {
		return nil
	}

	for _, v := range r["LIST"].([]interface{}) {
		appl := v.(map[string]interface{})
		serverStr := strings.TrimSpace(appl["NAME"].(string))
		info := strings.Split(serverStr, "_")
		s.Hostname = info[0]
		s.Sysnr = info[2]

		srv, err := connect(s)
		if err != nil {
			log.WithFields(log.Fields{
				"server": s.Hostname,
				"error":  err,
			}).Error("error from getServerConnections")
			continue
		}
		servers = append(servers, server{serverStr, srv})
	}
	return servers
}

// getSystemInfo - return info for all systems stored in the badger db
func (env envInfo) getSystemInfo() ([]*internal.SystemInfo, error) {

	// open badger systems db
	store, err := internal.GetStoreRo(filepath.Join(*env.dbstore, "systems"))
	if err != nil {
		return nil, errors.Wrap(err, " :GetSystemList - helpers.GetStore")
	}
	defer store.Close()

	// retrieve all systems
	dbSystems, err := store.GetValues()
	if err != nil {
		return nil, errors.Wrap(err, " :GetSystemList - store.GetValues")
	}

	// !!!!!!!!!!!!!!!!!!!!!!
	// if 0 == len(dbSystems) {
	// 	return nil, errors.New(" :No systems available.")
	// }

	var systems []*internal.SystemInfo
	for _, v := range dbSystems {
		var systemInfo internal.SystemInfo
		err = proto.Unmarshal(v, &systemInfo)
		if err != nil {
			return nil, errors.Wrap(err, " :GetSystemList - proto.Unmarshal")
		}

		// append info to systems slice
		systems = append(systems, &systemInfo)
	}
	return systems, nil
}

// savebadger - save transaction data of all systems to the db
func (env envInfo) saveBadger(sysData systemData) error {
	if !hasData(sysData.sapInfo) {
		return nil
	}
	log.Println("Scrape Time: ", sysData.sid, time.Since(env.ts))

	// check if path to system db exists, otherwise it will be created
	var path string
	for _, dir := range []string{*env.dbstore, sysData.sid} {
		path = filepath.Join(path, dir)
		err := checkDirectory(path)
		if err != nil {
			return errors.Wrap(err, " :saveBadger - checkDirectory")
		}
	}

	// ope system db
	dbStore, err := internal.GetStoreRw(path)
	if err != nil {
		return errors.Wrap(err, " :saveBadger - GetStore")
	}
	defer dbStore.Close()

	protoBytes, err := proto.Marshal(&sysData.sapInfo)
	if err != nil {
		return errors.Wrap(err, " :saveBadger - Marshal")
	}

	zstdBytes, err := zstd.CompressLevel(nil, protoBytes, 14)
	if err != nil {
		return errors.Wrap(err, " :saveBadger - zstd compression")
	}

	// save compressed system data for this scrape
	err = dbStore.SetValue([]byte(env.ts.Format("200601021504")), zstdBytes)
	if err != nil {
		return errors.Wrap(err, " :saveBadger - SetValue")
	}

	return nil
}

// hasData - check if system scrape is empty
func hasData(sysData internal.D3StringList) bool {
	for _, d := range sysData.Data3 {
		if d != nil {
			return true
		}
	}
	return false
}

// inFilter - check if values fulfill filter criteria
// for example WP_STATUS": []interface{}{"running", "!on hold"},
// without ! - only lines with substr are taken
// with ! - only lines without substr are taken
func inFilter(line map[string]interface{}, filter map[string][]interface{}) bool {
	var showLine bool
	for field, values := range filter {
		for _, value := range values {
			s := strings.ToUpper(i2String(line[strings.ToUpper(field)]))
			substr := strings.ToUpper(i2String(value))

			if string("!") == string(substr[0]) {
				if !fExclude(s, substr[1:]) {
					return false
				}
			} else {
				if fInclude(s, substr) {
					showLine = true
				}
			}
		}
	}
	return showLine
}

func fInclude(s, substr string) bool {
	if strings.Contains(s, substr) {
		return true
	}
	return false
}

func fExclude(s, substr string) bool {
	if strings.Contains(s, substr) {
		return false
	}
	return true
}

// i2String -  convert interface int values to string
func i2String(namePart interface{}) string {

	switch val := namePart.(type) {
	case string:
		return strings.TrimSpace(val)
	case int64, int32, int16, int8, int, uint64, uint32, uint8, uint:
		return strings.TrimSpace(fmt.Sprint(val))
	default:
		return ""
	}
}

// getSecret - select the secret information from the db
func (env envInfo) getSecret() (*internal.Secret, error) {
	var secret internal.Secret
	store, err := internal.GetStoreRw(filepath.Join(*env.dbstore, "preferences"))
	if err != nil {
		return nil, errors.Wrap(err, " :getSecret - internal.GetStore")
	}
	defer store.Close()

	secretBytes, err := store.GetValue([]byte("secret"))
	if err != nil {
		if err == badger.ErrKeyNotFound {

			// create secret key, if it doesn't exist
			secret.Name = make(map[string][]byte)
			secret.Name["secretKey"], err = internal.GetSecretKey()
			if err != nil {
				return nil, errors.Wrap(err, "getSecret - GetSecretKey")
			}
			return &secret, nil

		} else {
			return nil, errors.Wrap(err, " :getSecret - store.GetValue")
		}
	}

	// unmarshal secret and return secret map
	if err := proto.Unmarshal(secretBytes, &secret); err != nil {
		return nil, errors.Wrap(err, " unable to unmarshal secret")
	}
	return &secret, nil
}

// setSecret -  write the secret information back to the db
func (env envInfo) setSecret(sid, password string) error {

	secret, err := env.getSecret()
	if err != nil {
		return errors.Wrap(err, " :setSecret - getSecret")
	}

	encPw, err := internal.PwEncrypt([]byte(password), secret.Name["secretKey"])
	if err != nil {
		return errors.Wrap(err, "setSecret - PwEncrypt ")
	}

	// add password to secret map
	secret.Name[sid] = encPw

	// marshal the secret map
	protoSecret, err := proto.Marshal(secret)
	if err != nil {
		return errors.Wrap(err, " ImportTenants - Marshal")
	}

	store, err := internal.GetStoreRw(filepath.Join(*env.dbstore, "preferences"))
	if err != nil {
		return errors.Wrap(err, " :setSecret - helpers.GetStore")
	}
	defer store.Close()

	err = store.SetValue([]byte("secret"), protoSecret)
	if err != nil {
		return errors.Wrap(err, " :setSecret - store.SetValue")
	}
	return nil
}

// getPw - get pw in clear text
func getPw(secret *internal.Secret, sid string) (string, error) {

	// get encrypted tenant pw
	if _, ok := secret.Name[sid]; !ok {
		return "", errors.New("encrypted system pw info does not exist")
	}

	// decrypt tenant password
	pw, err := internal.PwDecrypt(secret.Name[sid], secret.Name["secretKey"])
	if err != nil {
		return "", err
	}
	return pw, nil
}

// getTransaction - return information of used transactions
func (env envInfo) getTransactions() []transaction {
	var transactions []transaction

	tsDay := env.ts.Format("20060102")
	tsDayAll := env.ts.Format("20060102") + "000000"

	tsUTC := time.Now().UTC()
	tLastMinute := tsUTC.Add(time.Minute * -1)

	tLastMinuteS := tLastMinute.Format("1504")
	t1s := tLastMinuteS + "00"
	t2s := tLastMinuteS + "59"

	// !!!!!!!!! err abfangen
	d, err1 := time.Parse("20060102", tsDay)
	t1, err2 := time.Parse("150405", t1s)
	t2, err3 := time.Parse("150405", t2s)

	transactions = append(transactions,
		[]transaction{
			transaction{
				"SM66",
				[]string{"Server", "Type", "Username", "Client", "WP Num", "Time", "Status", "Reason", "Action", "Table", "Report"},
				"TH_WPINFO",
				map[string]interface{}{
					// "SRVNAME": "SRV§§§",
				},
				map[string][]interface{}{
					"WP_STATUS": []interface{}{"on hold", "running"}, // only active processes
					"WP_REPORT": []interface{}{"!saplthfb"},          // no own workprocesses used for analysis
				},
				"WPLIST",
				// allServers transaction -> server will be prepended automatically in collectStandard
				[]string{"WP_TYP", "WP_BNAME", "WP_MANDT", "WP_NO", "WP_ELTIME", "WP_STATUS", "WP_WAITING", "WP_ACTION", "WP_TABLE", "WP_REPORT"},
				true,
				collectStandard,
			},
			transaction{
				"SM37",
				[]string{"Name", "Job Nr.", "Step Nr.", "Server", "WP", "Client", "Pl. Date", "Pl. Time (UTC)", "St. Date", "St. Time (UTC)", "Status", "User"},
				"RFC_READ_TABLE",
				map[string]interface{}{
					"QUERY_TABLE": "TBTCO",
					"DELIMITER":   "|",
					"OPTIONS": []map[string]interface{}{
						// R - running, S - released, Y - ready, A - aborted, F - finished
						map[string]interface{}{"TEXT": "STATUS IN ('R','Y') or STATUS = 'A' AND ENDDATE = '" + tsDay + "'"},
					},
					"FIELDS": []map[string]interface{}{
						map[string]interface{}{"FIELDNAME": "JOBNAME"},
						map[string]interface{}{"FIELDNAME": "JOBCOUNT"},
						map[string]interface{}{"FIELDNAME": "STEPCOUNT"},
						map[string]interface{}{"FIELDNAME": "REAXSERVER"},
						map[string]interface{}{"FIELDNAME": "WPNUMBER"},
						map[string]interface{}{"FIELDNAME": "AUTHCKMAN"},
						map[string]interface{}{"FIELDNAME": "RELDATE"},
						map[string]interface{}{"FIELDNAME": "RELTIME"},
						map[string]interface{}{"FIELDNAME": "STRTDATE"},
						map[string]interface{}{"FIELDNAME": "STRTTIME"},
						map[string]interface{}{"FIELDNAME": "STATUS"},
						map[string]interface{}{"FIELDNAME": "RELUNAME"},
					},
				},
				map[string][]interface{}{},
				"DATA",
				[]string{"WA"},
				false,
				collectTable,
			},
			transaction{
				"SM13",
				[]string{"Client", "Username", "Report", "Tcode", "Server", "Timestamp (UTC)", "RC", "State"},
				"RFC_READ_TABLE",
				map[string]interface{}{
					"QUERY_TABLE": "VBHDR",
					"DELIMITER":   "|",
					"OPTIONS": []map[string]interface{}{
						// map[string]interface{}{"TEXT": "VBDATE BETWEEN " + tsAll + " and '99991231000000'"},
						map[string]interface{}{"TEXT": "VBDATE >= " + tsDayAll},
					},
					"FIELDS": []map[string]interface{}{
						map[string]interface{}{"FIELDNAME": "VBMANDT"},
						map[string]interface{}{"FIELDNAME": "VBUSR"},
						map[string]interface{}{"FIELDNAME": "VBREPORT"},
						map[string]interface{}{"FIELDNAME": "VBTCODE"},
						map[string]interface{}{"FIELDNAME": "VBNAME"},
						map[string]interface{}{"FIELDNAME": "VBDATE"},
						map[string]interface{}{"FIELDNAME": "VBRC"},
						map[string]interface{}{"FIELDNAME": "VBSTATE"},
					},
				},
				map[string][]interface{}{},
				"DATA",
				[]string{"WA"},
				false,
				collectTable,
			},
			transaction{
				"SM12",
				[]string{"Client", "Username", "Table", "Argument", "Mode", "TCode"},
				"ENQUE_READ",
				map[string]interface{}{
					"GARG":    "",
					"GCLIENT": "",
					"GNAME":   "",
					"GUNAME":  "",
				},
				map[string][]interface{}{},
				"ENQ",
				[]string{"GCLIENT", "GUNAME", "GNAME", "GARG", "GMODE", "GTCODE"},
				false,
				collectStandard,
			},
			transaction{
				"SM58",
				[]string{"Caller", "FuMo", "Target System", "Date", "Time (UTC)", "Status", "Transaction Id", "Host", "Transaction", "RC", "Reserv"},
				"RFC_READ_TABLE",
				map[string]interface{}{
					"QUERY_TABLE": "ARFCSSTATE",
					"DELIMITER":   "|",
					"OPTIONS": []map[string]interface{}{
						map[string]interface{}{"TEXT": "ARFCDATUM = '" + tsDay + "' AND ARFCRETURN = ''"},
						// map[string]interface{}{"TEXT": "ARFCDATUM = '<day>' AND ARFCRETURN = ''"},
						// map[string]interface{}{"TEXT": "ARFCDATUM = '20200504' AND ARFCRETURN = '' ORDER BY ARFCDATUM,ARFCUZEIT"},
						// map[string]interface{}{"TEXT": "ARFCRETURN = ''"},
						// map[string]interface{}{"TEXT": "ORDER BY ARFCDATUM,ARFCUZEIT"},
					},
					"FIELDS": []map[string]interface{}{
						map[string]interface{}{"FIELDNAME": "ARFCUSER"},
						map[string]interface{}{"FIELDNAME": "ARFCFNAM"},
						map[string]interface{}{"FIELDNAME": "ARFCDEST"},
						map[string]interface{}{"FIELDNAME": "ARFCDATUM"},
						map[string]interface{}{"FIELDNAME": "ARFCUZEIT"},
						map[string]interface{}{"FIELDNAME": "ARFCSTATE"},
						map[string]interface{}{"FIELDNAME": "ARFCTIDCNT"},
						map[string]interface{}{"FIELDNAME": "ARFCRHOST"},
						map[string]interface{}{"FIELDNAME": "ARFCTCODE"},
						map[string]interface{}{"FIELDNAME": "ARFCRETURN"},
						map[string]interface{}{"FIELDNAME": "ARFCRESERV"},
					},
				},
				map[string][]interface{}{},
				"DATA",
				[]string{"WA"},
				false,
				collectTable,
			},
			transaction{
				"WE02",
				[]string{"IDoc", "Segments", "Status", "Partner Type", "Basic Type", "Extension", "Created On", "Created at (UTC)", "Msg Type", "Direction", "Port"},
				"RFC_READ_TABLE",
				map[string]interface{}{
					"QUERY_TABLE": "EDIDC",
					"DELIMITER":   "|",
					"OPTIONS": []map[string]interface{}{
						// map[string]interface{}{"TEXT": "VBDATE BETWEEN " + tsAll + " and '99991231000000'"},
						// map[string]interface{}{"TEXT": "CREDAT = '" + tsDay + "'"},
						map[string]interface{}{"TEXT": "CREDAT = '" + tsDay + "' and CRETIM between '" + t1s + "' and '" + t2s + "'"},
					},
					"FIELDS": []map[string]interface{}{
						map[string]interface{}{"FIELDNAME": "DOCNUM"},
						map[string]interface{}{"FIELDNAME": "MAXSEGNUM"},
						map[string]interface{}{"FIELDNAME": "STATUS"},
						map[string]interface{}{"FIELDNAME": "RCVPRT"},
						map[string]interface{}{"FIELDNAME": "RCVPRN"},
						map[string]interface{}{"FIELDNAME": "IDOCTP"},
						map[string]interface{}{"FIELDNAME": "CREDAT"},
						map[string]interface{}{"FIELDNAME": "CRETIM"},
						map[string]interface{}{"FIELDNAME": "MESTYP"},
						map[string]interface{}{"FIELDNAME": "DIRECT"},
						map[string]interface{}{"FIELDNAME": "RCVPOR"},
					},
				},
				map[string][]interface{}{},
				"DATA",
				[]string{"WA"},
				false,
				collectTable,
			},
			transaction{
				"AL08",
				[]string{"Server", "Client", "Username", "Tcode", "Terminal", "Time (UTC)", "Type", "Gui-Version", "IP-Address"},
				"TH_USER_LIST",
				map[string]interface{}{},
				map[string][]interface{}{
					// show only entries with action in the last minute
					// !!!!!!!!!!!
					"ZEIT": []interface{}{tLastMinuteS},
				},
				"USRLIST",
				// allServers transaction -> server will be prepended automatically in collectStandard
				[]string{"MANDT", "BNAME", "TCODE", "TERM", "ZEIT", "TYPE", "GUIVERSION", "HOSTADDR"},
				true,
				collectStandard,
			},
		}...)

	if err1 == nil && err2 == nil && err3 == nil {

		// log.Println("OOOOOOOOOOOOOOOOOOOO: ", d, t1, t2)
		transactions = append(transactions,

			transaction{
				"SM21",
				[]string{"Time (UTC)", "User", "Host", "Component", "Report", "Transaction", "Text"},
				"/SDF/GET_SYS_LOG",
				map[string]interface{}{
					"DATE_FROM": d,
					"TIME_FROM": t1,
					"DATE_TO":   d,
					"TIME_TO":   t2,
				},
				map[string][]interface{}{},
				"ET_E2E_LOG",
				[]string{"E2E_TIME", "E2E_USER", "E2E_HOST", "FIELD2", "FIELD4", "FIELD5", "FIELD6"},
				false,
				collectStandard,
			})
	}
	return transactions
}
