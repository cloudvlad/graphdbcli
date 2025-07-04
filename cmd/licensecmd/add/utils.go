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
	"io"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

// LicenseFileNoteContent is a globally used varibale, mainly set by the
// CLI on execution by providing the -n/--note flag.
var LicenseFileNoteContent string

// NewLicenseFileName specified the new name of the license file after storing it
// CLI on exection by providing the -r/--rename flag.
var NewLicenseFileName string

// StoreLicenseFile is more abstract method, containing additional
// safety checks so on.
func StoreLicenseFile(licenseFilepath string) {
	if cu.CheckDoesLicensesDirExists() {
		licenseInfo, err := os.Stat(licenseFilepath)
		if err != nil {
			fmt.Println("The provided license does not exists")
			logging.LOGGER.Fatal("the provided license does not exists", zap.Error(err))
		}

		if !licenseInfo.Mode().IsRegular() {
			fmt.Println("The provided license file should be a regular file")
			logging.LOGGER.Fatal("the provided license file was not a regular file")
		}

		licenseFilename := filepath.Base(licenseFilepath)
		if !cu.CheckDoesLicenseFileExists(licenseFilename) {
			storeProvidedLicense(licenseFilepath)
		}

	} else {
		fmt.Println("The licenses directory doesn't exists.")
		logging.LOGGER.Fatal("The licenses directory doesn't exists.")
	}
}

// storeProvidedLicense saves the license by copying the file located at the
// provided path into the directory dedicated for storing license files.
func storeProvidedLicense(licenseFilepath string) {
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
	if NewLicenseFileName != "" {
		licenseFilename = NewLicenseFileName
	}

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
