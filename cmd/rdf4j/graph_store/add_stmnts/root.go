package add_stmnts

import (
	"context"
	"graphdbcli/internal/tui/common_components"
	dlo "graphdbcli/internal/tui/rdf4j"

	"github.com/spf13/cobra"
)

var (
	graph string
)

func AddStatements(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	command := &cobra.Command{
		Use:     "add-statements <data-file-path>",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		Aliases: []string{"as", "add-stmts", "a-s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			addStatements(ctx, ctxCancel, args[0], graph)

			return nil
		},
	}

	command.Flags().StringVarP(&graph, "graph", "g", "", "named graph name")
	command.PersistentFlags().StringVarP(&dlo.RdfFormat, "rdfFormat", "f", "text/turtle", "RDF Serialization format")

	return command
}
