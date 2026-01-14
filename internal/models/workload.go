package models

import (
	"fmt"
	"io"
)

type WorkloadMetrics struct {
	Counter   int64   `json:"counter"`
	Hz        float64 `json:"hz"`
	Roughness float64 `json:"roughness"`
}

type WorkloadBytes struct {
	Read    WorkloadMetrics `json:"read"`
	Written WorkloadMetrics `json:"written"`
}

type WorkloadKeys struct {
	Read WorkloadMetrics `json:"read"`
}

type WorkloadOperations struct {
	LocationRequests WorkloadMetrics `json:"location_requests"`
	LowPriorityReads WorkloadMetrics `json:"low_priority_reads"`
	MemoryErrors     WorkloadMetrics `json:"memory_errors"`
	ReadRequests     WorkloadMetrics `json:"read_requests"`
	Reads            WorkloadMetrics `json:"reads"`
	Writes           WorkloadMetrics `json:"writes"`
}

func (o *WorkloadOperations) DumpPromethusMetrics(w io.Writer) error {
	fmt.Fprintln(w, "# HELP Count of workload read request operations")
	fmt.Fprintln(w, "# TYPE fdb_cluster_workload_operations_read_requests counter")
	fmt.Fprintf(w, "fdb_cluster_workload_operations_read_requests %d\n", o.ReadRequests.Counter)

	fmt.Fprintln(w, "# HELP Count of workload read operations")
	fmt.Fprintln(w, "# TYPE fdb_cluster_workload_operations_reads counter")
	fmt.Fprintf(w, "fdb_cluster_workload_operations_reads %d\n", o.Reads.Counter)

	fmt.Fprintln(w, "# HELP Count of workload write operations")
	fmt.Fprintln(w, "# TYPE fdb_cluster_workload_operations_writes counter")
	fmt.Fprintf(w, "fdb_cluster_workload_operations_writes %d\n", o.Writes.Counter)

	fmt.Fprintln(w, "# HELP Count of workload location request operations")
	fmt.Fprintln(w, "# TYPE fdb_cluster_workload_operations_location_requests counter")
	fmt.Fprintf(w, "fdb_cluster_workload_operations_location_requests %d\n", o.LocationRequests.Counter)

	fmt.Fprintln(w, "# HELP Count of workload low priority read operations")
	fmt.Fprintln(w, "# TYPE fdb_cluster_workload_operations_low_priority_reads counter")
	fmt.Fprintf(w, "fdb_cluster_workload_operations_low_priority_reads %d\n", o.LowPriorityReads.Counter)

	fmt.Fprintln(w, "# HELP Count of workload memory errors")
	fmt.Fprintln(w, "# TYPE fdb_cluster_workload_operations_memory_errors counter")
	fmt.Fprintf(w, "fdb_cluster_workload_operations_memory_errors %d\n", o.MemoryErrors.Counter)
	return nil
}

type WorkloadTransactions struct {
	Committed                WorkloadMetrics `json:"committed"`
	Conflicted               WorkloadMetrics `json:"conflicted"`
	RejectedForQueuedTooLong WorkloadMetrics `json:"rejected_for_queued_too_long"`
	Started                  WorkloadMetrics `json:"started"`
	StartedBatchPriority     WorkloadMetrics `json:"started_batch_priority"`
	StartedDefaultPriority   WorkloadMetrics `json:"started_default_priority"`
	StartedImmediatePriority WorkloadMetrics `json:"started_immediate_priority"`
}

func (t *WorkloadTransactions) DumpPromethusMetrics(w io.Writer) error {
	fmt.Fprintln(w, "# HELP Count of workload started transactions")
	fmt.Fprintln(w, "# TYPE fdb_cluster_workload_transactions_started counter")
	fmt.Fprintf(w, "fdb_cluster_workload_transactions_started %d\n", t.Started.Counter)

	fmt.Fprintln(w, "# HELP Count of workload committed transactions")
	fmt.Fprintln(w, "# TYPE fdb_cluster_workload_transactions_committed counter")
	fmt.Fprintf(w, "fdb_cluster_workload_transactions_committed %d\n", t.Committed.Counter)

	fmt.Fprintln(w, "# HELP Count of workload conflicted transactions")
	fmt.Fprintln(w, "# TYPE fdb_cluster_workload_transactions_conflicted counter")
	fmt.Fprintf(w, "fdb_cluster_workload_transactions_conflicted %d\n", t.Conflicted.Counter)

	fmt.Fprintln(w, "# HELP Count of workload transactions rejected for being queued too long")
	fmt.Fprintln(w, "# TYPE fdb_cluster_workload_transactions_rejected_for_queued_too_long counter")
	fmt.Fprintf(w, "fdb_cluster_workload_transactions_rejected_for_queued_too_long %d\n", t.RejectedForQueuedTooLong.Counter)
	return nil
}

type Workload struct {
	Bytes        WorkloadBytes        `json:"bytes"`
	Keys         WorkloadKeys         `json:"keys"`
	Operations   WorkloadOperations   `json:"operations"`
	Transactions WorkloadTransactions `json:"transactions"`
}

func (l *Workload) DumpPromethusMetrics(w io.Writer) error {
	l.Operations.DumpPromethusMetrics(w)
	l.Transactions.DumpPromethusMetrics(w)

	fmt.Fprintln(w, "# HELP Count of workload keys read")
	fmt.Fprintln(w, "# TYPE fdb_cluster_workload_keys_read counter")
	fmt.Fprintf(w, "fdb_cluster_workload_keys_read %d\n", l.Keys.Read.Counter)

	fmt.Fprintln(w, "# HELP Count of workload bytes read")
	fmt.Fprintln(w, "# TYPE fdb_cluster_workload_bytes_read counter")
	fmt.Fprintf(w, "fdb_cluster_workload_bytes_read %d\n", l.Bytes.Read.Counter)
	fmt.Fprintln(w, "# HELP Count of workload bytes written")
	fmt.Fprintln(w, "# TYPE fdb_cluster_workload_bytes_written counter")
	fmt.Fprintf(w, "fdb_cluster_workload_bytes_written %d\n", l.Bytes.Written.Counter)

	return nil
}
