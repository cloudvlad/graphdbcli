package upload

import (
	"context"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

var (
	baseUri string
)

func Upload(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	command := &cobra.Command{
		Use:     "upload <data-file-path>",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		Aliases: []string{"u"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			uploadData(ctx, ctxCancel, args[0], baseUri)

			return nil
		},
	}

	command.Flags().StringVarP(&baseUri, "baseUri", "b", "", "specifies the base URI to resolve any relative URIs found in uploaded data against")

	return command
}
