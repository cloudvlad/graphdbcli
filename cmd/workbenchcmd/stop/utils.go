package stop

import (
	"fmt"
	ini "graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tui/common_components"
	"os"
	"path"
	"regexp"
	"strconv"
	"syscall"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/enescakir/emoji"
	"go.uber.org/zap"
)

var p *tea.Program

var isRegexPattern bool
var forceStop bool

// stopInstance manages stopping one or more instances. If isRegexPattern is true,
// it will stop all instances matching the regex pattern in instanceName.
func stopInstance(instanceName string) {
	workbenchesDir := ini.GetWorkbenchDirectory()

	if !isRegexPattern {
		instancePath := path.Join(workbenchesDir, instanceName)
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
	instances, err := os.ReadDir(workbenchesDir)
	if err != nil {
		logging.LOGGER.Fatal("Failed to read clusters directory",
			zap.String("workbenchesDir", workbenchesDir),
			zap.Error(err))
	}
	for _, instance := range instances {
		name := instance.Name()
		matched, err := regexp.MatchString(instanceName, name)
		if err != nil {
			fmt.Printf("%s Error matching regex pattern", err)
			logging.LOGGER.Error("Failed to match regex pattern", zap.Error(err))
			continue
		}
		if matched {
			logging.LOGGER.Debug("Stopping instance", zap.String("name", name))
			instancePath := path.Join(workbenchesDir, name)
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
		fmt.Printf("%s Invalid PID file content", common_components.PadStatusIndicator(emoji.CrossMark.String(), 0))
		logging.LOGGER.Fatal("invalid PID file content", zap.String("path file", pidFilePath), zap.Error(err))
	}

	var signal syscall.Signal

	if forceStop {
		signal = syscall.SIGKILL
	} else {
		signal = syscall.SIGTERM
	}

	err = syscall.Kill(pid, signal)
	if err != nil {
		fmt.Printf("%s Failed to terminate process with PID %d", common_components.PadStatusIndicator(emoji.CrossMark.String(), 0), pid)
		logging.LOGGER.Fatal("Failed to terminate process", zap.Error(err), zap.Int("pid", pid))
	}

	fmt.Printf("%s Successfully initiated instance shutdown for %s\n",
		common_components.PadStatusIndicator(emoji.CheckMark.String(), 0),
		instanceName)
	logging.LOGGER.Info("successfully initiated instance shutdown", zap.String("instance name", instanceName))
}
