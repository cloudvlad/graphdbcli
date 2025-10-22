package graph_store

import (
	"context"
	"graphdbcli/cmd/rdf4j/graph_store/add_stmnts"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func GraphStore(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	command := &cobra.Command{
		Use:     "graph-store",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			return nil
		},
	}

	command.AddCommand(add_stmnts.AddStatements(ctx, ctxCancel))

	return command
}
