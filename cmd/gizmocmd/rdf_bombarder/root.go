// Package rdf_bombarder provides the command for sending multiple statements to GraphDB.
package rdf_bombarder

import (
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:     "rdf_bombarder",
		Short:   "Bomabrds graphdb with statemtns",
		Example: common_components.PadExamples(examples),
		RunE: func(cmd *cobra.Command, args []string) error {
			bombard()

			return nil
		},
	}

	command.Flags().StringVarP(&instanceAddress, "instanceAddress", "a", "", "instance address")
	command.Flags().StringVarP(&repoName, "repositoryName", "r", "", "repository name")
	command.Flags().IntVarP(&numberOfNamedGraphs, "numberOfNamedGraphs", "g", 0, "number of named graphs")
	command.Flags().IntVarP(&numberOfStatementsPerNamedGraph, "numberOfStatementsPerNamedGraph", "n", 10, "number Of Statements Per Named Graph")
	command.Flags().IntVarP(&numberOfThreads, "numberOfThreads", "t", 1, "number of threads")
	command.Flags().IntVarP(&numberOfStatementsPerRequest, "numberOfStatementsPerRequest", "q", 10, "number of statements per request")
	command.Flags().Float64VarP(&interconnection, "interconnection", "i", 0.0, "interconnection")

	cobra.MarkFlagRequired(command.Flags(), "instanceAddress")
	cobra.MarkFlagRequired(command.Flags(), "repositoryName")

	return command
}
