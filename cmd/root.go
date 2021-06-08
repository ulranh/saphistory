package cmd

import (
	"flag"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/ulranh/saphistory/internal"
)

// environment variables
type envInfo struct {
	dbstore *string
	// period  *uint
	fetch   *bool
	timeout *uint
	port    int
	ts      time.Time
}

// Root -  cmd starting point
func Root() {

	var env envInfo
	env.dbstore = flag.String("dbstore", "", "Path of database directory. (required)")
	// env.period = flag.Uint("period", 1, "Retrieve data every <period> minutes")
	env.timeout = flag.Uint("timeout", 5, "Timeout of sapnwrfc call in <timeout> seconds")
	env.fetch = flag.Bool("fetch", true, "Fetch SAP data")
	flag.Parse()

	if *env.dbstore == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// if no port assignment through environment variable -> port: 8000
	port, err := strconv.Atoi(os.Getenv("VUE_APP_PORT"))
	if err == nil && port > 0 {
		env.port = port
	} else {
		env.port = 8000
	}

	// create database path for sap systems
	for _, path := range []string{*env.dbstore, filepath.Join(*env.dbstore, "systems"), filepath.Join(*env.dbstore, "preferences"), filepath.Join(*env.dbstore, "tcodes"), filepath.Join(*env.dbstore, "headlines")} {
		err := checkDirectory(path)
		if err != nil {
			log.Fatal("Problem creating database diretory: ", path)
		}
	}

	// check if something regarding headlines or transactions has changed
	env.ts = time.Now()
	if env.transactionChanged() {

		// update headline and transaction entries in database
		err = env.updateTransactionInfo()
		if err != nil {
			log.Fatal("Problem with Root - updateTransactionInfo: ", err)
		}
	}

	err = env.web()
	if err != nil {
		log.Error("Server problem", err)
		os.Exit(1)
	}
}

func (env envInfo) transactionChanged() bool {

	headlines, err := env.getHeadlines(env.ts.Format("200601021504"))
	if err != nil {
		return true
	}
	tcodes, err := env.getTcodes(env.ts.Format("200601021504"))
	if err != nil {
		return true
	}

	transactions := env.getTransactions()

	// number of transactions has changed
	if len(transactions) != len(tcodes.Data1) {
		return true
	}

	for i, t := range transactions {

		// order description of tcodes has changed
		if t.name != tcodes.Data1[i] {
			return true
		}

		// number of headline columns has changed
		if len(t.head) != len(headlines.Data2[i].Data1) {
			return true
		}

		// order description of tcodes has changed
		if !equalSlice(t.head, headlines.Data2[i].Data1) {
			return true
		}
	}
	return false
}

// save correct tcode and headline values to the badger db
func (env envInfo) updateTransactionInfo() error {

	var hData []*internal.D1StringList
	var tcodes []string
	for _, t := range env.getTransactions() {

		hData = append(hData, &internal.D1StringList{
			Data1: t.head,
		})
		tcodes = append(tcodes, t.name)
	}

	// marshal tcode values ...
	protoTcode, err := proto.Marshal(&internal.D1StringList{
		Data1: tcodes,
	})
	if err != nil {
		return errors.Wrap(err, " :setValue - proto.Marshal")
	}

	// ... open db ..
	store, err := internal.GetStoreRw(filepath.Join(*env.dbstore, "tcodes"))
	if err != nil {
		return errors.Wrap(err, " :GetSystemList - helpers.GetStore")
	}
	defer store.Close()

	// ... save tcodes to the db
	err = store.SetValue([]byte(env.ts.Format("200601021504")), protoTcode)
	if err != nil {
		return errors.Wrap(err, " :GetSystemList - store.GetValues")
	}

	// marshal headline values ...
	protoHead, err := proto.Marshal(&internal.D2StringList{
		Data2: hData,
	})
	if err != nil {
		return errors.Wrap(err, " :setValue - proto.Marshal")
	}

	// ... open db ...
	store, err = internal.GetStoreRw(filepath.Join(*env.dbstore, "headlines"))
	if err != nil {
		return errors.Wrap(err, " :GetSystemList - helpers.GetStore")
	}
	defer store.Close()

	// ... save headlines to the database
	err = store.SetValue([]byte(env.ts.Format("200601021504")), protoHead)
	if err != nil {
		return errors.Wrap(err, " :GetSystemList - store.GetValues")
	}

	return nil
}

// equalSlice - check if two slices are equal
func equalSlice(sl1, sl2 []string) bool {
	if (sl1 == nil) != (sl2 == nil) {
		return false
	}

	if len(sl1) != len(sl2) {
		return false
	}

	for i := range sl1 {
		if sl2[i] != sl2[i] {
			return false
		}
	}
	return true
}

// checkDirectory - create directory if it does not exist yet
func checkDirectory(path string) error {

	if !existsDirectory(path) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return errors.Wrap(err, " :checkDirectory - MkdirAll")
		}
	}
	return nil
}

// existsDirectory - check if directory exists
func existsDirectory(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
