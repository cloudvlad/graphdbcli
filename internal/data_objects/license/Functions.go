package license

import (
	"fmt"
	"graphdbcli/internal/tool_configurations/initialization"
	"graphdbcli/internal/tool_configurations/logging"
	"log"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
)

// GetLicensesData loads the licenses information
//
// It fetches all the files that are ending on '.note'.
// Based on those we fetch the actual licenses.
// By removing the suffix the actual license name will be used.
func GetLicensesData() []LicenseData {
	licensesDir := initialization.GetLicensesDirectory()

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

// GetLicenseContent returns the actual content of the
// stored license by using the provided license name.
// The license name corresponds to the name of the file.
func GetLicenseContent(licenseName string) []byte {
	licensesFilePath := filepath.Join(initialization.GetLicensesDirectory(), licenseName)
	_, err := os.Stat(licensesFilePath)
	if err != nil {
		fmt.Println("There was problem checking the license file")
		logging.LOGGER.Fatal("There was problem checking the license file",
			zap.String("license_name", licenseName))
	}

	content, err := os.ReadFile(licensesFilePath)
	if err != nil {
		fmt.Println("There was problem reading the license file")
		logging.LOGGER.Fatal("There was problem reading the license file",
			zap.String("license_name", licenseName))
	}

	return content
}
