// Package remove
//
// Responsible for removing licenses and the
// data stored alongside them.
//
// Supports TUI and no-TUI mode.
package remove

import (
	"graphdbcli/internal/tool_configurations/logging"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var log = logging.LOGGER

func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:     "remove",
		Short:   "remove stored licenses",
		Example: "remove license.graphdb",
		Aliases: []string{"rm"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				log.Info("Removing license", zap.String("license name", args[0]))
				RemoveLicenseFile(args[0])
			}

			return nil
		},
	}

	return command
}
