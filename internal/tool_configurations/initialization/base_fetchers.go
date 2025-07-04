package initialization

import (
	"fmt"
	"go.uber.org/zap"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tool_configurations/statics"
	"os"
	"path/filepath"
)

func GetUserHomeDirectory() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Unable to get user home directory")
		logging.LOGGER.Fatal("Unable to get user home directory:",
			zap.Error(err))
	}
	logging.LOGGER.Debug("Home directory: " + homeDir)

	return homeDir
}

func GetDistDirectory() string {
	homeDir := GetUserHomeDirectory()
	gdbDir := filepath.Join(homeDir, statics.HomeDirSpaceName)
	distDir := filepath.Join(gdbDir, statics.DistDirName)

	return distDir
}

func GetClustersDirectory() string {
	homeDir := GetUserHomeDirectory()
	gdbDir := filepath.Join(homeDir, statics.HomeDirSpaceName)
	clustersDir := filepath.Join(gdbDir, statics.InstancesDirName)

	return clustersDir
}

func GetLicensesDirectory() string {
	homeDir := GetUserHomeDirectory()
	gdbDir := filepath.Join(homeDir, statics.HomeDirSpaceName)
	licensesDir := filepath.Join(gdbDir, statics.LicensesDirName)

	return licensesDir
}

func GetResourcesDirectory() string {
	homeDir := GetUserHomeDirectory()
	gdbDir := filepath.Join(homeDir, statics.HomeDirSpaceName)
	resourcesDir := filepath.Join(gdbDir, statics.ResourcesDirName)

	return resourcesDir
}
