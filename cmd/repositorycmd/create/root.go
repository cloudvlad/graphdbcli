package create

import (
	"context"

	"github.com/spf13/cobra"
)

func Create(ctx context.Context) *cobra.Command {
	var command = &cobra.Command{
		Use:     "create",
		Short:   "show stored licenses",
		Example: "show",
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return command
}
