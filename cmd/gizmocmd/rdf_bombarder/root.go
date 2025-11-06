// Package rdf_bombarder provides the command for sending multiple statements to GraphDB.
package rdf_bombarder

import (
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var instanceAddress string
	var repoName string

	var numberOfNamedGraphs int
	var numberOfStatementsPerNamedGraph int
	var numberOfThreads int
	var numberOfStatementsPerRequest int
	var interconnection float64

	var command = &cobra.Command{
		Use:     "rdf_bombarder",
		Short:   "Bombards GraphDB repository with statements",
		Example: common_components.PadExamples(examples),
		RunE: func(cmd *cobra.Command, args []string) error {
			bombard(instanceAddress, repoName, numberOfNamedGraphs, numberOfStatementsPerNamedGraph, numberOfThreads, numberOfStatementsPerRequest, interconnection)

			return nil
		},
	}

	command.Flags().StringVarP(&instanceAddress, "instanceAddress", "a", "http://localhost:7200", "GraphDB instance address")
	command.Flags().StringVarP(&repoName, "repositoryName", "r", "", "repository name")
	command.Flags().IntVarP(&numberOfNamedGraphs, "numberOfNamedGraphs", "g", 0, "number of named graphs")
	command.Flags().IntVarP(&numberOfStatementsPerNamedGraph, "numberOfStatementsPerNamedGraph", "n", 10, "number Of Statements Per Named Graph")
	command.Flags().IntVarP(&numberOfThreads, "numberOfThreads", "t", 1, "number of threads")
	command.Flags().IntVarP(&numberOfStatementsPerRequest, "numberOfStatementsPerRequest", "q", 10, "number of statements per request")
	command.Flags().Float64VarP(&interconnection, "interconnection", "i", 0.0, "interconnection")

	cobra.MarkFlagRequired(command.Flags(), "repositoryName")

	return command
}
