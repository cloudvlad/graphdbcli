// Package workbenchcmd provides the command for managing GraphDB versions.
package workbenchcmd

import (
	"context"
	"graphdbcli/cmd/workbenchcmd/add"
	"graphdbcli/cmd/workbenchcmd/config"
	"graphdbcli/cmd/workbenchcmd/purge"
	"graphdbcli/cmd/workbenchcmd/show"
	"graphdbcli/cmd/workbenchcmd/start"
	"graphdbcli/cmd/workbenchcmd/stop"

	"github.com/spf13/cobra"
)

func Workbench(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	command := &cobra.Command{
		Use:     "workbench",
		Short:   shortDescription,
		Long:    longDescription,
		Example: examples,
		Aliases: []string{"wb", "w"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			return nil
		},
		Hidden: true,
	}

	updateWorkbenchesMetadata()

	command.AddCommand(config.Command(ctx))
	command.AddCommand(add.Command())
	command.AddCommand(purge.Command(ctx))
	command.AddCommand(start.Command(ctx, ctxCancel))
	command.AddCommand(stop.Command(ctx, ctxCancel))
	command.AddCommand(show.Command())

	return command
}
