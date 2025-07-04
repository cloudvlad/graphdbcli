package graphdb_cluster

import (
	"fmt"
	ini "graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	tc "graphdbcli/internal/tool_configurations/statics"
	"graphdbcli/internal/tui/common_components"
	"io"
	"os"
	"path/filepath"

	"github.com/enescakir/emoji"
	"go.uber.org/zap"
)

// ReadPropertyOverrides reads the provided property file by using the provided filepath
// returns the content of the file as string
func ReadPropertyOverrides() string {
	logging.LOGGER.Info("Reading property overrides...")

	if PropertyOverridesFilepath == "" {
		fmt.Printf("%s Property overrides file not specified. Skipping...\n",
			common_components.PadStatusIndicator(string(emoji.NextTrackButton), 0))
		logging.LOGGER.Info("Property overrides file not specified. Skipping...")
		return ""
	}
	_, err := os.Stat(PropertyOverridesFilepath)
	if err != nil {
		fmt.Printf("%s An error was encountered while trying to read the property overrides",
			common_components.PadStatusIndicator(string(emoji.CrossMark), tc.NotTUIStatusIndicatorAdditionalPadding))
		logging.LOGGER.Fatal("Error was encountered whilst reading property overrides file",
			zap.Error(err))
	}

	overridesFile, err := os.Open(PropertyOverridesFilepath)
	if err != nil {
		fmt.Println("Error encountered while trying to access property overrides file")
		logging.LOGGER.Fatal("Error was encountered while trying to open property overrides")
	}

	defer func(overridesFile *os.File) {
		err := overridesFile.Close()
		if err != nil {
			logging.LOGGER.Debug("Error was encountered whilst closing overrides file",
				zap.Error(err))
		}
	}(overridesFile)

	var rawData []byte
	rawData, err = io.ReadAll(overridesFile)
	if err != nil {
		fmt.Println("Error reading property overrides")
		logging.LOGGER.Fatal("Error was encountered while reading property overrides file")
	}

	return string(rawData)
}

// CheckDoesInstancesDirExists checks if the directory for the instances exists
// returns true if it does, if does not - exit with non-zero status
func CheckDoesInstancesDirExists() bool {
	homeDir := ini.GetUserHomeDirectory()
	gdbCliDir := filepath.Join(homeDir, tc.HomeDirSpaceName)
	clustersDir := filepath.Join(gdbCliDir, tc.InstancesDirName)

	if _, err := os.Stat(clustersDir); os.IsNotExist(err) {
		logging.LOGGER.Fatal("The instances directory does not exists",
			zap.String("instances directory", clustersDir))
	}

	logging.LOGGER.Info("The clusters directory exists",
		zap.String("instances directory", clustersDir))
	return true
}
