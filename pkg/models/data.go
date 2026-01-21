package models

import (
	"fmt"
	"io"
)

type MovingData struct {
	HighestPriority   int64 `json:"highest_priority"`
	InFlightBytes     int64 `json:"in_flight_bytes"`
	InQueueBytes      int64 `json:"in_queue_bytes"`
	TotalWrittenBytes int64 `json:"total_written_bytes"`
}

type State struct {
	Description          string `json:"description"`
	Healthy              bool   `json:"healthy"`
	MinReplicasRemaining int64  `json:"min_replicas_remaining"`
	Name                 string `json:"name"`
}

type TeamTracker struct {
	InFlightBytes    int64  `json:"in_flight_bytes"`
	Primary          bool   `json:"primary"`
	State            *State `json:"state"`
	UnhealthyServers int64  `json:"unhealthy_servers"`
}

type Data struct {
	AveragePartitionSizeBytes             int64         `json:"average_partition_size_bytes"`
	LeastOperatingSpaceBytesLogServer     int64         `json:"least_operating_space_bytes_log_server"`
	LeastOperatingSpaceBytesStorageServer int64         `json:"least_operating_space_bytes_storage_server"`
	MovingData                            MovingData    `json:"moving_data"`
	PartitionsCount                       int64         `json:"partitions_count"`
	State                                 State         `json:"state"`
	SystemKvSizeBytes                     int64         `json:"system_kv_size_bytes"`
	TeamTrackers                          []TeamTracker `json:"team_trackers"`
	TotalDiskUsedBytes                    int64         `json:"total_disk_used_bytes"`
	TotalKvSizeBytes                      int64         `json:"total_kv_size_bytes"`
}

func (d *Data) DumpPromethusMetrics(w io.Writer) error {
	fmt.Fprintln(w, "# HELP Indicates data state")
	fmt.Fprintln(w, "# TYPE fdb_cluster_data_state gauge")
	fmt.Fprintf(w, "fdb_cluster_data_state{state=\"%s\"} 1\n", d.State.Name)

	fmt.Fprintln(w, "# HELP Operating space on most full log server")
	fmt.Fprintln(w, "# TYPE fdb_cluster_data_least_operating_space_log_server_bytes gauge")
	fmt.Fprintf(w, "fdb_cluster_data_least_operating_space_log_server_bytes %d\n", d.LeastOperatingSpaceBytesLogServer)

	fmt.Fprintln(w, "# HELP Operating space on most full storage server")
	fmt.Fprintln(w, "# TYPE fdb_cluster_data_least_operating_space_storage_server_bytes gauge")
	fmt.Fprintf(w, "fdb_cluster_data_least_operating_space_storage_server_bytes %d\n", d.LeastOperatingSpaceBytesStorageServer)

	fmt.Fprintln(w, "# HELP Average partition size")
	fmt.Fprintln(w, "# TYPE fdb_cluster_data_average_partition_size_bytes gauge")
	fmt.Fprintf(w, "fdb_cluster_data_average_partition_size_bytes %d\n", d.AveragePartitionSizeBytes)

	fmt.Fprintln(w, "# HELP Average partition count")
	fmt.Fprintln(w, "# TYPE fdb_cluster_data_partitions gauge")
	fmt.Fprintf(w, "fdb_cluster_data_partitions %d\n", d.PartitionsCount)

	fmt.Fprintln(w, "# HELP Total disk used")
	fmt.Fprintln(w, "# TYPE fdb_cluster_data_total_disk_used_bytes gauge")
	fmt.Fprintf(w, "fdb_cluster_data_total_disk_used_bytes %d\n", d.TotalDiskUsedBytes)

	fmt.Fprintln(w, "# HELP Total KV size")
	fmt.Fprintln(w, "# TYPE fdb_cluster_data_total_kv_size_bytes gauge")
	fmt.Fprintf(w, "fdb_cluster_data_total_kv_size_bytes %d\n", d.TotalKvSizeBytes)

	fmt.Fprintln(w, "# HELP System KV size")
	fmt.Fprintln(w, "# TYPE fdb_cluster_data_system_kv_size_bytes gauge")
	fmt.Fprintf(w, "fdb_cluster_data_system_kv_size_bytes %d\n", d.SystemKvSizeBytes)

	fmt.Fprintln(w, "# HELP Moving data in-flight bytes")
	fmt.Fprintln(w, "# TYPE fdb_cluster_data_moving_in_flight_bytes gauge")
	fmt.Fprintf(w, "fdb_cluster_data_moving_in_flight_byte %d\n", d.MovingData.InFlightBytes)

	fmt.Fprintln(w, "# HELP Moving data in-queue bytes")
	fmt.Fprintln(w, "# TYPE fdb_cluster_data_moving_in_queue_bytes gauge")
	fmt.Fprintf(w, "fdb_cluster_data_moving_in_queue_bytes %d\n", d.MovingData.InQueueBytes)

	fmt.Fprintln(w, "# HELP Moving data total written bytes")
	fmt.Fprintln(w, "# TYPE fdb_cluster_data_moving_total_written_bytes counter")
	fmt.Fprintf(w, "fdb_cluster_data_moving_total_written_bytes %d\n", d.MovingData.TotalWrittenBytes)

	return nil
}
