// Package cmd ...
package cmd

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	strings "strings"

	"net/http"
	"path/filepath"
	"time"

	"github.com/DataDog/zstd"
	"github.com/improbable-eng/grpc-web/go/grpcweb"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/ulranh/saphistory/internal"
)

// start data fetching in go routine
// afterwards start web server
func (env *envInfo) web() error {

	// if fetch flag == true -> get sap data at the beginning of every minute
	if *env.fetch {
		go func() {
			// wait for start of next minute
			sleeptime := 60 - time.Now().Second()
			time.Sleep(time.Duration(sleeptime) * time.Second)

			for {
				env.ts = time.Now()
				err := env.save()
				if err != nil {
					log.WithFields(log.Fields{
						"error": err,
					}).Error("Cannot select data for systems")
				}

				// sleep until next scrape
				time.Sleep(time.Minute - time.Since(env.ts))
			}
		}()
	}

	// server mux
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(assets))

	// after browser refresh start with / or /systems
	mux.HandleFunc("/systemStatus/", func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/", http.StatusMovedPermanently)
	})
	mux.HandleFunc("/systems/", func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/", http.StatusMovedPermanently)
	})

	// grpc server
	grpcServer := grpc.NewServer()
	internal.RegisterSapHistoryServiceServer(grpcServer, env)
	wrappedServer := grpcweb.WrapServer(grpcServer)

	// create tcp connection
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", env.port))
	if err != nil {
		return errors.Wrap(err, " web - net.Listen")
	}

	// tls config
	var tlsConfig *tls.Config
	tlsPath, ok := os.LookupEnv("VUE_APP_TLS_PATH")
	if ok && existsDirectory(tlsPath) {

		tlsCert := filepath.Join(tlsPath, "tls.crt")
		tlsKey := filepath.Join(tlsPath, "tls.key")
		cert, err := tls.LoadX509KeyPair(tlsCert, tlsKey)
		if err != nil {
			return errors.Wrap(err, " web - LoadX509KeyPair")
		}
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
			NextProtos:   []string{"h2"},
		}
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", env.port),
		Handler:      matchingHandlerFunc(wrappedServer, mux),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		TLSConfig:    tlsConfig,
	}

	// with correct tlsConfig start secure server otherwise standard htp
	if tlsConfig != nil {
		err = srv.Serve(tls.NewListener(conn, srv.TLSConfig))
	} else {
		err = srv.Serve(conn)
	}

	return nil
}

// // development version!
// // start data fetching in go routine
// // afterwards start web server
// func (env *envInfo) web() error {
//
// 	// if fetch flag == true -> get sap data at the beginning of every minute
// 	if *env.fetch {
// 		go func() {
// 			// waiting for start of next minute
// 			sleeptime := 60 - time.Now().Second()
// 			time.Sleep(time.Duration(sleeptime) * time.Second)
//
// 			for {
//
// 				// !!!!!!! UTC checken ??????
// 				// ts := time.Now().UTC()
// 				env.ts = time.Now()
// 				err := env.save()
// 				if err != nil {
// 					log.WithFields(log.Fields{
// 						"error": err,
// 					}).Error("Cannot select data for systems")
// 				}
//
// 				// sleep until next scrape
// 				time.Sleep(time.Minute - time.Since(env.ts))
// 			}
// 		}()
// 	}
//
// 	grpcServer := grpc.NewServer()
// 	internal.RegisterSapHistoryServiceServer(grpcServer, env)
//
// 	wrappedServer := grpcweb.WrapServer(grpcServer)
// 	grpcHandler := func(resp http.ResponseWriter, req *http.Request) {
// 		allowCors(resp, req)
// 		wrappedServer.ServeHTTP(resp, req)
// 	}
//
// 	httpServer := http.Server{
// 		Addr:         fmt.Sprintf(":%d", env.port),
// 		Handler:      http.HandlerFunc(grpcHandler),
// 		WriteTimeout: 10 * time.Second,
// 		ReadTimeout:  10 * time.Second,
// 	}
//
// 	if err := httpServer.ListenAndServe(); err != nil {
// 		grpclog.Fatalf("failed starting http server: %v", err)
// 	}
//
// 	return nil
// }

// matchingHandlerFunc - differentiate between grpc and other packets. this way it is possible to use
// one port for grpc and the rest.
func matchingHandlerFunc(grpcServer *grpcweb.WrappedGrpcServer, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			// fmt.Println("GRPC: ", grpcServer)
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
		allowCors(w, r)
	})
}

// allowCors - cors settings (?????)
func allowCors(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Credentials", "true")
	resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	resp.Header().Set("Access-Control-Expose-Headers", "grpc-status, grpc-message")
	resp.Header().Set("Access-Control-Allow-Headers", "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, XMLHttpRequest, x-user-agent, x-grpc-web, grpc-status, grpc-message")
}

