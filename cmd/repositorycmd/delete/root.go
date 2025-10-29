package delete

import (
	"context"

	"github.com/spf13/cobra"
)

func Delete(ctx context.Context) *cobra.Command {
	var command = &cobra.Command{
		Use:     "remove",
		Short:   "show stored licenses",
		Example: "show",
		Aliases: []string{"r", "rm"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return command
}
