package rdf4j

import (
	"context"
	"graphdbcli/cmd/rdf4j/graph_store"
	"graphdbcli/cmd/rdf4j/sparql"
	dlo "graphdbcli/internal/tui/rdf4j"

	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Rdf4J(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	command := &cobra.Command{
		Use:     "rdf4j <subcommand>",
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

	command.AddCommand(sparql.Sparql(ctx, ctxCancel))
	command.AddCommand(graph_store.GraphStore(ctx, ctxCancel))

	command.PersistentFlags().StringVarP(&dlo.GraphdbAddress, "graphdb-address", "a", "http://localhost:7200", "GraphDB instance address")
	command.PersistentFlags().StringVarP(&dlo.Repository, "repository", "r", "", "GraphDB repository name")
	command.PersistentFlags().StringVarP(&dlo.RdfFormat, "rdfFormat", "f", "text/turtle", "RDF Format header")

	command.MarkFlagRequired("repository")

	return command
}
