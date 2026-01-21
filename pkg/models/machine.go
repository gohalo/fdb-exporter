package models

import (
	"fmt"
	"io"
)

type MachineCpu struct {
	LogicalCoreUtilization float64 `json:"logical_core_utilization"`
}

type MachineLocality struct {
	MachineId string `json:"machineid"`
	ProcessId string `json:"processid"`
	ZoneId    string `json:"zoneid"`
}

type MachineMemory struct {
	CommittedBytes int `json:"committed_bytes"`
	FreeBytes      int `json:"free_bytes"`
	TotalBytes     int `json:"total_bytes"`
}

type MachineNetwork struct {
	MegabitsReceived         Hz `json:"megabits_received"`
	MegabitsSent             Hz `json:"megabits_sent"`
	TcpSegmentsRetransmitted Hz `json:"tcp_segments_retransmitted"`
}

type Machine struct {
	Address             string          `json:"address"`
	MachineID           string          `json:"machine_id"`
	ContributingWorkers int             `json:"contributing_workers"`
	Cpu                 MachineCpu      `json:"cpu"`
	Excluded            bool            `json:"excluded"`
	Locality            MachineLocality `json:"locality"`
	Memory              MachineMemory   `json:"memory"`
	Network             MachineNetwork  `json:"network"`
}

func (m *Machine) DumpPromethusMetrics(w io.Writer) error {
	fmt.Fprintln(w, "# HELP Workers on this machine")
	fmt.Fprintln(w, "# TYPE fdb_machine_workers gauge")
	fmt.Fprintf(w, "fdb_machine_workers{id=\"%s\",address=\"%s\"} %d\n",
		m.MachineID, m.Address, m.ContributingWorkers)

	fmt.Fprintln(w, "# HELP Logical core utilization")
	fmt.Fprintln(w, "# TYPE fdb_machine_logical_core_utils gauge")
	fmt.Fprintf(w, "fdb_machine_logical_core_utils{id=\"%s\",address=\"%s\"} %f\n",
		m.MachineID, m.Address, m.Cpu.LogicalCoreUtilization)

	fmt.Fprintln(w, "# HELP Total memory in bytes")
	fmt.Fprintln(w, "# TYPE fdb_machine_memory_total_bytes gauge")
	fmt.Fprintf(w, "fdb_machine_memory_total_bytes{id=\"%s\",address=\"%s\"} %d\n",
		m.MachineID, m.Address, m.Memory.TotalBytes)

	fmt.Fprintln(w, "# HELP Free memory in bytes")
	fmt.Fprintln(w, "# TYPE fdb_machine_memory_free_bytes gauge")
	fmt.Fprintf(w, "fdb_machine_memory_free_bytes{id=\"%s\",address=\"%s\"} %d\n",
		m.MachineID, m.Address, m.Memory.FreeBytes)

	fmt.Fprintln(w, "# HELP Committed memory in bytes")
	fmt.Fprintln(w, "# TYPE fdb_machine_memory_committed_bytes gauge")
	fmt.Fprintf(w, "fdb_machine_memory_committed_bytes{id=\"%s\",address=\"%s\"} %d\n",
		m.MachineID, m.Address, m.Memory.CommittedBytes)

	fmt.Fprintln(w, "# HELP Received data rate in megabits per second")
	fmt.Fprintln(w, "# TYPE fdb_machine_network_megabits_received_rate gauge")
	fmt.Fprintf(w, "fdb_machine_network_megabits_received_rate{id=\"%s\",address=\"%s\"} %f\n",
		m.MachineID, m.Address, m.Network.MegabitsReceived.Hz)

	fmt.Fprintln(w, "# HELP Sent data rate in megabits per second")
	fmt.Fprintln(w, "# TYPE fdb_machine_network_megabits_sent_rate gauge")
	fmt.Fprintf(w, "fdb_machine_network_megabits_sent_rate{id=\"%s\",address=\"%s\"} %f\n",
		m.MachineID, m.Address, m.Network.MegabitsSent.Hz)

	fmt.Fprintln(w, "# HELP Retransmitted TCP segments per second")
	fmt.Fprintln(w, "# TYPE fdb_machine_tcp_segements_retransmitted gauge")
	fmt.Fprintf(w, "fdb_machine_tcp_segements_retransmitted{id=\"%s\",address=\"%s\"} %f\n",
		m.MachineID, m.Address, m.Network.TcpSegmentsRetransmitted.Hz)
	return nil
}
