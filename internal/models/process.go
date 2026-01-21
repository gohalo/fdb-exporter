package models

import (
	"fmt"
	"io"
)

type ProcessCpu struct {
	UsageCores float64 `json:"usage_cores"`
}

type ProcessDiskCounter struct {
	Counter int     `json:"counter"`
	Hz      float64 `json:"hz"`
	Sectors float64 `json:"sectors"`
}

type ProcessDisk struct {
	Busy       float64            `json:"busy"`
	FreeBytes  int                `json:"free_bytes"`
	Reads      ProcessDiskCounter `json:"reads"`
	TotalBytes int                `json:"total_bytes"`
	Writes     ProcessDiskCounter `json:"writes"`
}

func (d *ProcessDisk) DumpPromethusMetrics(w io.Writer, id string, p *Process) error {
	fmt.Fprintln(w, "# HELP Disk usage from 0.0 (idle) to 1.0 (fully busy)")
	fmt.Fprintln(w, "# TYPE fdb_process_disk_usage_ratio gauge")
	fmt.Fprintf(w, "fdb_process_disk_usage_ratio{id=\"%s\",address=\"%s\",class_type=\"%s\"} %f\n",
		id, p.Address, p.ClassType, d.Busy)

	fmt.Fprintln(w, "# HELP Amount of free disk space in bytes")
	fmt.Fprintln(w, "# TYPE fdb_process_disk_free_bytes gauge")
	fmt.Fprintf(w, "fdb_process_disk_free_bytes{id=\"%s\",address=\"%s\",class_type=\"%s\"} %d\n",
		id, p.Address, p.ClassType, d.FreeBytes)

	fmt.Fprintln(w, "# HELP Total amount of disk space in bytes")
	fmt.Fprintln(w, "# TYPE fdb_process_disk_total_bytes gauge")
	fmt.Fprintf(w, "fdb_process_disk_total_bytes{id=\"%s\",address=\"%s\",class_type=\"%s\"} %d\n",
		id, p.Address, p.ClassType, d.TotalBytes)

	fmt.Fprintln(w, "# HELP Count of disk read operations")
	fmt.Fprintln(w, "# TYPE fdb_process_disk_reads counter")
	fmt.Fprintf(w, "fdb_process_disk_reads{id=\"%s\",address=\"%s\",class_type=\"%s\"} %d\n",
		id, p.Address, p.ClassType, d.Reads.Counter)

	fmt.Fprintln(w, "# HELP Count of disk write operations")
	fmt.Fprintln(w, "# TYPE fdb_process_disk_writes counter")
	fmt.Fprintf(w, "fdb_process_disk_writes{id=\"%s\",address=\"%s\",class_type=\"%s\"} %d\n",
		id, p.Address, p.ClassType, d.Writes.Counter)

	return nil
}

type ProcessLocality struct {
	InstanceId string `json:"instance_id"`
	MachineId  string `json:"machineid"`
	ProcessId  string `json:"processid"`
	Process_Id string `json:"process_id"`
	ZoneId     string `json:"zoneid"`
}

type ProcessMemory struct {
	AvailableBytes        int `json:"available_bytes"`
	LimitBytes            int `json:"limit_bytes"`
	UnusedAllocatedMemory int `json:"unused_allocated_memory"`
	UsedBytes             int `json:"used_bytes"`
	RssBytes              int `json:"rss_bytes"`
}

