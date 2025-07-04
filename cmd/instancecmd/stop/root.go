// Package stop provides the command for stopping GraphDB instances.
package stop

import (
	"errors"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:     "stop <instance-name>",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("stop requires instance name")
			}

			instanceName = args[0]
			stopInstance()

			return nil
		},
	}

	command.Flags().BoolVarP(&forceStop, "force", "f", false, "force stop the instance")
	command.Flags().BoolVarP(&isRegexPattern, "regex", "r", false, "classifies the argument as a regex pattern")

	return command
}
