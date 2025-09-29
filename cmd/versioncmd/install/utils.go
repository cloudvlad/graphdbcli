// Package install /*
//
// Contains the actual logic for installing
// the GraphDB Platform independent distribution file.
package install

import (
	"context"
	"crypto/sha256"
	"fmt"
	"graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tool_configurations/statics"
	"graphdbcli/internal/tui/versiontui"
	"io"
	"os"
	"path/filepath"

	"github.com/enescakir/emoji"
	"go.uber.org/zap"
)

var isIntegrityCheckNeeded = false

// installSelectedVersion install the provided version, that must follow the
// rules for Semantic versioning.
func installSelectedVersion(version statics.Version) {
	distDir := initialization.GetDistDirectory()
	zipFile := filepath.Join(distDir, fmt.Sprintf("graphdb-%s.zip", version.Version))

	if _, err := os.Stat(zipFile); os.IsNotExist(err) {
		url := fmt.Sprintf("https://ipfs.io/ipfs/%s", version.IpfsCID)
		if statics.IsTuiDisabled {
			versiontui.NotTUIinstallVersion(context.Background(), version)
		} else {
			versiontui.DownloadWithProgressBar(version.Version, url, zipFile)
		}

		if versiontui.IsFileDownloadedSuccessful {
			fmt.Printf("%s GraphDB version %s has been downloaded!\n", emoji.Rocket, version.Version)
		} else {
			fmt.Printf("%s GraphDB installation has been cancelled!\n", emoji.StopSign)
			err = os.Remove(zipFile)
			if err != nil {
				fmt.Printf("Error removing file %s: %s\n", zipFile, err)
				logging.LOGGER.Fatal("Error occured whilst removing the zip file", zap.Error(err), zap.String("zipFile", zipFile))
				return
			}
		}
	} else {
		fmt.Printf("%s GraphDB version %s is already present!\n", emoji.BeachWithUmbrella, version.Version)
	}
}

// integrityCheck checks if the file was tempered with by
// comparing the previously calculated and saved hash
// with the one cacluated after downloading it
func integrityCheck(version statics.Version) {
	if !isIntegrityCheckNeeded {
		logging.LOGGER.Debug("Skipping integrity check...")
		return
	}

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
		fmt.Printf("%s Integrity check failed\n", emoji.CrossMark)
		logging.LOGGER.Fatal("integrity check failed: hash mismatch",
			zap.String("expected", version.Sha256sum),
			zap.String("actual", calculatedSum))
	} else {
		fmt.Printf("%s Integrity check passed\n", emoji.CheckMark)
		logging.LOGGER.Info("integrity check passed")
	}
}