// GetSapStatus - retrieve correct system transaction data from badger db and return them to the client
func (env *envInfo) GetSapStatus(ctx context.Context, selValues *internal.SapSelection) (*internal.TransactionData, error) {

	// get tcodes
	tcodes, err := env.getTcodes(selValues.Ts)
	if err != nil {
		return nil, errors.Wrap(err, " :GetSapStatus - headline")
	}

	// get headlines
	headlines, err := env.getHeadlines(selValues.Ts)
	if err != nil {
		return nil, errors.Wrap(err, " :GetSapStatus - headline")
	}

	// open badger database for the correct system
	store, err := internal.GetStoreRo(filepath.Join(*env.dbstore, selValues.Sid))
	if err != nil {
		return nil, errors.Wrap(err, " :GetSapStatus - GetStoreRo")

	}
	defer store.Close()

	// find matching key for the timestamp
	k, err := store.FindMatchingKey(selValues.Ts, selValues.Direction)
	if err != nil {
		return nil, errors.Wrap(err, " :GetSapStatus - FindMatchingKey")
	}

	// get compressed data for the key
	v, err := store.GetValue(k)
	if err != nil {
		return nil, errors.Wrap(err, " :GetSapStatus - store.GetValue")
	}

	// uncompress data
	protoBytes, err := zstd.Decompress(nil, v)
	if err != nil {
		return nil, errors.Wrap(err, " :GetSapStatus - zstd.Decompress")
	}

	// unmarshal protobuf data
	var transactionData internal.D3StringList
	err = proto.Unmarshal(protoBytes, &transactionData)
	if err != nil {
		return nil, errors.Wrap(err, " :GetSapStatus - Unmarshal")
	}

	return &internal.TransactionData{
		Ts:     string(k),
		Tcodes: tcodes,
		Hdata:  headlines,
		Tdata:  &transactionData,
	}, nil
}

// GetSystemList - retrieve systems from badger db and return data to the client
func (env *envInfo) GetSystemList(ctx context.Context, nothing *internal.Nothing) (*internal.SystemList, error) {

	systems, err := env.getSystemInfo()
	if err != nil {
		return nil, errors.Wrap(err, " :GetSystemList - proto.Unmarshal")
	}

	return &internal.SystemList{
		Systems: systems,
	}, nil
}

// UpdateSystem - update information of one system in the badger db
func (env *envInfo) UpdateSystem(ctx context.Context, system *internal.SystemInfo) (*internal.Nothing, error) {

	// save new encrypted password
	if len(system.Password) > 0 {

		// store decrypted password
		err := env.setSecret(system.Sid, system.Password)
		if err != nil {
			return nil, errors.Wrap(err, " :GetSystemList - helpers.GetStore")
		}

	}
	system.Password = ""

	// connect to db store
	store, err := internal.GetStoreRw(filepath.Join(*env.dbstore, "systems"))
	if err != nil {
		return nil, errors.Wrap(err, " :GetSystemList - helpers.GetStore")
	}
	defer store.Close()

	// marshal the system information except the password ...
	protoSystem, err := proto.Marshal(system)
	if err != nil {
		return nil, errors.Wrap(err, " :setValue - proto.Marshal")
	}

	// ... save data to the database
	err = store.SetValue([]byte(system.Sid), protoSystem)
	if err != nil {
		return nil, errors.Wrap(err, " :GetSystemList - store.GetValues")
	}
	return &internal.Nothing{}, nil
}

// DeleteSystem - delete system from badger db
func (env *envInfo) DeleteSystem(ctx context.Context, system *internal.SystemInfo) (*internal.Nothing, error) {
	dbPath := filepath.Join(*env.dbstore, "systems")
	store, err := internal.GetStoreRw(dbPath)
	if err != nil {
		return nil, errors.Wrap(err, " DeleteSystem - internal.GetStore")
	}
	defer store.Close()

	err = store.DeleteValue([]byte(system.Sid))
	if err != nil {
		return nil, errors.Wrap(err, " :DeleteSystem - store.DeleteValue")
	}

	return &internal.Nothing{}, nil
}

// helper functions
//

// getHeadlines - return headlines from the badger db
func (env *envInfo) getHeadlines(ts string) (*internal.D2StringList, error) {
	dbPath := filepath.Join(*env.dbstore, "headlines")
	store, err := internal.GetStoreRo(dbPath)
	if err != nil {
		return nil, errors.Wrap(err, " GetHeadline - internal.GetStore")
	}
	defer store.Close()

	// find matching key for the timestamp
	// k, err := store.FindMatchingKey(env.ts.Format("200601021504"), 0)
	k, err := store.FindMatchingKey(ts, 0)
	if err != nil {
		return nil, errors.Wrap(err, " :getHeadline - FindMatchingKey")
	}
	v, err := store.GetValue(k)
	if err != nil {
		return nil, errors.Wrap(err, " :getHeadline - store.GetValue")
	}

	var protoData internal.D2StringList
	err = proto.Unmarshal(v, &protoData)
	if err != nil {
		return nil, errors.Wrap(err, " :getHeadline - Unmarshal")
	}
	return &protoData, nil
}

// getTcodes - return tcodes from the badger db
func (env *envInfo) getTcodes(ts string) (*internal.D1StringList, error) {
	dbPath := filepath.Join(*env.dbstore, "tcodes")
	store, err := internal.GetStoreRo(dbPath)
	if err != nil {
		return nil, errors.Wrap(err, " getTcodes - internal.GetStore")
	}
	defer store.Close()

	// find matching key for the timestamp
	// k, err := store.FindMatchingKey(env.ts.Format("200601021504"), 0)
	k, err := store.FindMatchingKey(ts, 0)
	if err != nil {
		return nil, errors.Wrap(err, " :getTcodes - FindMatchingKey")
	}

	v, err := store.GetValue(k)
	if err != nil {
		return nil, errors.Wrap(err, " :getTcodes - store.GetValue")
	}

	var protoData internal.D1StringList
	err = proto.Unmarshal(v, &protoData)
	if err != nil {
		return nil, errors.Wrap(err, " :getTcodes - Unmarshal")
	}
	return &protoData, nil
}