func (m *ProcessMemory) DumpPromethusMetrics(w io.Writer, id string, p *Process) error {
	fmt.Fprintln(w, "# HELP Memory in bytes available to process")
	fmt.Fprintln(w, "# TYPE fdb_process_memory_available_bytes gauge")
	fmt.Fprintf(w, "fdb_process_memory_available_bytes{id=\"%s\",address=\"%s\",class_type=\"%s\"} %d\n",
		id, p.Address, p.ClassType, m.AvailableBytes)

	fmt.Fprintln(w, "# HELP Memory in bytes used by process")
	fmt.Fprintln(w, "# TYPE fdb_process_memory_used_bytes gauge")
	fmt.Fprintf(w, "fdb_process_memory_used_bytes{id=\"%s\",address=\"%s\",class_type=\"%s\"} %d\n",
		id, p.Address, p.ClassType, m.UsedBytes)

	fmt.Fprintln(w, "# HELP Memory limit in bytes for process")
	fmt.Fprintln(w, "# TYPE fdb_process_memory_limit_bytes gauge")
	fmt.Fprintf(w, "fdb_process_memory_limit_bytes{id=\"%s\",address=\"%s\",class_type=\"%s\"} %d\n",
		id, p.Address, p.ClassType, m.LimitBytes)

	fmt.Fprintln(w, "# HELP Unused memory in bytes allocated by process")
	fmt.Fprintln(w, "# TYPE fdb_process_memory_unused_allocated_memory gauge")
	fmt.Fprintf(w, "fdb_process_memory_unused_allocated_memory{id=\"%s\",address=\"%s\",class_type=\"%s\"} %d\n",
		id, p.Address, p.ClassType, m.UnusedAllocatedMemory)

	fmt.Fprintln(w, "# HELP Resident set size (RSS) memory in bytes for process")
	fmt.Fprintln(w, "# TYPE fdb_process_memory_rss_bytes gauge")
	fmt.Fprintf(w, "fdb_process_memory_rss_bytes{id=\"%s\",address=\"%s\",class_type=\"%s\"} %d\n",
		id, p.Address, p.ClassType, m.RssBytes)
	return nil
}

type ProcessMessage struct {
	Description   string  `json:"description"`
	Name          string  `json:"name"`
	RawLogMessage string  `json:"raw_log_message"`
	Time          float64 `json:"time"`
	Type          string  `json:"type"`
}

type LatencyStats struct {
	Count  int     `json:"count"`
	Max    float64 `json:"max"`
	Mean   float64 `json:"mean"`
	Median float64 `json:"median"`
	Min    float64 `json:"min"`
	P25    float64 `json:"p25"`
	P90    float64 `json:"p90"`
	P95    float64 `json:"p95"`
	P99    float64 `json:"p99"`
	P999   float64 `json:"p99.9"`
}

type ProcessNetwork struct {
	ConnectionErrors       Hz  `json:"connection_errors"`
	ConnectionsClosed      Hz  `json:"connections_closed"`
	ConnectionsEstablished Hz  `json:"connections_established"`
	CurrentConnections     int `json:"current_connections"`
	MegabitsReceived       Hz  `json:"megabits_received"`
	MegabitsSent           Hz  `json:"megabits_sent"`
	TlsPolicyFailures      Hz  `json:"tls_policy_failures"`
}

func (n *ProcessNetwork) DumpPromethusMetrics(w io.Writer, id string, p *Process) error {
	fmt.Fprintln(w, "# HELP Number of current connections to process")
	fmt.Fprintln(w, "# TYPE fdb_process_network_current_connections gauge")
	fmt.Fprintf(w, "fdb_process_network_current_connections{id=\"%s\",address=\"%s\",class_type=\"%s\"} %d\n",
		id, p.Address, p.ClassType, n.CurrentConnections)

	fmt.Fprintln(w, "# HELP Connection errors per second")
	fmt.Fprintln(w, "# TYPE fdb_process_network_connection_errors_rate gauge")
	fmt.Fprintf(w, "fdb_process_network_connection_errors_rate{id=\"%s\",address=\"%s\",class_type=\"%s\"} %f\n",
		id, p.Address, p.ClassType, n.ConnectionErrors.Hz)

	fmt.Fprintln(w, "# HELP Connections closed per second")
	fmt.Fprintln(w, "# TYPE fdb_process_network_connection_closed_rate gauge")
	fmt.Fprintf(w, "fdb_process_network_connection_closed_rate{id=\"%s\",address=\"%s\",class_type=\"%s\"} %f\n",
		id, p.Address, p.ClassType, n.ConnectionsClosed.Hz)

	fmt.Fprintln(w, "# HELP Connections established per second")
	fmt.Fprintln(w, "# TYPE fdb_process_network_connection_established_rate gauge")
	fmt.Fprintf(w, "fdb_process_network_connection_established_rate{id=\"%s\",address=\"%s\",class_type=\"%s\"} %f\n",
		id, p.Address, p.ClassType, n.ConnectionsEstablished.Hz)

	fmt.Fprintln(w, "# HELP Received data rate in megabits per second")
	fmt.Fprintln(w, "# TYPE fdb_process_network_megabits_received_rate gauge")
	fmt.Fprintf(w, "fdb_process_network_megabits_received_rate{id=\"%s\",address=\"%s\",class_type=\"%s\"} %f\n",
		id, p.Address, p.ClassType, n.MegabitsReceived.Hz)

	fmt.Fprintln(w, "# HELP Sent data rate in megabits per second")
	fmt.Fprintln(w, "# TYPE fdb_process_network_megabits_sent_rate gauge")
	fmt.Fprintf(w, "fdb_process_network_megabits_sent_rate{id=\"%s\",address=\"%s\",class_type=\"%s\"} %f\n",
		id, p.Address, p.ClassType, n.MegabitsSent.Hz)
	return nil
}

