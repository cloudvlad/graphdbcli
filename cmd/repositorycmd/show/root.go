package show

import (
	"context"

	"github.com/spf13/cobra"
)

func Show(ctx context.Context) *cobra.Command {
	var command = &cobra.Command{
		Use:     "show",
		Short:   shortDescription,
		Long:    longDescription,
		Example: examples,
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return command
}
