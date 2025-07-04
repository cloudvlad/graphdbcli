// Package show /*
//
// Contains the utilities logic for the
// listing of licenses.
package show

import (
	"fmt"
	"graphdbcli/internal/data_objects/license"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tui/licensetui"
	"os"

	"github.com/charmbracelet/bubbles/list"
)

// TuiListStoredLicenses show the stored licenses of
// the tool in TUI mode.
func TuiListStoredLicenses() {
	var items []list.Item

	for _, licenseRecord := range license.GetLicensesData() {
		logging.LOGGER.Debug("License found: " + licenseRecord.GetName())
		item := licensetui.Item{
			Name: licenseRecord.GetName(),
			Note: licenseRecord.GetNote(),
		}

		items = append(items, item)
	}

	logging.LOGGER.Info("Selected licenseRecord: " + licensetui.SelectedLicense)

	p := licensetui.Initialize(items)
	if _, err := p.Run(); err != nil {
		logging.LOGGER.Error("Error running program: " + err.Error())
		os.Exit(1)
	}
}

// listStoredLicenses show the stored licenses of the tool in
// no-TUI mode.
func listStoredLicenses() {
	for _, licenseRecord := range license.GetLicensesData() {
		fmt.Printf("License: %s, Note: %s\n", licenseRecord.GetName(), licenseRecord.GetNote())
	}
}