type GrvLatencyStats struct {
	Batch   LatencyStats `json:"batch"`
	Default LatencyStats `json:"default"`
}

type ProcessRole struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	// GRV proxy specific
	GrvLatencyStatistics GrvLatencyStats `json:"grv_latency_statistics"`
	// Commit Proxy specific
	CommitLatencyStatistics  LatencyStats `json:"commit_latency_statistics"`
	CommitBatchingWindowSize LatencyStats `json:"commit_batching_window_size"`
	// Storage and Log specific
	KvStoreAvailableBytes   int64           `json:"kvstore_available_bytes"`
	KvStoreFreeBytes        int64           `json:"kvstore_free_bytes"`
	KvStoreTotalBytes       int64           `json:"kvstore_total_bytes"`
	KvStoreUsedBytes        int64           `json:"kvstore_used_bytes"`
	QueueDiskAvailableBytes int64           `json:"queue_disk_available_bytes"`
	QueueDiskFreeBytes      int64           `json:"queue_disk_free_bytes"`
	QueueDiskTotalBytes     int64           `json:"queue_disk_total_bytes"`
	QueueDiskUsedBytes      int64           `json:"queue_disk_used_bytes"`
	DataVersion             int64           `json:"data_version"`
	DurableBytes            WorkloadMetrics `json:"durable_bytes"`
	InputBytes              WorkloadMetrics `json:"input_bytes"`
	// Only storage specific
	BytesQueried          WorkloadMetrics `json:"bytes_queried"`
	DataLag               Lag             `json:"data_lag"`
	DurabilityLag         Lag             `json:"durability_lag"`
	DurableVersion        int64           `json:"durable_version"`
	FetchedVersions       WorkloadMetrics `json:"fetched_versions"`
	FetchesFromLogs       WorkloadMetrics `json:"fetches_from_logs"`
	FinishedQueries       WorkloadMetrics `json:"finished_queries"`
	KeysQueried           WorkloadMetrics `json:"keys_queried"`
	LocalRate             int64           `json:"local_rate"`
	LowPriorityQueries    WorkloadMetrics `json:"low_priority_queries"`
	MutationBytes         WorkloadMetrics `json:"mutation_bytes"`
	Mutations             WorkloadMetrics `json:"mutations"`
	QueryQueueMax         int64           `json:"query_queue_max"`
	ReadLatencyStatistics LatencyStats    `json:"read_latency_statistics"`
	StoredBytes           int64           `json:"stored_bytes"`
	TotalQueries          WorkloadMetrics `json:"total_queries"`
}

