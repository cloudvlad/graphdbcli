package sparql

import (
	"context"
	"graphdbcli/cmd/rdf4j/sparql/upload"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Sparql(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	command := &cobra.Command{
		Use:     "sparql",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			return nil
		},
	}

	command.AddCommand(upload.Upload(ctx, ctxCancel))

	return command
}
