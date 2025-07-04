package instancecmd

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"syscall"

	"graphdbcli/internal/data_objects/intance_metadata"
	"graphdbcli/internal/tool_configurations/initialization"

	"gopkg.in/yaml.v3"
)

// updateInstancesMetadata is used to update the metadata and the process id file for every instance
func updateInstancesMetadata() {
	instancesPath := initialization.GetClustersDirectory()
	instances, err := os.ReadDir(instancesPath)
	if err != nil {
		fmt.Println("Error reading clusters directory:", err)
		return
	}

	for _, instance := range instances {
		if !instance.IsDir() {
			continue
		}
		instanceDir := path.Join(instancesPath, instance.Name())
		pidFile := path.Join(instanceDir, ".instance_pid")
		metadataFile := path.Join(instanceDir, "metadata.yaml")

		// Read metadata
		metaBytes, err := os.ReadFile(metadataFile)
		if err != nil {
			continue
		}
		var meta intance_metadata.InstanceMetadata
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
