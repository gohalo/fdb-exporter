package models

import (
	"fmt"
	"io"
)

type LockState struct {
	Locked bool `json:"locked"`
}

type Layers struct {
	Valid bool `json:"_valid"`
	// Backup       *Backup   `json:"backup"`
	// DrBackup     *DrBackup `json:"dr_backup"`
	// DrBackupDest *DrBackup `json:"dr_backup_dest"`
}

type Clients struct {
	Count int `json:"count"`
}

type LatencyProbe struct {
	BatchPriorityTransactionStartSeconds     float64 `json:"batch_priority_transaction_start_seconds"`
	CommitSeconds                            float64 `json:"commit_seconds"`
	ImmediatePriorityTransactionStartSeconds float64 `json:"immediate_priority_transaction_start_seconds"`
	ReadSeconds                              float64 `json:"read_seconds"`
	TransactionStartSeconds                  float64 `json:"transaction_start_seconds"`
}

func (p *LatencyProbe) DumpPromethusMetrics(w io.Writer) error {
	fmt.Fprintln(w, "# HELP Time to perform a single read")
	fmt.Fprintln(w, "# TYPE fdb_cluster_latency_probe_read_seconds gauge")
	fmt.Fprintf(w, "fdb_cluster_latency_probe_read_seconds %f\n", p.ReadSeconds)

	fmt.Fprintln(w, "# HELP Time to commit a sample transaction")
	fmt.Fprintln(w, "# TYPE fdb_cluster_latency_probe_commit_seconds gauge")
	fmt.Fprintf(w, "fdb_cluster_latency_probe_commit_seconds %f\n", p.CommitSeconds)

	fmt.Fprintln(w, "# HELP Time to start a sample transaction at normal priority")
	fmt.Fprintln(w, "# TYPE fdb_cluster_latency_probe_transaction_start_seconds gauge")
	fmt.Fprintf(w, "fdb_cluster_latency_probe_transaction_start_seconds %f\n", p.TransactionStartSeconds)

	fmt.Fprintln(w, "# HELP Time to start a sample transaction at system immediate priority")
	fmt.Fprintln(w, "# TYPE fdb_cluster_latency_probe_immediate_priority_transaction_start_seconds gauge")
	fmt.Fprintf(w, "fdb_cluster_latency_probe_immediate_priority_transaction_start_seconds %f\n", p.ImmediatePriorityTransactionStartSeconds)

	fmt.Fprintln(w, "# HELP Time to start a sample transaction at batch priority")
	fmt.Fprintln(w, "# TYPE fdb_cluster_latency_probe_batch_priority_transaction_start_seconds gauge")
	fmt.Fprintf(w, "fdb_cluster_latency_probe_batch_priority_transaction_start_seconds %f\n", p.BatchPriorityTransactionStartSeconds)
	return nil
}

type ExcludedServer struct {
	Address string `json:"address"`
}

type Configuration struct {
	BackupWorkerEnabled            int64            `json:"backup_worker_enabled"`
	BlobGranulesEnabled            int64            `json:"blob_granules_enabled"`
	CommitProxies                  int64            `json:"commit_proxies"`
	CoordinatorsCount              int64            `json:"coordinators_count"`
	ExcludedServers                []ExcludedServer `json:"excluded_servers"`
	GrvProxies                     int64            `json:"grv_proxies"`
	LogRouters                     int64            `json:"log_routers"`
	Logs                           int64            `json:"logs"`
	PerpetualStorageWiggle         int64            `json:"perpetual_storage_wiggle"`
	PerpetualStorageWiggleLocality string           `json:"perpetual_storage_wiggle_locality"`
	Proxies                        int64            `json:"proxies"`
	RedundancyMode                 string           `json:"redundancy_mode"`
	RemoteLogs                     int64            `json:"remote_logs"`
	Resolvers                      int64            `json:"resolvers"`
	StorageEngine                  string           `json:"storage_engine"`
	StorageMigrationType           string           `json:"storage_migration_type"`
	TenantMode                     string           `json:"tenant_mode"`
	UsableRegions                  int64            `json:"usable_regions"`
}

type FaultTolerance struct {
	MaxZoneFailuresWithoutLosingAvailability int64 `json:"max_zone_failures_without_losing_availability"`
	MaxZoneFailuresWithoutLosingData         int64 `json:"max_zone_failures_without_losing_data"`
}

type Lag struct {
	Seconds  float64 `json:"seconds"`
	Versions int64   `json:"versions"`
}

type Hz struct {
	Hz float64 `json:"hz"`
}

type PageCache struct {
	LogHitRate     float64 `json:"log_hit_rate"`
	StorageHitRate float64 `json:"storage_hit_rate"`
}

