// Package remove /*
//
// Responsible for removing licenses and the
// data stored alongside them.
//
// Supports TUI and no-TUI mode.
package remove

import (
	"graphdbcli/cmd/versioncmd/list"
	"graphdbcli/internal/tool_configurations/statics"
	"graphdbcli/internal/tui/versiontui"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:     "remove",
		Short:   "remove stored licenses",
		Example: "remove license.graphdb",
		Aliases: []string{"rm"},
		RunE: func(cmd *cobra.Command, args []string) error {
			var selectedLicense string

			if len(args) == 0 && !statics.IsTuiDisabled {
				list.TuiListAvailableVersions()
				selectedLicense = versiontui.SelectedVersion
			}

			if len(args) == 1 {
				selectedLicense = args[0]
			}

			RemoveLicenseFile(selectedLicense)
			return nil
		},
	}

	return command
}
