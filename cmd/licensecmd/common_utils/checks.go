// Package common_utils provides common utility functions for license management.
package common_utils

import (
	"fmt"
	"graphdbcli/internal/tool_configurations/logging"
	tc "graphdbcli/internal/tool_configurations/statics"
	"os"
	"path"
	"path/filepath"

	"go.uber.org/zap"
)

// CheckDoesLicensesDirExists checks if the directory for storing licences exists.
// It returns true if the directory exists, otherwise it logs an error and returns false.
// It is used to ensure that the directory used for managing licenses is accessible.
func CheckDoesLicensesDirExists() bool {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Unable to get user home directory")
		logging.LOGGER.Fatal("Unable to get user home directory:",
			zap.Error(err))
	}

	logging.LOGGER.Debug("Home directory: " + homeDir)

	gdbCliDir := filepath.Join(homeDir, tc.HomeDirSpaceName)
	licensesDir := filepath.Join(gdbCliDir, tc.LicensesDirName)

	_, err = os.Stat(licensesDir)
	if err != nil {
		fmt.Println(logging.ErrorMessages[003].External)
		logging.LOGGER.Fatal("The licenses directory does not exist",
			zap.String("licenses directory", licensesDir),
			zap.Error(err),
		)
		return false
	}

	logging.LOGGER.Info("The licenses directory exists",
		zap.String("licenses directory", licensesDir))
	return true
}

// CheckDoesLicenseFileExists checks if the provided license file exists.
// It returns true if the file exists, otherwise it logs an error and returns false.
// It is used to ensure that the file provided by the user is valid before attempting to store it.
func CheckDoesLicenseFileExists(licenseFilename string) bool {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Unable to get user home directory")
		logging.LOGGER.Fatal("Unable to get user home directory:", zap.Error(err))
	}

	licenseFilepath := path.Join(homeDir, tc.HomeDirSpaceName, tc.LicensesDirName, licenseFilename)
	_, err = os.Stat(licenseFilepath)
	if err == nil {
		logging.LOGGER.Info("The provided file does exist",
			zap.String("file", licenseFilepath))
		return true
	}

	logging.LOGGER.Info("The provided license file was not found",
		zap.String("file", licenseFilepath))
	return false
}
