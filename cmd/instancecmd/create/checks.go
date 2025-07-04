package create

import (
	c "graphdbcli/internal/data_objects/graphdb_cluster"
	"graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	"net"
	"os"
	"path"
)

// isPortOpen check if the provided port is open or occupied by another service.
func isPortOpen(port string) bool {
	host := ":" + port
	server, err := net.Listen("tcp", host)

	if err != nil {
		return false
	}
	server.Close()

	return true
}

func IsClusterPresent() bool {
	fullClusterPath := path.Join(initialization.GetClustersDirectory(), c.Instance.Name)

	_, err := os.Stat(fullClusterPath)
	if err == nil {
		logging.LOGGER.Error("The cluster is already present")
		return true
	}

	return false
}
