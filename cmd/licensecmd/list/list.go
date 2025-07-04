// Package list /*
//
// Responsible for listing licenses.
//
// Supports TUI and no-TUI mode.
package list

import (
	"graphdbcli/internal/tool_configurations/logging"
	tc "graphdbcli/internal/tool_configurations/statics"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:     "list",
		Short:   "list stored licenses",
		Example: "list",
		Aliases: []string{"l"},
		RunE: func(cmd *cobra.Command, args []string) error {
			logging.LOGGER.Info("Listing licenses...")
			if tc.IsTuiDisabled {
				listStoredLicenses()
			} else {
				TuiListStoredLicenses()
			}
			return nil
		},
	}

	return command
}
