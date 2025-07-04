// Package install /*
//
// Contains the actual logic for installing
// the GraphDB Platform independent distribution file.
package install

import (
	"fmt"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tool_configurations/statics"
	"graphdbcli/internal/tui/versiontui"
	"github.com/enescakir/emoji"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

// installSelectedVersion install the provided version, that must follow the
// rules for Semantic versioning.
func installSelectedVersion(selectedVersion string) {
	versionIndex := statics.GetVersionIndex(selectedVersion)
	if versionIndex < 0 {
		fmt.Printf("version %s not found. The latest available version: %s\n", selectedVersion, statics.Versions[0].Version)
		logging.LOGGER.Fatal("unavailable versionw was specified",
			zap.String("version", selectedVersion))
	}

	version := statics.Versions[versionIndex]

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: Unable to get user home directory:", err)
		return
	}

	gdbCliDir := filepath.Join(homeDir, statics.HomeDirSpaceName)
	distDir := filepath.Join(gdbCliDir, statics.DistDirName)
	zipFile := filepath.Join(distDir, fmt.Sprintf("graphdb-%s.zip", selectedVersion))

	if _, err := os.Stat(zipFile); os.IsNotExist(err) {
		versiontui.DownloadWithProgressBar(selectedVersion, version.Url, zipFile)
		if versiontui.IsFileDownloadedSucucessful() {
			fmt.Printf("%s GraphDB version %s has been downloaded!\n", emoji.Rocket, selectedVersion)
		} else {
			fmt.Printf("%s GraphDB installation has been cancelled!\n", emoji.StopSign)
			os.Remove(zipFile)
		}

	} else {
		fmt.Printf("%s GraphDB version %s is already present!\n", emoji.BeachWithUmbrella, selectedVersion)
	}
}
