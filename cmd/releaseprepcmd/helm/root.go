package helm

import (
	"context"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

var updateConfFilePath string
var updatedTargetDirPath string
var removeConfFileWhenDone bool

func Command(ctx context.Context) *cobra.Command {
	var command = &cobra.Command{
		Use:     "helm",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		RunE: func(cmd *cobra.Command, args []string) error {

			prepare(updateConfFilePath, updatedTargetDirPath, removeConfFileWhenDone)

			return nil
		},
	}

	command.Flags().StringVarP(&updateConfFilePath, "update-conf", "c", "./", "path to the Update YAML file")
	command.Flags().StringVarP(&updatedTargetDirPath, "updated-target", "u", "", "path to the updated target directory")
	command.Flags().BoolVar(&removeConfFileWhenDone, "remove-conf", false, "remove the configuration file when done")

	return command
}
