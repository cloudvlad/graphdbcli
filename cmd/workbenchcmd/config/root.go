// Package configcmd provides the command for managing GraphDB versions.
package config

import (
	"context"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Command(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:     "config <workbench-name>",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		Aliases: []string{"c"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			ConfigureCustomWorkbench(args[0])

			return nil
		},
	}

	command.Flags().UintVarP(&graphdbPort, "graphdb-port", "g", 7200, "GraphDB instance port")
	command.Flags().StringVarP(&graphdbHost, "graphdb-host", "l", "localhost", "GraphDB instance location")
	command.Flags().UintVarP(&workbenchPort, "workbench-port", "w", 9000, "Workbench instance port")

	return command
}
