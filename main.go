package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/rs/zerolog/log"

	"github.com/gohalo/fdb-exporter/internal/models"
)

var (
	db        fdb.Database
	statusKey []byte
)

var (
	address        = flag.String("addr", ":8080", "Listen address for http service")
	fdbClusterFile = flag.String("fdb_cluster_file", "/etc/foundationdb/fdb.cluster", "Cluster file to connect to fdb")
	fdbAPIVersion  = flag.String("fdb_api_version", "710", "API version to use for the fdb connection")
)

func init() {
	statusKey = append([]byte{255, 255}, []byte("/status/json")...)
}

// import (
// 	"bufio"
// 	"encoding/json"
// 	"os"
// 	"regexp"
//
// 	"github.com/rs/zerolog/log"
// 	ulog "github.com/tigrisdata/fdb-exporter/util/log"
//
// 	"github.com/tigrisdata/fdb-exporter/models"
// )
//
// const DefaultApiVersion = "710"
//
// var (
// 	tls     = false
// 	netopts fdb.NetworkOptions
// )
//
// func getFdb() fdb.Database {
// 	clusterFile := os.Getenv("FDB_CLUSTER_FILE")
// 	if clusterFile == "" {
// 		log.Error().Msg("set FDB_CLUSTER_FILE environment variable")
// 		os.Exit(1)
// 	}
// 	apiVersionStr := os.Getenv("FDB_API_VERSION")
// 	if apiVersionStr == "" {
// 		apiVersionStr = DefaultApiVersion
// 	}
// 	apiVersion, err := strconv.Atoi(apiVersionStr)
// 	if err != nil {
// 		log.Error().Str("FDB_API_VERSION", apiVersionStr).Msg("Could not convert api version to integer")
// 	}
// 	tlsCaFile := os.Getenv("FDB_TLS_CA_FILE")
// 	tlsCertFile := os.Getenv("FDB_TLS_CERT_FILE")
// 	tlsKeyFile := os.Getenv("FDB_TLS_KEY_FILE")
// 	tlsCheckValid := os.Getenv("FDB_TLS_VERIFY_PEERS")
//
// 	fdb.MustAPIVersion(apiVersion)
// 	netopts = fdb.Options()
//
// 	// Check TLS in clusterfile
// 	tls = isTLSmode(&clusterFile)
// 	if tls {
// 		// TLS params for fdbclient
// 		if len(tlsCaFile) > 0 {
// 			err := netopts.SetTLSCaPath(tlsCaFile)
// 			if err != nil {
// 				log.Error().Err(err).Str("msg", "Error cannot set CA").Msg("TLS CA error")
// 			}
// 			log.Info().Str("msg", "TLS CA File").Str("fdb.tls-ca-file", tlsCaFile).Msg("TLS CA file set")
// 		}
//
// 		if len(tlsCertFile) > 0 {
// 			err := netopts.SetTLSCertPath(tlsCertFile)
// 			if err != nil {
// 				log.Error().Err(err).Str("msg", "Error cannot set Cert").Msg("TLS Cert error")
// 			}
// 			log.Info().Str("msg", "TLS Cert File").Str("fdb.tls-cert-file", tlsCertFile).Msg("TLS cert file set")
// 		}
//
// 		if len(tlsKeyFile) > 0 {
// 			err := netopts.SetTLSKeyPath(tlsKeyFile)
// 			if err != nil {
// 				log.Error().Err(err).Str("msg", "Error cannot set Key").Msg("TLS Key error")
// 			}
// 			log.Info().Str("msg", "TLS Private Key").Str("fdb.tls-key-file", tlsKeyFile).Msg("TLS private key set")
// 		}
//
// 		err := netopts.SetTLSVerifyPeers([]byte(tlsCheckValid))
// 		if err != nil {
// 			log.Error().Err(err).Str("msg", "Error cannot set VerifyPeers").Msg("TLS VerifyPeers error")
// 		}
// 		log.Info().Str("msg", "TLS VerifyPeers").Str("fdb.tls-check-valid", tlsCheckValid).Msg("TLS VerifyPeers set")
// 	}
// 	db, err := fdb.OpenDatabase(clusterFile)
// 	if err != nil {
// 		log.Error().Str("cluster_file", clusterFile).Msg("failed to open database using cluster file")
// 		os.Exit(1)
// 	}
// 	return db
// }
//
// func GetStatus() (*models.FullStatus, error) {
// 	conn := getFdb()
// 	var status models.FullStatus
// 	statusKey := append([]byte{255, 255}, []byte("/status/json")...)
// 	statusJson, err := conn.ReadTransact(func(t fdb.ReadTransaction) (interface{}, error) {
// 		return t.Get(fdb.Key(statusKey)).Get()
// 	})
// 	if err != nil {
// 		msg := "failed to get status"
// 		ulog.E(err, msg)
// 		return nil, err
// 	}
//
// 	err = json.Unmarshal(statusJson.([]byte), &status)
// 	if err != nil {
// 		msg := "failed to unmarshal status"
// 		ulog.E(err, msg)
// 		return nil, err
// 	}
// 	return &status, nil
// }
//
// func isTLSmode(c *string) bool {
// 	// Find if must run in TLS mode
// 	file, err := os.Open(*c)
// 	if err != nil {
// 		log.Fatal().Err(err).Msgf("failed to open %s", *c)
// 	}
//
// 	scanner := bufio.NewScanner(file)
// 	scanner.Split(bufio.ScanLines)
// 	var text []string
//
// 	for scanner.Scan() {
// 		text = append(text, scanner.Text())
// 	}
//
// 	// The method os.File.Close() is called
// 	// on the os.File object to close the file
// 	file.Close()
//
// 	// and then a loop iterates through
// 	// and prints each of the slice values.
// 	for _, eachLn := range text {
// 		//docker:docker@172.19.0.2:4500:tls
// 		tls, err = regexp.Match(`[0-9]+:tls`, []byte(eachLn))
// 	}
// 	return tls
// }

