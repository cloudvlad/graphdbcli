// Package licensecmd /*
//
// Defines the license command and its subcommands.
// The subcommands definition and logic is moved to
// separate packages.
package licensecmd

import (
	"graphdbcli/cmd/licensecmd/add"
	"graphdbcli/cmd/licensecmd/list"
	"graphdbcli/cmd/licensecmd/remove"
	"github.com/spf13/cobra"
)

func License() *cobra.Command {
	command := &cobra.Command{
		Use:     "license",
		Short:   "Manages GraphDB licenses",
		Example: "licenses",
		Aliases: []string{"l", "licenses"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			return nil
		},
	}

	command.AddCommand(add.Command())
	command.AddCommand(remove.Command())
	command.AddCommand(list.Command())

	return command
}
