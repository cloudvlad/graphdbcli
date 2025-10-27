package stop

import (
	"fmt"
	ini "graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tui/common_components"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"syscall"

	"github.com/enescakir/emoji"
	"go.uber.org/zap"
)

var instanceName string

var forceStop bool
var isRegexPattern bool

// stopInstance manages stopping one or more instances. If isRegexPattern is true,
// it will stop all instances matching the regex pattern in instanceName.
func stopInstance() {
	clustersDir := ini.GetWorkbenchesDirectory()
	if !isRegexPattern {
		instancePath := path.Join(clustersDir, instanceName)
		_, err := os.Stat(instancePath)
		if err != nil {
			logging.LOGGER.Fatal("Failed to read instance directory",
				zap.String("instancePath", instancePath),
				zap.Error(err))
		}
		stopProcess(instancePath, instanceName)
		return
	}

	// Regex pattern: stop all matching instances
	instances, err := os.ReadDir(clustersDir)
	if err != nil {
		logging.LOGGER.Fatal("Failed to read clusters directory",
			zap.String("clustersDir", clustersDir),
			zap.Error(err))
	}
	for _, instance := range instances {
		name := instance.Name()
		matched, err := regexp.MatchString(instanceName, name)
		if err != nil {
			log.Printf("Error matching regex pattern: %v", err)
			continue
		}
		if matched {
			instancePath := path.Join(clustersDir, name)
			stopProcess(instancePath, name)
		}
	}
}

// stopProcess accepts the instance path and there it checks the
// PID file, that was previously created. If still running, send the
// appropriate signal for shutting it down.
func stopProcess(instancePath, instanceName string) {
	logging.LOGGER.Info("Stopping instance", zap.String("instance name", instanceName), zap.Bool("forceStop", forceStop))
	pidFilePath := path.Join(instancePath, ".instance_pid")

	// Read the PID from the .instance_pid file
	pidData, err := os.ReadFile(pidFilePath)
	if err != nil {
		logging.LOGGER.Fatal("Failed to read PID file",
			zap.String("pidFile", pidFilePath),
			zap.Error(err))
	}

	// Convert the PID to an integer
	pid, err := strconv.Atoi(string(pidData))
	if err != nil {
		log.Printf("Invalid PID in file %s: %v", pidFilePath, err)
		return
	}

	var signal syscall.Signal

	if forceStop {
		signal = syscall.SIGKILL
	} else {
		signal = syscall.SIGTERM
	}

	err = syscall.Kill(pid, signal)
	if err != nil {
		log.Printf("Failed to terminate process with PID %d: %v", pid, err)
		return
	}

	fmt.Printf("%s Successfully initiated instance shutdown for %s\n",
		common_components.PadStatusIndicator(emoji.CheckMark.String(), 0),
		instanceName)
	logging.LOGGER.Info("successfully initiated instance shutdown", zap.String("instance name", instanceName))
}
