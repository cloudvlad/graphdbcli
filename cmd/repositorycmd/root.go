// Package repositorycmd manages GraphDB repositories
package repositorycmd

import (
	"context"
	"graphdbcli/cmd/repositorycmd/create"
	"graphdbcli/cmd/repositorycmd/remove"
	"graphdbcli/cmd/repositorycmd/show"

	"github.com/spf13/cobra"
)

func Repository(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:     "repository",
		Short:   shortDescription,
		Long:    longDescription,
		Example: examples,
		Aliases: []string{"r", "repositories"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		Hidden: true,
	}

	command.AddCommand(create.Create(ctx))
	command.AddCommand(show.Show(ctx))
	command.AddCommand(remove.Remove(ctx))

	return command
}