func FDBMetrics(w http.ResponseWriter, r *http.Request) {
	var status models.Status

	start := time.Now()
	if status_, err := db.ReadTransact(func(t fdb.ReadTransaction) (any, error) {
		return t.Get(fdb.Key(statusKey)).Get()
	}); err != nil {
		log.Error().Msgf("Read from fdb failed, %v.", err)
		return
	} else if err = json.Unmarshal(status_.([]byte), &status); err != nil {
		log.Error().Msgf("Unmarshal status failed, %v.", err)
		return
	}
	elapsed := time.Since(start)
	log.Info().Msgf("Get status elapsed %dms.", elapsed.Milliseconds())

	status.Client.DumpPromethusMetrics(w)
	status.Cluster.DumpPromethusMetrics(w)

	fmt.Fprintln(w, "# TYPE fdb_exporter_latency_ms gauge")
	fmt.Fprintf(w, "fdb_exporter_latency_ms %d\n", elapsed.Milliseconds())
}

func getFromEnvWithDefault(envs string, defval string) string {
	if enval := os.Getenv(envs); enval != "" {
		return enval
	}
	return defval
}

func fdbInit() error {
	apiver := getFromEnvWithDefault("FDB_API_VERSION", *fdbAPIVersion)
	version, err := strconv.Atoi(apiver)
	if err != nil {
		log.Error().Str("FDB_API_VERSION", apiver).Msg("Could not convert api version to integer")
		return err
	}
	fdb.MustAPIVersion(version)

	// export FDB_CLUSTER_FILE="/home/olap/fdb/home/conf/fdb.cluster"
	clusterfile := getFromEnvWithDefault("FDB_CLUSTER_FILE", *fdbClusterFile)
	if db_, err := fdb.OpenDatabase(clusterfile); err != nil {
		log.Info().Str("FDB_CLUSTER_FILE", clusterfile).Msgf("Open with cluster file failed, %v.", err)
		return err
	} else {
		db = db_
	}
	return nil
}

func main() {
	if err := fdbInit(); err != nil {
		return
	}

	http.HandleFunc("/", FDBMetrics)
	http.HandleFunc("/metrics", FDBMetrics)

	log.Info().Msgf("Start to listen on '%s'.", *address)
	if err := http.ListenAndServe(*address, nil); err != nil {
		log.Fatal().Msgf("Listen to '%s' failed, %v.", *address, err)
	}
}
