// Package install provides the command for installing a specified version of the Platform Independent Distribution File.
package install

import (
	"context"
	"fmt"
	"graphdbcli/internal/data_objects/install"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tool_configurations/statics"
	"graphdbcli/internal/tui/common_components"

	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var isIntegrityCheckNeeded bool

func Command(ctx context.Context) *cobra.Command {
	var command = &cobra.Command{
		Use:     "install",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		Aliases: []string{"i"},
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) == 0 {
				TuiListAvailableVersions()
			} else {
				install.SelectedVersion = args[0]
			}

			if len(args) > 1 {
				return fmt.Errorf("too many arguments specified")
			}

			versionIndex := statics.GetVersionIndex(install.SelectedVersion)
			var version statics.Version
			if versionIndex >= 0 {
				version = statics.Versions[versionIndex]
			}

			installSelectedVersion(install.SelectedVersion)

			if isIntegrityCheckNeeded {
				if versionIndex < 0 {
					fmt.Printf("%s The checksum for version %s not indexed. Integrity check is skippied...", common_components.PadStatusIndicator(emoji.NextTrackButton.String(), 0), install.SelectedVersion)
					logging.LOGGER.Info("unavailable version was specified",
						zap.String("version", install.SelectedVersion))
				} else {
					integrityCheck(version)
				}
			}

			return nil
		},
	}

	command.Flags().BoolVarP(&isIntegrityCheckNeeded, "integrity-check", "c", false, "Do an integrity check on the version that is being installed")

	return command
}
