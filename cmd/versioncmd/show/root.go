// Package show provides the command for showing available versions.
package show

import (
	"context"
	"graphdbcli/internal/tool_configurations/logging"

	"github.com/spf13/cobra"
)

func Command(ctx context.Context) *cobra.Command {

	var command = &cobra.Command{
		Use:     "show",
		Short:   "List available versions",
		Example: "show",
		RunE: func(cmd *cobra.Command, args []string) error {
			logging.LOGGER.Info("Listing available versions...")
			TuiListAvailableVersions()
			return nil
		},
	}

	return command
}
