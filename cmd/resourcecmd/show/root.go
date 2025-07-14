// Package show provides the command for listing resources.
package show

import (
	"graphdbcli/internal/tool_configurations/logging"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:     "show",
		Short:   "show the available resources",
		Example: "show",
		RunE: func(cmd *cobra.Command, args []string) error {
			logging.LOGGER.Info("Listing resources...")
			return nil
		},
	}

	return command
}
