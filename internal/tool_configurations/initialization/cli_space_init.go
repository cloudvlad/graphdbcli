// Package initialization /*
//
// The package is used for initializing the tool and create
// the needed directories
package initialization

import (
	"go.uber.org/zap"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tool_configurations/statics"
	"os"
	"path/filepath"
)

// InitializeCLIHomeDirectory Initializes all needed directories for the tool
func InitializeCLIHomeDirectory() {
	homeDir := GetUserHomeDirectory()
	gdbDir := filepath.Join(homeDir, statics.HomeDirSpaceName)
	distDir := filepath.Join(gdbDir, statics.DistDirName)
	logsDir := filepath.Join(gdbDir, statics.LogsDirName)
	licensesDir := filepath.Join(gdbDir, statics.LicensesDirName)
	clustersDir := filepath.Join(gdbDir, statics.InstancesDirName)
	resourcesDir := filepath.Join(gdbDir, statics.ResourcesDirName)

	directories := []string{gdbDir, distDir, logsDir, licensesDir, clustersDir, resourcesDir}
	for _, directory := range directories {
		if _, err := os.Stat(directory); os.IsNotExist(err) {
			err = os.Mkdir(directory, 0750)
			if err != nil {
				logging.LOGGER.Fatal("Error: Unable to create graphdb internal directory",
					zap.String("tried directory", gdbDir),
					zap.Error(err))
			}
		}
	}
}
