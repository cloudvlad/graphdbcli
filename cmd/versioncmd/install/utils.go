// Package install /*
//
// Contains the actual logic for installing
// the GraphDB Platform independent distribution file.
package install

import (
	"crypto/sha256"
	"fmt"
	"graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tool_configurations/statics"
	"graphdbcli/internal/tui/common_components"
	"graphdbcli/internal/tui/versiontui"
	"io"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	"github.com/enescakir/emoji"
	"go.uber.org/zap"
)

// installSelectedVersion install the provided version, that must follow the
// rules for Semantic versioning.
func installSelectedVersion(selectedVersion string) {
	distDir := initialization.GetDistDirectory()
	zipFile := filepath.Join(distDir, fmt.Sprintf("graphdb-%s.zip", selectedVersion))

	if _, err := os.Stat(zipFile); os.IsNotExist(err) {
		url := fmt.Sprintf("https://maven.ontotext.com/repository/owlim-releases/com/ontotext/graphdb/graphdb/%s/graphdb-%s-dist.zip", selectedVersion, selectedVersion)

		versiontui.DownloadWithProgressBar(selectedVersion, url, zipFile)

		if versiontui.IsFileDownloadedSuccessful {
			fmt.Printf("%s GraphDB version %s has been downloaded!\n", common_components.PadStatusIndicator(emoji.Rocket.String(), 0), selectedVersion)
		} else {
			fmt.Printf("%s GraphDB installation has been cancelled!\n", common_components.PadStatusIndicator(emoji.StopSign.String(), 0))
			err = os.Remove(zipFile)
			if err != nil {
				fmt.Printf("Error removing file %s: %s\n", zipFile, err)
				logging.LOGGER.Fatal("Error occured whilst removing the zip file", zap.Error(err), zap.String("zipFile", zipFile))
				return
			}
		}
	} else {
		fmt.Printf("%s GraphDB version %s is already present!\n", common_components.PadStatusIndicator(emoji.BeachWithUmbrella.String(), 0), selectedVersion)
	}
}

// integrityCheck checks if the file was tempered with by
// comparing the previously calculated and saved hash
// with the one cacluated after downloading it
func integrityCheck(version statics.Version) {
	distDir := initialization.GetDistDirectory()
	zipFile := filepath.Join(distDir, fmt.Sprintf("graphdb-%s.zip", version.Version))

	file, err := os.Open(zipFile)
	if err != nil {
		logging.LOGGER.Fatal("failed to open file for integrity check", zap.Error(err))
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		logging.LOGGER.Fatal("failed to hash file", zap.Error(err))
	}
	calculatedSum := fmt.Sprintf("%x", hasher.Sum(nil))

	if calculatedSum != version.Sha256sum {
		fmt.Printf("%s Integrity check failed\n", common_components.PadStatusIndicator(emoji.CrossMark.String(), 0))
		logging.LOGGER.Fatal("integrity check failed: hash mismatch",
			zap.String("expected", version.Sha256sum),
			zap.String("actual", calculatedSum))
	} else {
		fmt.Printf("%s Integrity check passed\n", common_components.PadStatusIndicator(emoji.CheckMark.String(), 0))
		logging.LOGGER.Info("integrity check passed")
	}
}

// TuiListAvailableVersions show the available
// versions in TUI mode
func TuiListAvailableVersions() {
	var items []list.Item

	for _, version := range statics.Versions {
		item := versiontui.Item{
			Version:     version.Version,
			ReleaseDate: version.ReleaseDate,
		}

		items = append(items, item)
	}

	p := versiontui.Initialize(items)
	if _, err := p.Run(); err != nil {
		logging.LOGGER.Error("Error running program: " + err.Error())
		os.Exit(1)
	}
}
