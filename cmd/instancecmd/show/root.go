// Package show provides the command for displaying GraphDB instances.
package show

import (
	"graphdbcli/internal/tool_configurations/logging"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:     "show",
		Short:   "shows the instances",
		Example: "instance show\n",
		RunE: func(cmd *cobra.Command, args []string) error {
			logging.LOGGER.Info("Showing clusters...")
			showInstances()
			return nil
		},
	}

	return command
}
