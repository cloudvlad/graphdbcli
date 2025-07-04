// Package remove /*
//
// Responsible for the removing of the stored
// licenses logic.
package remove

import (
	"fmt"
	cu "graphdbcli/cmd/licensecmd/common_utils"
	"graphdbcli/internal/tool_configurations/logging"
	tc "graphdbcli/internal/tool_configurations/statics"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

// RemoveLicenseFile is responsible for the deletion of a stored
// license and the related information with it.
func RemoveLicenseFile(licenseFilepath string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logging.LOGGER.Fatal("Unable to get user home directory:",
			zap.Error(err))
		fmt.Println("Unable to get user home directory")
	}

	gdbDir := filepath.Join(homeDir, tc.HomeDirSpaceName)
	licensesDir := filepath.Join(gdbDir, tc.LicensesDirName)
	licenseFilename := filepath.Base(licenseFilepath)
	storedLicensePath := filepath.Join(licensesDir, licenseFilename)

	storedLicenseFile, err := os.Create(storedLicensePath)
	defer storedLicenseFile.Close()
	if err != nil {
		fmt.Println("Error: Unable to create license file", tc.LicensesDirName)
		logging.LOGGER.Fatal("The licenses directory could not be opened",
			zap.Error(err),
			zap.String("licensesDirName", tc.LicensesDirName),
		)
	}

	if !cu.CheckDoesLicensesDirExists() || !cu.CheckDoesLicenseFileExists(licenseFilepath) {
		fmt.Println("Error: License file or licenses directory does not exist.")
		return
	}

	// Removes the actual stored license
	err = os.Remove(storedLicensePath)
	if err != nil {
		fmt.Println("Error: Unable to delete license file.")
		logging.LOGGER.Fatal("The license file could not be deleted",
			zap.Error(err),
			zap.String("licenseFilepath", storedLicensePath),
		)
	}

	// Removes the license note
	storedLicenseNotePath := storedLicensePath + ".note"
	err = os.Remove(storedLicenseNotePath)
	if err != nil {
		fmt.Println("Warning: Unable to delete license note file.")
		logging.LOGGER.Warn("The license note file could not be deleted",
			zap.Error(err),
			zap.String("licenseNotePath", storedLicenseNotePath),
		)
	}

	logging.LOGGER.Info("The license file and its note were successfully deleted!")
}
