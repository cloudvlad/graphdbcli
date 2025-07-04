// Package list /*
//
// Contains the utilities logic for the
// listing of licenses.
package list

import (
	"fmt"
	cu "graphdbcli/cmd/licensecmd/common_utils"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tui/licensetui"
	"github.com/charmbracelet/bubbles/list"
	"os"
)

// TuiListStoredLicenses list the stored licenses of
// the tool in TUI mode.
func TuiListStoredLicenses() {
	var items []list.Item

	for _, license := range cu.GetLicensesData() {
		logging.LOGGER.Info("License found: " + license.Name)
		item := licensetui.Item{
			Name: license.Name,
			Note: license.Note,
		}

		items = append(items, item)
	}

	logging.LOGGER.Info("Selected license: " + licensetui.SelectedLicense)

	p := licensetui.Initialize(items)
	if _, err := p.Run(); err != nil {
		logging.LOGGER.Error("Error running program: " + err.Error())
		os.Exit(1)
	}
}

// listStoredLicenses list the stored licenses of the tool in
// no-TUI mode.
func listStoredLicenses() {
	for _, license := range cu.GetLicensesData() {
		fmt.Printf("License: %s, Note: %s\n", license.Name, license.Note)
	}
}
