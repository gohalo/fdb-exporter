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
