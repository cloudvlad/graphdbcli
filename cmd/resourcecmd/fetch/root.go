// Package fetch provides the command to fetch a pre-uploaded resources.
package fetch

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:     "fetch [resource name]",
		Short:   "Fetch a resource",
		Example: "fetch starwars-data.ttl",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			resourceName := args[0]
			resource, err := findResourceByName(resourceName)
			FetchResource(*resource)

			return err
		},
	}

	return command
}
