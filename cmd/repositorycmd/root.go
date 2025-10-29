// Package repositorycmd manages GraphDB repositories
package repositorycmd

import (
	"context"
	"graphdbcli/cmd/repositorycmd/create"
	"graphdbcli/cmd/repositorycmd/delete"
	"graphdbcli/cmd/repositorycmd/info"
	"graphdbcli/cmd/repositorycmd/list"
	"graphdbcli/cmd/repositorycmd/restart"
	pf "graphdbcli/internal/flags/repositorycmd"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Repository(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:     "repository",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		Aliases: []string{"r", "repositories"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		Hidden: true,
	}

	command.AddCommand(create.Create(ctx))
	command.AddCommand(delete.Delete(ctx))
	command.AddCommand(info.Info(ctx))
	command.AddCommand(list.List(ctx))
	command.AddCommand(restart.Restart(ctx))

	command.PersistentFlags().StringVarP(&pf.GraphdbAddress, "graphdb-address", "a", "http://localhost:7200", "GraphDB instance address")
	command.PersistentFlags().StringVarP(&pf.Repository, "repository", "r", "", "GraphDB repository name")

	command.MarkFlagRequired("repository")

	return command
}
