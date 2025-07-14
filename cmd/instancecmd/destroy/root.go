// Package destroy provides the command for destroying GraphDB instances.
// Stops the GraphDB instance and deletes the files associated with it.
package destroy

import (
	"context"
	"errors"

	"github.com/spf13/cobra"
)

func Command(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	var command = &cobra.Command{
		Use:     "destroy",
		Short:   "Destroy a cluster",
		Example: "destroy <name>\n",
		Aliases: []string{"d", "remove", "rm", "purge"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("destroy requires exactly one argument")
			}

			destroy(ctx, ctxCancel, args[0])

			return nil
		},
	}

	command.Flags().BoolVarP(&isRegexPattern, "regex", "r", false, "classifies the argument as a regex pattern")

	return command
}
