// Package show /*
//
// List the available GraphDB versions.
package list

import (
	"fmt"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tool_configurations/statics"
)

// listAvailableVersions show the available versions
// in no-TUI mode
func listAvailableVersions() {
	logging.LOGGER.Info("Listing available versions...")

	for versionNumber, _ := range statics.Versions {
		version := statics.Versions[versionNumber].Version
		releaseDate := statics.Versions[versionNumber].ReleaseDate
		fmt.Println(version, " - ", releaseDate)
	}
}
