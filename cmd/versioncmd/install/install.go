// Package install /*
//
// Responsible for installing, context of downloading
// GraphDB platform independent distribution files.
package install

import (
	"fmt"
	"graphdbcli/cmd/versioncmd/list"
	"graphdbcli/internal/tool_configurations/statics"
	"graphdbcli/internal/tui/versiontui"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:     "install",
		Short:   "Install a version",
		Example: "install 11.0.1",
		RunE: func(cmd *cobra.Command, args []string) error {
			var selectedVersion string

			if len(args) == 0 && !statics.IsTuiDisabled {
				list.TuiListAvailableVersions()
				selectedVersion = versiontui.SelectedVersion
			}

			if len(args) == 1 {
				selectedVersion = args[0]
			}

			if len(args) > 1 {
				return fmt.Errorf("too many arguments specified")
			}

			installSelectedVersion(selectedVersion)

			return nil
		},
	}

	return command
}
