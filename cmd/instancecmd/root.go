// Package instancecmd provides the command for managing GraphDB instances.
package instancecmd

import (
	"context"
	"graphdbcli/cmd/instancecmd/create"
	"graphdbcli/cmd/instancecmd/destroy"
	"graphdbcli/cmd/instancecmd/show"
	"graphdbcli/cmd/instancecmd/stop"

	"github.com/spf13/cobra"
)

func Cluster(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	command := &cobra.Command{
		Use:     "instance",
		Short:   shortDescription,
		Long:    longDescription,
		Aliases: []string{"instances", "i"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	// Executed right before running any other subcommand.
	// So we can ensure up-to-date state of the stored configurations.
	updateInstancesMetadata()

	command.AddCommand(create.Create(ctx, ctxCancel))
	command.AddCommand(destroy.Command(ctx, ctxCancel))
	command.AddCommand(show.Command())
	command.AddCommand(stop.Command())

	return command
}
