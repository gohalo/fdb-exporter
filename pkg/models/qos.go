package models

import (
	"fmt"
	"io"
)

type PerformanceLimitedBy struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	ReasonId    int64  `json:"reason_id"`
}

type ThrottledTags struct {
	Auto   AutoThrottledTags   `json:"auto"`
	Manual ManualThrottledTags `json:"manual"`
}

type AutoThrottledTags struct {
	BusyRead        int64 `json:"busy_read"`
	BusyWrite       int64 `json:"busy_write"`
	Count           int64 `json:"count"`
	RecommendedOnly int64 `json:"recommended_only"`
}

type ManualThrottledTags struct {
	Count int64 `json:"count"`
}

type Qos struct {
	BatchPerformanceLimitedBy          PerformanceLimitedBy `json:"batch_performance_limited_by"`
	BatchReleasedTransactionsPerSecond float64              `json:"batch_released_transactions_per_second"`
	BatchTransactionsPerSecondLimit    float64              `json:"batch_transactions_per_second_limit"`
	LimitingDataLagStorageServer       Lag                  `json:"limiting_data_lag_storage_server"`
	LimitingDurabilityLagStorageServer Lag                  `json:"limiting_durability_lag_storage_server"`
	LimitingQueueBytesStorageServer    int64                `json:"limiting_queue_bytes_storage_server"`
	PerformanceLimitedBy               PerformanceLimitedBy `json:"performance_limited_by"`
	ReleasedTransactionsPerSecond      float64              `json:"released_transactions_per_second"`
	ThrottledTags                      ThrottledTags        `json:"throttled_tags"`
	TransactionsPerSecondLimit         float64              `json:"transactions_per_second_limit"`
	WorstDataLagStorageServer          Lag                  `json:"worst_data_lag_storage_server"`
	WorstDurabilityLagStorageServer    Lag                  `json:"worst_durability_lag_storage_server"`
	WorstQueueBytesLogServer           int64                `json:"worst_queue_bytes_log_server"`
	WorstQueueBytesStorageServer       int64                `json:"worst_queue_bytes_storage_server"`
}