type RecoveryState struct {
	ActiveGenerations         int64   `json:"active_generations"`
	Description               string  `json:"description"`
	Name                      string  `json:"name"`
	SecondsSinceLastRecovered float64 `json:"seconds_since_last_recovered"`
}

func (r *RecoveryState) DumpPromethusMetrics(w io.Writer) error {
	fmt.Fprintln(w, "# HELP Recovery state info")
	fmt.Fprintln(w, "# TYPE fdb_cluster_recovery_state gauge")
	fmt.Fprintf(w, "fdb_cluster_recovery_state{name=\"%s\",description=\"%s\"} 1\n", r.Name, r.Description)

	fmt.Fprintln(w, "# HELP Recovery state active generations count")
	fmt.Fprintln(w, "# TYPE fdb_cluster_recovery_state_active_generations gauge")
	fmt.Fprintf(w, "fdb_cluster_recovery_state_active_generations %d\n", r.ActiveGenerations)

	fmt.Fprintln(w, "# HELP Recovery state active generations count")
	fmt.Fprintln(w, "# TYPE fdb_cluster_recovery_state_seconds_since_last_recovered gauge")
	fmt.Fprintf(w, "fdb_cluster_recovery_state_seconds_since_last_recovered %f\n", r.SecondsSinceLastRecovered)
	return nil
}

// Cluster top level status
type ClusterStatus struct {
	ClusterControllerTimestamp int64              `json:"cluster_controller_timestamp"`
	Configuration              Configuration      `json:"configuration"`
	ConnectionString           string             `json:"connection_string"`
	Data                       Data               `json:"data"`
	DatabaseAvailable          bool               `json:"database_available"`
	DatabaseLockState          LockState          `json:"database_lock_state"`
	Clients                    Clients            `json:"clients"`
	DatacenterLag              Lag                `json:"datacenter_lag"`
	DegradedProcesses          int64              `json:"degraded_processes"`
	FaultTolerance             *FaultTolerance    `json:"fault_tolerance"`
	FullReplication            bool               `json:"full_replication"`
	Generation                 int64              `json:"generation"`
	LatencyProbe               LatencyProbe       `json:"latency_probe"`
	ProtocolVersion            string             `json:"protocol_version"`
	Qos                        Qos                `json:"qos"`
	RecoveryState              RecoveryState      `json:"recovery_state"`
	Workload                   Workload           `json:"workload"`
	Layers                     Layers             `json:"layers"`
	Processes                  map[string]Process `json:"processes"`
	// Logs                       []Log              `json:"logs"`
	// Machines                   map[string]Machine `json:"machines"`
	// Messages                   []ClusterMessage   `json:"messages"`
	// PageCache                  *PageCache         `json:"page_cache"`
}

func (c *ClusterStatus) DumpPromethusMetrics(w io.Writer) error {
	fmt.Fprintln(w, "# HELP Whether or not database is available. 1 if is, 0 otherwise")
	fmt.Fprintln(w, "# TYPE fdb_cluster_database_available gauge")
	if c.DatabaseAvailable {
		fmt.Fprintln(w, "fdb_cluster_database_available 1")
	} else {
		fmt.Fprintln(w, "fdb_cluster_database_available 0")
	}

	fmt.Fprintln(w, "# HELP Whether or not database is locked. 1 if is, 0 otherwise")
	fmt.Fprintln(w, "# TYPE fdb_cluster_database_locked gauge")
	if c.DatabaseLockState.Locked {
		fmt.Fprintln(w, "fdb_cluster_database_locked 1")
	} else {
		fmt.Fprintln(w, "fdb_cluster_database_locked 0")
	}

	fmt.Fprintln(w, "# HELP Count of clients")
	fmt.Fprintln(w, "# TYPE fdb_cluster_clients_count gauge")
	fmt.Fprintf(w, "fdb_cluster_clients_count %d\n", c.Clients.Count)

	fmt.Fprintln(w, "# HELP Datacenter lag in seconds")
	fmt.Fprintln(w, "# TYPE fdb_cluster_datacenter_lag_seconds gauge")
	fmt.Fprintf(w, "fdb_cluster_datacenter_lag_seconds %f\n", c.DatacenterLag.Seconds)

	c.Workload.DumpPromethusMetrics(w)
	c.LatencyProbe.DumpPromethusMetrics(w)
	c.Data.DumpPromethusMetrics(w)
	c.Qos.DumpPromethusMetrics(w)
	c.RecoveryState.DumpPromethusMetrics(w)
	for id, proc := range c.Processes {
		proc.DumpPromethusMetrics(w, id)
	}
	return nil
}
