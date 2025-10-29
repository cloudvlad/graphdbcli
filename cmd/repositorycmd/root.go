// Package repositorycmd manages GraphDB repositories
package repositorycmd

import (
	"context"
	"graphdbcli/cmd/repositorycmd/create"
	pf "graphdbcli/internal/flags/repositorycmd"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Repository(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	command := &cobra.Command{
		Use:     "repository",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		Aliases: []string{"r", "repositories", "repo", "repos"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	command.AddCommand(create.Create(ctx, ctxCancel))

	command.PersistentFlags().StringVarP(&pf.GraphdbAddress, "graphdb-address", "a", "http://localhost:7200", "GraphDB instance address")
	command.PersistentFlags().StringVarP(&pf.Repository, "repository", "r", "", "GraphDB repository name")

	command.MarkFlagRequired("repository")

	return command
}
