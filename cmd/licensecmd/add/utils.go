// Package add /*
//
// Responsible for the logging
// regarding adding licenses.
package add

import (
	"fmt"
	cu "graphdbcli/cmd/licensecmd/common_utils"
	"graphdbcli/internal/tool_configurations/logging"
	tc "graphdbcli/internal/tool_configurations/statics"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
)

// LicenseFileNoteContent is a globally used varibale, mainly set by the
// CLI on execution by providing the -n/--note flag.
var LicenseFileNoteContent string

// StoreLicenseFile is more abstract method, containing additional
// safety checks so on.
func StoreLicenseFile(licenseFilepath string) {
	if cu.CheckDoesLicenseFileExists(licenseFilepath) &&
		cu.CheckDoesLicensesDirExists() {
		storeProvidedLicense(licenseFilepath)
	} else {
		fmt.Println("Internal error occurred while store license file")
		logging.LOGGER.Fatal("There was ",
			zap.String("licenseFilepath", licenseFilepath),
		)
	}
}

// storeProvidedLicense saves the license by copying the file located at the
// provided path into the directory dedicated for storing license files.
func storeProvidedLicense(licenseFilepath string) {
	fileInformation, _ := os.Stat(licenseFilepath)

	// cannot copy non-regular files (e.g., directories,
	// symlinks, devices, etc.)
	if !fileInformation.Mode().IsRegular() {
		fmt.Println(logging.ErrorMessages[002].External)
		logging.LOGGER.Fatal(logging.ErrorMessages[002].Internal,
			zap.String("licenseFilepath", licenseFilepath))
	}

	originalLicenseFile, err := os.Open(licenseFilepath)
	defer originalLicenseFile.Close()
	if err != nil {
		fmt.Println("There was problem opening the provided license file")
		logging.LOGGER.Fatal("There was problem opening the provided license file",
			zap.String("licenseFilepath", licenseFilepath),
			zap.Error(err),
		)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(logging.ErrorMessages[004].External)
		logging.LOGGER.Fatal(logging.ErrorMessages[004].Internal,
			zap.Error(err),
			zap.String("licenseFilepath", licenseFilepath),
		)
	}

	homeDirSpace := tc.HomeDirSpaceName
	licensesDirName := tc.LicensesDirName

	gdbDir := filepath.Join(homeDir, homeDirSpace)
	licensesDir := filepath.Join(gdbDir, licensesDirName)
	licenseFilename := filepath.Base(licenseFilepath)
	copiedLicensePath := filepath.Join(licensesDir, licenseFilename)

	storedLicenseFile, err := os.Create(copiedLicensePath)
	defer storedLicenseFile.Close()
	if err != nil {
		fmt.Println(logging.ErrorMessages[005].External)
		logging.LOGGER.Fatal(logging.ErrorMessages[005].Internal,
			zap.String("licensesDirName", licensesDirName),
			zap.Error(err),
		)
	}

	copiedLicenseNotePath := copiedLicensePath + ".note"
	storedLicenseFileNote, err := os.Create(copiedLicenseNotePath)
	defer storedLicenseFileNote.Close()
	if err != nil {
		fmt.Println(logging.ErrorMessages[006].External)
		logging.LOGGER.Fatal(logging.ErrorMessages[006].Internal,
			zap.String("licenseNoteFilename", copiedLicenseNotePath),
			zap.Error(err),
		)
	}

	_, err = io.Copy(storedLicenseFile, originalLicenseFile)
	if err != nil {
		fmt.Println(logging.ErrorMessages[006].External)
		logging.LOGGER.Fatal(logging.ErrorMessages[005].Internal,
			zap.Error(err),
		)
	}

	err = storedLicenseFile.Sync()
	if err != nil {
		fmt.Println("Error occurred while storing the note for the license")
		logging.LOGGER.Fatal("The licenses directory could not be opened and the note file for the license was not stored",
			zap.Error(err),
		)
	}

	_, err = io.WriteString(storedLicenseFileNote, LicenseFileNoteContent)
	if err != nil {
		fmt.Println("Problem occurred while writing note for license")
		logging.LOGGER.Fatal("Problem occurred while writing note for license",
			zap.Error(err),
		)
	}

	err = storedLicenseFileNote.Sync()
	if err != nil {
		logging.LOGGER.Fatal("Problem occurred while storing note for a license file",
			zap.Error(err),
		)
	}

	logging.LOGGER.Info("The license was successfully stored!")
}
