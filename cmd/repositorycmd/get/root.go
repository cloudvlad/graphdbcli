package get

import (
	"github.com/spf13/cobra"
)

func Get() *cobra.Command {
	var command = &cobra.Command{
		Use:     "show",
		Short:   "show stored licenses",
		Example: "show",
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return command
}
