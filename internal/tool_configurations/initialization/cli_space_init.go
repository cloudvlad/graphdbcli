// Package initialization /*
//
// The package is used for initializing the tool and create
// the needed directories
package initialization

import (
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tool_configurations/statics"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

// InitializeCLIHomeDirectory Initializes all needed directories for the tool
func InitializeCLIHomeDirectory() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		errorField := zap.Field{Key: err.Error()}
		logging.LOGGER.Error("Error: Unable to get user home directory:", errorField)
		return
	}

	gdbDir := filepath.Join(homeDir, statics.HomeDirSpaceName)
	distDir := filepath.Join(gdbDir, statics.DistDirName)
	logsDir := filepath.Join(gdbDir, statics.LogsDirName)
	licensesDir := filepath.Join(gdbDir, statics.LicensesDirName)

	directories := []string{gdbDir, distDir, logsDir, licensesDir}
	for _, directory := range directories {
		if _, err = os.Stat(directory); os.IsNotExist(err) {
			err = os.Mkdir(directory, 0750)
			if err != nil {
				gdbDirField := zap.Field{Key: gdbDir}
				errorField := zap.Field{Key: err.Error()}
				logging.LOGGER.Error("Error: Unable to create %s directory: %s", gdbDirField, errorField)
				return
			}
		}
	}
}
