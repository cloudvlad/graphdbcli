// Package list /*
//
// Responsible for the list command that
// shows the available GraphDB versions.
package list

import (
	"graphdbcli/internal/tool_configurations/logging"
	tc "graphdbcli/internal/tool_configurations/statics"
	"graphdbcli/internal/tui/licensetui"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {

	var command = &cobra.Command{
		Use:     "list",
		Short:   "List available versions",
		Example: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			logging.LOGGER.Info("Listing versions...")
			
			if tc.IsTuiDisabled {
				listAvailableVersions()
			} else {
				TuiListAvailableVersions()
			}

			println(licensetui.SelectedLicense)
			return nil
		},
	}

	return command
}
