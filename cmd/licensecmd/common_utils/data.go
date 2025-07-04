// Package common_utils /*
//
// Contains the License Data definition
// and the loading of data.
package common_utils

import (
	"fmt"
	tc "graphdbcli/internal/tool_configurations/statics"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// LicenseData represents the name of the license
// and the information related to it, referred as note.
type LicenseData struct {
	Name string
	Note string
}

// GetLicensesData loads the licenses information
//
// It fetches all the files that are ending on '.note'.
// Based on those we fetch the actual licenses. This is done
// By refersing the logging for adding licenses - whatever the provided
// license name is, the node '.note' suffix will be added for the note file.
// By removing the suffix the actual license name will used.
func GetLicensesData() []LicenseData {
	var homeDir, err = os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: Unable to get user home directory:", err)
		os.Exit(2)
	}

	var gdbDir = filepath.Join(homeDir, tc.HomeDirSpaceName)
	var licensesDir = filepath.Join(gdbDir, tc.LicensesDirName)

	files, err := os.ReadDir(licensesDir)
	if err != nil {
		log.Fatal(err)
	}

	var licensesData []LicenseData

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		filePath := filepath.Join(licensesDir, fileName)

		if filepath.Ext(fileName) == ".note" {
			content, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", filePath, err)
				continue
			}

			licensesData = append(licensesData, LicenseData{
				Name: strings.TrimSuffix(fileName, ".note"),
				Note: string(content),
			})
		}
	}

	return licensesData
}
