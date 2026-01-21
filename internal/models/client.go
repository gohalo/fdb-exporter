package models

import (
	"fmt"
	"io"
)

type ClusterFile struct {
	Path     string `json:"path"`
	UpToDate bool   `json:"up_to_date"`
}

type Coordinator struct {
	Address   string `json:"address"`
	Reachable bool   `json:"reachable"`
	Protocol  string `json:"protocol"`
}

type Coordinators struct {
	Coordinators    []Coordinator `json:"coordinators"`
	QuorumReachable bool          `json:"quorum_reachable"`
}

type DatabaseStatus struct {
	Available bool `json:"available"`
	Healthy   bool `json:"healthy"`
}

type ClientMessage struct {
	// Possible names
	// "inconsistent_cluster_file",
	// "unreachable_cluster_controller",
	// "no_cluster_controller",
	// "status_incomplete_client",
	// "status_incomplete_coordinators",
	// "status_incomplete_error",
	// "status_incomplete_timeout",
	// "status_incomplete_cluster",
	// "quorum_not_reachable"
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Client top level status
type ClientStatus struct {
	ClusterFile    ClusterFile     `json:"cluster_file"`
	Coordinators   Coordinators    `json:"coordinators"`
	DatabaseStatus DatabaseStatus  `json:"database_status"`
	Messages       []ClientMessage `json:"messages"`
	Timestamp      int64           `json:"timestamp"`
}

func (c *ClientStatus) IsValid() bool {
	return c.Coordinators.QuorumReachable
}

func (c *ClientStatus) DumpPromethusMetrics(w io.Writer) error {
	fmt.Fprintln(w, "# HELP Whether or not database is available. 1 if is, 0 otherwise")
	fmt.Fprintln(w, "# TYPE fdb_client_database_status_available gauge")
	if c.DatabaseStatus.Available {
		fmt.Fprintln(w, "fdb_client_database_status_available 1")
	} else {
		fmt.Fprintln(w, "fdb_client_database_status_available 0")
	}

	fmt.Fprintln(w, "# HELP The number of coordinators")
	fmt.Fprintln(w, "# TYPE fdb_client_coordinators gauge")
	fmt.Fprintf(w, "fdb_client_coordinators %d\n", len(c.Coordinators.Coordinators))
	return nil
}