func (r *ProcessRole) dumpStorageMetrics(w io.Writer, id string, p *Process) error {
	fmt.Fprintln(w, "# HELP Storage process data lag")
	fmt.Fprintln(w, "# TYPE fdb_process_storage_data_lag_seconds gauge")
	fmt.Fprintf(w, "fdb_process_storage_data_lag_seconds{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.DataLag.Seconds)

	fmt.Fprintln(w, "# HELP Storage process durability lag")
	fmt.Fprintln(w, "# TYPE fdb_process_storage_durability_lag_seconds gauge")
	fmt.Fprintf(w, "fdb_process_storage_durability_lag_seconds{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.DurabilityLag.Seconds)

	fmt.Fprintln(w, "# HELP Storage process input bytes")
	fmt.Fprintln(w, "# TYPE fdb_process_storage_input_bytes counter")
	fmt.Fprintf(w, "fdb_process_storage_input_bytes{id=\"%s\",address=\"%s\",role=\"%s\"} %d\n",
		id, p.Address, r.Role, r.InputBytes.Counter)

	fmt.Fprintln(w, "# HELP Storage process durable bytes")
	fmt.Fprintln(w, "# TYPE fdb_process_storage_durable_bytes counter")
	fmt.Fprintf(w, "fdb_process_storage_durable_bytes{id=\"%s\",address=\"%s\",role=\"%s\"} %d\n",
		id, p.Address, r.Role, r.DurableBytes.Counter)

	fmt.Fprintln(w, "# HELP Storage process stored bytes")
	fmt.Fprintln(w, "# TYPE fdb_process_storage_stored_bytes gauge")
	fmt.Fprintf(w, "fdb_process_storage_stored_bytes{id=\"%s\",address=\"%s\",role=\"%s\"} %d\n",
		id, p.Address, r.Role, r.StoredBytes)

	fmt.Fprintln(w, "# HELP Storage process KV store available bytes")
	fmt.Fprintln(w, "# TYPE fdb_process_storage_kvstore_available_bytes gauge")
	fmt.Fprintf(w, "fdb_process_storage_kvstore_available_bytes{id=\"%s\",address=\"%s\",role=\"%s\"} %d\n",
		id, p.Address, r.Role, r.KvStoreAvailableBytes)

	fmt.Fprintln(w, "# HELP Storage process KV store free bytes")
	fmt.Fprintln(w, "# TYPE fdb_process_storage_kvstore_free_bytes gauge")
	fmt.Fprintf(w, "fdb_process_storage_kvstore_free_bytes{id=\"%s\",address=\"%s\",role=\"%s\"} %d\n",
		id, p.Address, r.Role, r.KvStoreFreeBytes)

	fmt.Fprintln(w, "# HELP Storage process KV store total bytes")
	fmt.Fprintln(w, "# TYPE fdb_process_storage_kvstore_total_bytes gauge")
	fmt.Fprintf(w, "fdb_process_storage_kvstore_total_bytes{id=\"%s\",address=\"%s\",role=\"%s\"} %d\n",
		id, p.Address, r.Role, r.KvStoreTotalBytes)

	fmt.Fprintln(w, "# HELP Storage process KV store used bytes")
	fmt.Fprintln(w, "# TYPE fdb_process_storage_kvstore_used_bytes gauge")
	fmt.Fprintf(w, "fdb_process_storage_kvstore_used_bytes{id=\"%s\",address=\"%s\",role=\"%s\"} %d\n",
		id, p.Address, r.Role, r.KvStoreUsedBytes)

	fmt.Fprintln(w, "# HELP Storage process fetched versions count")
	fmt.Fprintln(w, "# TYPE fdb_process_storage_fetched_versions counter")
	fmt.Fprintf(w, "fdb_process_storage_fetched_versions{id=\"%s\",address=\"%s\",role=\"%s\"} %d\n",
		id, p.Address, r.Role, r.FetchedVersions.Counter)

	fmt.Fprintln(w, "# HELP Storage process fetches from logs count")
	fmt.Fprintln(w, "# TYPE fdb_process_storage_fetches_from_logs counter")
	fmt.Fprintf(w, "fdb_process_storage_fetches_from_logs{id=\"%s\",address=\"%s\",role=\"%s\"} %d\n",
		id, p.Address, r.Role, r.FetchesFromLogs.Counter)

	fmt.Fprintln(w, "# HELP Storage process low priority queries count")
	fmt.Fprintln(w, "# TYPE fdb_process_storage_low_priority_queries counter")
	fmt.Fprintf(w, "fdb_process_storage_low_priority_queries{id=\"%s\",address=\"%s\",role=\"%s\"} %d\n",
		id, p.Address, r.Role, r.LowPriorityQueries.Counter)

	fmt.Fprintln(w, "# HELP Storage process mean read latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_storage_read_latency_mean gauge")
	fmt.Fprintf(w, "fdb_process_storage_read_latency_mean{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.ReadLatencyStatistics.Mean)

	fmt.Fprintln(w, "# HELP Storage process median read latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_storage_read_latency_median gauge")
	fmt.Fprintf(w, "fdb_process_storage_read_latency_median{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.ReadLatencyStatistics.Median)

	fmt.Fprintln(w, "# HELP Storage process p95 read latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_storage_read_latency_p95 gauge")
	fmt.Fprintf(w, "fdb_process_storage_read_latency_p95{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.ReadLatencyStatistics.P95)

	fmt.Fprintln(w, "# HELP Storage process p99 read latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_storage_read_latency_p99 gauge")
	fmt.Fprintf(w, "fdb_process_storage_read_latency_p99{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.ReadLatencyStatistics.P99)

	fmt.Fprintln(w, "# HELP Storage process max read latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_storage_read_latency_max gauge")
	fmt.Fprintf(w, "fdb_process_storage_read_latency_max{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.ReadLatencyStatistics.Max)
	return nil
}

func (r *ProcessRole) dumpLogMetrics(w io.Writer, id string, p *Process) error {
	fmt.Fprintln(w, "# HELP Log process queue disk available bytes")
	fmt.Fprintln(w, "# TYPE fdb_process_log_queue_disk_available_bytes gauge")
	fmt.Fprintf(w, "fdb_process_log_queue_disk_available_bytes{id=\"%s\",address=\"%s\",role=\"%s\"} %d\n",
		id, p.Address, r.Role, r.QueueDiskAvailableBytes)

	fmt.Fprintln(w, "# HELP Log process queue disk free bytes")
	fmt.Fprintln(w, "# TYPE fdb_process_log_queue_disk_free_bytes gauge")
	fmt.Fprintf(w, "fdb_process_log_queue_disk_free_bytes{id=\"%s\",address=\"%s\",role=\"%s\"} %d\n",
		id, p.Address, r.Role, r.QueueDiskFreeBytes)

	fmt.Fprintln(w, "# HELP Log process queue disk total bytes")
	fmt.Fprintln(w, "# TYPE fdb_process_log_queue_disk_total_bytes gauge")
	fmt.Fprintf(w, "fdb_process_log_queue_disk_total_bytes{id=\"%s\",address=\"%s\",role=\"%s\"} %d\n",
		id, p.Address, r.Role, r.QueueDiskTotalBytes)

	fmt.Fprintln(w, "# HELP Log process queue disk used bytes")
	fmt.Fprintln(w, "# TYPE fdb_process_log_queue_disk_used_bytes gauge")
	fmt.Fprintf(w, "fdb_process_log_queue_disk_used_bytes{id=\"%s\",address=\"%s\",role=\"%s\"} %d\n",
		id, p.Address, r.Role, r.QueueDiskUsedBytes)

	fmt.Fprintln(w, "# HELP Log process input bytes")
	fmt.Fprintln(w, "# TYPE fdb_process_log_input_bytes counter")
	fmt.Fprintf(w, "fdb_process_log_input_bytes{id=\"%s\",address=\"%s\",role=\"%s\"} %d\n",
		id, p.Address, r.Role, r.InputBytes.Counter)

	fmt.Fprintln(w, "# HELP Log process durable bytes")
	fmt.Fprintln(w, "# TYPE fdb_process_log_durable_bytes counter")
	fmt.Fprintf(w, "fdb_process_log_durable_bytes{id=\"%s\",address=\"%s\",role=\"%s\"} %d\n",
		id, p.Address, r.Role, r.DurableBytes.Counter)
	return nil
}

func (r *ProcessRole) dumpCommitProxyMetrics(w io.Writer, id string, p *Process) error {
	fmt.Fprintln(w, "# HELP Commit proxy mean batching window size in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_commit_proxy_batching_window_size_mean gauge")
	fmt.Fprintf(w, "fdb_process_commit_proxy_batching_window_size_mean{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.CommitBatchingWindowSize.Mean)

	fmt.Fprintln(w, "# HELP Commit proxy median batching window size in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_commit_proxy_batching_window_size_median gauge")
	fmt.Fprintf(w, "fdb_process_commit_proxy_batching_window_size_median{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.CommitBatchingWindowSize.Median)

	fmt.Fprintln(w, "# HELP Commit proxy p95 batching window size in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_commit_proxy_batching_window_size_p95 gauge")
	fmt.Fprintf(w, "fdb_process_commit_proxy_batching_window_size_p95{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.CommitBatchingWindowSize.P95)

	fmt.Fprintln(w, "# HELP Commit proxy p99 batching window size in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_commit_proxy_batching_window_size_p99 gauge")
	fmt.Fprintf(w, "fdb_process_commit_proxy_batching_window_size_p99{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.CommitBatchingWindowSize.P99)

	fmt.Fprintln(w, "# HELP Commit proxy max batching window size in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_commit_proxy_batching_window_size_max gauge")
	fmt.Fprintf(w, "fdb_process_commit_proxy_batching_window_size_max{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.CommitBatchingWindowSize.Max)

	fmt.Fprintln(w, "# HELP Commit proxy mean commit latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_commit_proxy_commit_latency_mean gauge")
	fmt.Fprintf(w, "fdb_process_commit_proxy_commit_latency_mean{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.CommitLatencyStatistics.Mean)

	fmt.Fprintln(w, "# HELP Commit proxy median commit latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_commit_proxy_commit_latency_median gauge")
	fmt.Fprintf(w, "fdb_process_commit_proxy_commit_latency_median{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.CommitLatencyStatistics.Median)

	fmt.Fprintln(w, "# HELP Commit proxy p95 commit latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_commit_proxy_commit_latency_p95 gauge")
	fmt.Fprintf(w, "fdb_process_commit_proxy_commit_latency_p95{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.CommitLatencyStatistics.P95)

	fmt.Fprintln(w, "# HELP Commit proxy p99 commit latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_commit_proxy_commit_latency_p99 gauge")
	fmt.Fprintf(w, "fdb_process_commit_proxy_commit_latency_p99{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.CommitLatencyStatistics.P99)

	fmt.Fprintln(w, "# HELP Commit proxy max commit latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_commit_proxy_commit_latency_max gauge")
	fmt.Fprintf(w, "fdb_process_commit_proxy_commit_latency_max{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.CommitLatencyStatistics.Max)
	return nil
}

func (r *ProcessRole) dumpGRVProxyMetrics(w io.Writer, id string, p *Process) error {
	fmt.Fprintln(w, "# HELP GRV proxy mean batch priority latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_grv_proxy_batch_latency_mean gauge")
	fmt.Fprintf(w, "fdb_process_grv_proxy_batch_latency_mean{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.GrvLatencyStatistics.Batch.Mean)

	fmt.Fprintln(w, "# HELP GRV proxy median batch priority latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_grv_proxy_batch_latency_median gauge")
	fmt.Fprintf(w, "fdb_process_grv_proxy_batch_latency_median{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.GrvLatencyStatistics.Batch.Median)

	fmt.Fprintln(w, "# HELP GRV proxy p95 batch priority latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_grv_proxy_batch_latency_p95 gauge")
	fmt.Fprintf(w, "fdb_process_grv_proxy_batch_latency_p95{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.GrvLatencyStatistics.Batch.P95)

	fmt.Fprintln(w, "# HELP GRV proxy p99 batch priority latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_grv_proxy_batch_latency_p99 gauge")
	fmt.Fprintf(w, "fdb_process_grv_proxy_batch_latency_p99{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.GrvLatencyStatistics.Batch.P99)

	fmt.Fprintln(w, "# HELP GRV proxy mean default priority latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_grv_proxy_default_latency_mean gauge")
	fmt.Fprintf(w, "fdb_process_grv_proxy_default_latency_mean{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.GrvLatencyStatistics.Batch.Mean)

	fmt.Fprintln(w, "# HELP GRV proxy median default priority latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_grv_proxy_default_latency_median gauge")
	fmt.Fprintf(w, "fdb_process_grv_proxy_default_latency_median{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.GrvLatencyStatistics.Batch.Median)

	fmt.Fprintln(w, "# HELP GRV proxy p95 default priority latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_grv_proxy_default_latency_p95 gauge")
	fmt.Fprintf(w, "fdb_process_grv_proxy_default_latency_p95{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.GrvLatencyStatistics.Batch.P95)

	fmt.Fprintln(w, "# HELP GRV proxy p99 default priority latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_grv_proxy_default_latency_p99 gauge")
	fmt.Fprintf(w, "fdb_process_grv_proxy_default_latency_p99{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.GrvLatencyStatistics.Batch.P99)

	fmt.Fprintln(w, "# HELP GRV proxy max default priority latency in seconds")
	fmt.Fprintln(w, "# TYPE fdb_process_grv_proxy_default_latency_max gauge")
	fmt.Fprintf(w, "fdb_process_grv_proxy_default_latency_max{id=\"%s\",address=\"%s\",role=\"%s\"} %f\n",
		id, p.Address, r.Role, r.GrvLatencyStatistics.Batch.Max)

	return nil
}

func (r *ProcessRole) DumpPromethusMetrics(w io.Writer, id string, p *Process) error {
	switch r.Role {
	case "storage":
		r.dumpStorageMetrics(w, id, p)
	case "log":
		r.dumpLogMetrics(w, id, p)
	case "commit_proxy":
		r.dumpCommitProxyMetrics(w, id, p)
	case "grv_proxy":
		r.dumpGRVProxyMetrics(w, id, p)
	}
	return nil
}

type Process struct {
	Address       string           `json:"address"`
	ClassSource   string           `json:"class_source"`
	ClassType     string           `json:"class_type"`
	CommandLine   string           `json:"command_line"`
	UptimeSeconds float64          `json:"uptime_seconds"`
	Cpu           ProcessCpu       `json:"cpu"`
	Disk          ProcessDisk      `json:"disk"`
	Excluded      bool             `json:"excluded"`
	Degraded      bool             `json:"degraded"`
	FaultDomain   string           `json:"fault_domain"`
	Locality      ProcessLocality  `json:"locality"`
	MachineId     string           `json:"machine_id"`
	Memory        ProcessMemory    `json:"memory"`
	Messages      []ProcessMessage `json:"messages"`
	Network       ProcessNetwork   `json:"network"`
	Roles         []ProcessRole    `json:"roles"`
	RunLoopBusy   float64          `json:"run_loop_busy"`
}

func (p *Process) DumpPromethusMetrics(w io.Writer, id string) error {
	fmt.Fprintln(w, "# HELP Process uptime")
	fmt.Fprintln(w, "# TYPE fdb_process_uptime_seconds gauge")
	fmt.Fprintf(w, "fdb_process_uptime_seconds{id=\"%s\",address=\"%s\",class_type=\"%s\"} %f\n",
		id, p.Address, p.ClassType, p.UptimeSeconds)

	fmt.Fprintln(w, "# HELP Fraction of time the run loop was busy")
	fmt.Fprintln(w, "# TYPE fdb_process_run_loop_busy_ratio gauge")
	fmt.Fprintf(w, "fdb_process_run_loop_busy_ratio{id=\"%s\",address=\"%s\",class_type=\"%s\"} %f\n",
		id, p.Address, p.ClassType, p.RunLoopBusy)

	fmt.Fprintln(w, "# HELP Amount of CPU cores used by process")
	fmt.Fprintln(w, "# TYPE fdb_process_cpu_usage_cores gauge")
	fmt.Fprintf(w, "fdb_process_cpu_usage_cores{id=\"%s\",address=\"%s\",class_type=\"%s\"} %f\n",
		id, p.Address, p.ClassType, p.Cpu.UsageCores)

	p.Memory.DumpPromethusMetrics(w, id, p)
	p.Disk.DumpPromethusMetrics(w, id, p)
	p.Network.DumpPromethusMetrics(w, id, p)
	for _, role := range p.Roles {
		role.DumpPromethusMetrics(w, id, p)
	}
	return nil
}
