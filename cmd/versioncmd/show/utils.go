// Package show /*
//
// List the available GraphDB versions.
package show

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tool_configurations/statics"
	"graphdbcli/internal/tui/versiontui"
	"os"
)

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

// listAvailableVersions show the available versions
// in no-TUI mode
func listAvailableVersions() {
	for versionNumber, _ := range statics.Versions {
		version := statics.Versions[versionNumber].Version
		releaseDate := statics.Versions[versionNumber].ReleaseDate
		fmt.Println(version, " - ", releaseDate)
	}
}
