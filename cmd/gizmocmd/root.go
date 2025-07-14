package gizmocmd

import (
	"graphdbcli/cmd/gizmocmd/perf_test"
	"graphdbcli/cmd/gizmocmd/rdf_bombarder"
	"graphdbcli/cmd/gizmocmd/statements_stats"
	"graphdbcli/cmd/gizmocmd/ttl_generator"

	"github.com/spf13/cobra"
)

func Gizmo() *cobra.Command {
	command := &cobra.Command{
		Use:     "gizmo",
		Short:   "Collection of GraphDB-related gizmo commands",
		Aliases: []string{"gizmos", "g"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		Hidden: true,
	}

	command.AddCommand(statements_stats.Command())
	command.AddCommand(perf_test.Command())
	command.AddCommand(rdf_bombarder.Command())
	command.AddCommand(ttl_generator.Command())

	return command
}