func (q *Qos) DumpPromethusMetrics(w io.Writer) error {
	fmt.Fprintln(w, "# HELP QoS limiting data lag among storage servers")
	fmt.Fprintln(w, "# TYPE fdb_cluster_qos_limiting_storage_server_data_lag_seconds gauge")
	fmt.Fprintf(w, "fdb_cluster_qos_limiting_storage_server_data_lag_seconds %f\n", q.LimitingDataLagStorageServer.Seconds)

	fmt.Fprintln(w, "# HELP QoS limiting durability lag among storage servers")
	fmt.Fprintln(w, "# TYPE fdb_cluster_qos_limiting_storage_server_durability_lag_seconds gauge")
	fmt.Fprintf(w, "fdb_cluster_qos_limiting_storage_server_durability_lag_seconds %f\n", q.LimitingDurabilityLagStorageServer.Seconds)

	fmt.Fprintln(w, "# HELP QoS limiting queue bytes among storage servers")
	fmt.Fprintln(w, "# TYPE fdb_cluster_qos_limiting_storage_server_queue_bytes gauge")
	fmt.Fprintf(w, "fdb_cluster_qos_limiting_storage_server_queue_bytes %d\n", q.LimitingQueueBytesStorageServer)

	fmt.Fprintln(w, "# HELP QoS worst data lag among storage servers")
	fmt.Fprintln(w, "# TYPE fdb_cluster_qos_worst_storage_server_data_lag_seconds gauge")
	fmt.Fprintf(w, "fdb_cluster_qos_worst_storage_server_data_lag_seconds %f\n", q.WorstDataLagStorageServer.Seconds)

	fmt.Fprintln(w, "# HELP QoS worst durability lag among storage servers")
	fmt.Fprintln(w, "# TYPE fdb_cluster_qos_worst_storage_server_durability_lag_seconds gauge")
	fmt.Fprintf(w, "fdb_cluster_qos_worst_storage_server_durability_lag_seconds %f\n", q.WorstDurabilityLagStorageServer.Seconds)

	fmt.Fprintln(w, "# HELP Qos worst queue bytes among storage servers")
	fmt.Fprintln(w, "# TYPE fdb_cluster_qos_worst_storage_server_queue_bytes gauge")
	fmt.Fprintf(w, "fdb_cluster_qos_worst_storage_server_queue_bytes %d\n", q.WorstQueueBytesStorageServer)

	fmt.Fprintln(w, "# HELP Qos worst queue bytes among log servers")
	fmt.Fprintln(w, "# TYPE fdb_cluster_qos_worst_log_server_queue_bytes gauge")
	fmt.Fprintf(w, "fdb_cluster_qos_worst_log_server_queue_bytes %d\n", q.WorstQueueBytesLogServer)

	fmt.Fprintln(w, "# HELP QoS released transactions per second")
	fmt.Fprintln(w, "# TYPE fdb_cluster_qos_released_transactions_per_second gauge")
	fmt.Fprintf(w, "fdb_cluster_qos_released_transactions_per_second %f\n", q.ReleasedTransactionsPerSecond)

	fmt.Fprintln(w, "# HELP QoS transactions per second limit")
	fmt.Fprintln(w, "# TYPE fdb_cluster_qos_transactions_per_second_limit gauge")
	fmt.Fprintf(w, "fdb_cluster_qos_transactions_per_second_limit %f\n", q.TransactionsPerSecondLimit)

	fmt.Fprintln(w, "# HELP Indicates the reason for limited performance")
	fmt.Fprintln(w, "# TYPE fdb_cluster_qos_performance_limited_by gauge")
	fmt.Fprintf(w, "fdb_cluster_qos_performance_limited_by{name=\"%s\",description=\"%s\"} 1\n",
		q.PerformanceLimitedBy.Name, q.PerformanceLimitedBy.Description)

	fmt.Fprintln(w, "# HELP QoS released transactions per second")
	fmt.Fprintln(w, "# TYPE fdb_cluster_qos_batch_released_transactions_per_second gauge")
	fmt.Fprintf(w, "fdb_cluster_qos_batch_released_transactions_per_second %f\n", q.BatchReleasedTransactionsPerSecond)

	fmt.Fprintln(w, "# HELP QoS transactions per second limit")
	fmt.Fprintln(w, "# TYPE fdb_cluster_qos_batch_transactions_per_second_limit gauge")
	fmt.Fprintf(w, "fdb_cluster_qos_batch_transactions_per_second_limit %f\n", q.BatchTransactionsPerSecondLimit)

	fmt.Fprintln(w, "# HELP Indicates the reason for limited performance")
	fmt.Fprintln(w, "# TYPE fdb_cluster_qos_batch_performance_limited_by gauge")
	fmt.Fprintf(w, "fdb_cluster_qos_batch_performance_limited_by{name=\"%s\",description=\"%s\"} 1\n",
		q.BatchPerformanceLimitedBy.Name, q.BatchPerformanceLimitedBy.Description)

	fmt.Fprintln(w, "# HELP Number of automatically throttled tags")
	fmt.Fprintln(w, "# TYPE fdb_cluster_qos_throttled_tags_auto_count gauge")
	fmt.Fprintf(w, "fdb_cluster_qos_throttled_tags_auto_count %d\n", q.ThrottledTags.Auto.Count)

	fmt.Fprintln(w, "# HELP Number of automatically throttled tags for busy reads")
	fmt.Fprintln(w, "# TYPE fdb_cluster_qos_throttled_tags_auto_busy_read gauge")
	fmt.Fprintf(w, "fdb_cluster_qos_throttled_tags_auto_busy_read %d\n", q.ThrottledTags.Auto.BusyRead)

	fmt.Fprintln(w, "# HELP Number of automatically throttled tags for busy writes")
	fmt.Fprintln(w, "# TYPE fdb_cluster_qos_throttled_tags_auto_busy_write gauge")
	fmt.Fprintf(w, "fdb_cluster_qos_throttled_tags_auto_busy_write %d\n", q.ThrottledTags.Auto.BusyWrite)

	fmt.Fprintln(w, "# HELP Number of manually throttled tags")
	fmt.Fprintln(w, "# TYPE fdb_cluster_qos_throttled_tags_manual_count gauge")
	fmt.Fprintf(w, "fdb_cluster_qos_throttled_tags_manual_count %d\n", q.ThrottledTags.Manual.Count)
	return nil
}
