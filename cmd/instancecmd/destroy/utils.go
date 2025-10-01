package destroy

import (
	"context"
	"fmt"
	"graphdbcli/internal/tool_configurations/initialization"
	ini "graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	tc "graphdbcli/internal/tool_configurations/statics"
	"graphdbcli/internal/tui/common_components"
	"os"
	"path"
	"regexp"

	"github.com/enescakir/emoji"
	"go.uber.org/zap"
)

var isRegexPattern bool

// destroy stop running instances, if there are such and deletes
// the instance designated directory
func destroy(context context.Context, ctxCancel context.CancelFunc, instanceName string) {
	if !isRegexPattern {
		destroyInstance(instanceName)
		return
	}

	instancesPath := initialization.GetClustersDirectory()
	instances, err := os.ReadDir(instancesPath)
	if err != nil {
		fmt.Println("Error reading clusters directory:", err)
		os.Exit(1)
	}

	for _, instance := range instances {
		name := instance.Name()
		matched, err := regexp.MatchString(instanceName, name)
		if err != nil {
			fmt.Printf("Error matching regex pattern: %v\n", err)
			continue
		}
		if matched {
			destroyInstance(name)
		}
	}
}

func destroyInstance(instanceName string) {
	instancePath := path.Join(ini.GetClustersDirectory(), instanceName)
	if !isInstancePresent(instancePath) {
		fmt.Printf("%s Instance %s not found. Exiting... \n",
			common_components.PadStatusIndicator(string(emoji.FastForwardButton), tc.NotTUIStatusIndicatorAdditionalPadding), instanceName)
	}

	pidFile := path.Join(instancePath, ".instance_pid")
	if isPidFilePresent(pidFile) {
		fmt.Printf("%s Instance %s appears to be running. Skipping... \n",
			common_components.PadStatusIndicator(emoji.Warning.String(), 0), instanceName)
		return
	}

	// Remove the instance directory and all its contents
	err := os.RemoveAll(instancePath)
	if err != nil {
		return
	}

	fmt.Printf("%s Instance %s removed \n",
		common_components.PadStatusIndicator(emoji.CheckMark.String(), 0),
		instanceName)
	logging.LOGGER.Info("instance destroyed", zap.String("instance name", instanceName))
}

func isInstancePresent(instancePath string) bool {
	if _, err := os.Stat(instancePath); err == nil {
		return true
	}
	return false
}

// isPidFilePresent checks if the pid file exists.
// Returns true if present, false otherwise.
func isPidFilePresent(pidFile string) bool {
	if _, err := os.Stat(pidFile); err == nil {
		return true
	}
	return false
}
