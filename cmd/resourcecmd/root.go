// Package resourcecmd provides the command for managing resources.
package resourcecmd

import (
	"context"
	"graphdbcli/cmd/resourcecmd/fetch"
	"graphdbcli/cmd/resourcecmd/show"

	"github.com/spf13/cobra"
)

func Resource(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:     "resource",
		Short:   "Manages resources",
		Example: "dataset show",
		Aliases: []string{"r", "resources"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			return nil
		},
		Hidden: true,
	}

	command.AddCommand(show.Command())
	command.AddCommand(fetch.Command())

	return command
}
