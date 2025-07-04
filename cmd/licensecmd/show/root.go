// Package show is responsible for showing licenses
package show

import (
	"graphdbcli/internal/tool_configurations/logging"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:     "show",
		Short:   "show stored licenses",
		Example: "show",
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			logging.LOGGER.Info("Listing licenses...")
			TuiListStoredLicenses()
			return nil
		},
	}

	return command
}
