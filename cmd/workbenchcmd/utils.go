package workbenchcmd

import (
	"fmt"
	wbmd "graphdbcli/internal/data_objects/workbench_metadata"
	"graphdbcli/internal/tool_configurations/logging"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"syscall"

	"graphdbcli/internal/tool_configurations/initialization"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

// updateWorkbenchesMetadata is used to update the metadata and the process id file for every instance
func updateWorkbenchesMetadata() {
	workbenchesPath := initialization.GetWorkbenchesDirectory()
	workbenches, err := os.ReadDir(workbenchesPath)
	if err != nil {
		fmt.Println("Error reading clusters directory:", err)
		return
	}

	for _, workbench := range workbenches {
		if !workbench.IsDir() {
			continue
		}

		workbenchDir := path.Join(workbenchesPath, workbench.Name())
		pidFile := path.Join(workbenchDir, ".workbench_pid")
		metadataFile := path.Join(workbenchDir, "metadata.yaml")

		// Read metadata
		metaBytes, err := os.ReadFile(metadataFile)
		if err != nil {
			continue
		}
		var meta wbmd.Data
		if err := yaml.Unmarshal(metaBytes, &meta); err != nil {
			continue
		}

		pidRunning := updatePidPresence(pidFile)

		// Update status logic
		if pidRunning && meta.Status != "Active" {
			meta.Status = "Active"
		} else if !pidRunning && meta.Status != "Inactive" {
			meta.Status = "Inactive"
		}

		gdbHost, gdbPort, wbPort := getLatestConfigurations(workbenchDir)

		meta.WorkbenchPort = wbPort
		meta.GraphDBHost = gdbHost
		meta.GraphDBPort = gdbPort

		// Write back if changed
		newMetaBytes, err := yaml.Marshal(&meta)
		if err == nil {
			_ = os.WriteFile(metadataFile, newMetaBytes, 0644)
		}
	}
}

// updatePidPresence sync the file presence of the pid file (.pid_instance) in the instance directory
// if the file is present, but the process is no longer running (due to reboot or manual shutdown)
func updatePidPresence(pidFile string) bool {
	pidBytes, err := os.ReadFile(pidFile)
	pidExists := err == nil
	pidRunning := false
	if pidExists {
		pidStr := string(pidBytes)
		pid, err := strconv.Atoi(pidStr)
		if err == nil {
			// Check if process is running
			process, err := os.FindProcess(pid)
			if err == nil {
				// On Unix, sending signal 0 checks if process exists
				if err := process.Signal(syscall.Signal(0)); err == nil {
					pidRunning = true
				} else {
					os.Remove(pidFile)
				}
			}
		}
	}

	return pidRunning
}

func getLatestConfigurations(workbenchPath string) (string, string, string) {
	var graphdbHost string
	var graphdbPort string
	var workbenchPort string

	webpackFile := filepath.Join(workbenchPath, "webpack.config.dev.js")

	data, err := os.ReadFile(webpackFile)
	if err != nil {
		logging.LOGGER.Debug("webpack config not found or unreadable", zap.Error(err))
		return "", "", ""
	}

	content := string(data)

	// Each property appears only once, so capture group 1 contains the value
	reHost := regexp.MustCompile(`(?m)const\s+host\s*=\s*['"](.*?)['"];?`)
	rePortHere := regexp.MustCompile(`(?m)const\s+portHere\s*=\s*(\d+);?`)
	rePortThere := regexp.MustCompile(`(?m)const\s+portThere\s*=\s*(\d+);?`)

	if m := reHost.FindStringSubmatch(content); m != nil && len(m) > 1 {
		graphdbHost = m[1]
	}

	if m := rePortThere.FindStringSubmatch(content); m != nil && len(m) > 1 {
		graphdbPort = m[1]
	}

	if m := rePortHere.FindStringSubmatch(content); m != nil && len(m) > 1 {
		workbenchPort = m[1]
	}

	return graphdbHost, graphdbPort, workbenchPort
}
