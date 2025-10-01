// Package create provides the command for creating GraphDB instances.
package create

import (
	"context"
	c "graphdbcli/internal/data_objects/graphdb_cluster"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Create(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	var command = &cobra.Command{
		Use:     "create",
		Short:   "Create a GraphDB instance",
		Example: common_components.PadExamples(examples),
		Aliases: []string{"c"},
		RunE: func(cmd *cobra.Command, args []string) error {
			createGraphDBInstance(ctx, ctxCancel)

			return nil
		},
	}

	command.Flags().StringVarP(&c.Instance.Name, "name", "n", "", "name of the GraphDB cluster")
	command.Flags().StringVarP(&c.Instance.Version, "version", "v", "", "version of the GraphDB cluster")
	command.Flags().StringVarP(&c.Instance.StoredLicenseFilename, "license", "l", "", "stored license file name")
	command.Flags().BoolVarP(&c.Instance.IsActive, "activate", "a", true, "activates the instance right after bootstrapping it")
	command.Flags().StringVarP(&c.PropertyOverridesFilepath, "properties", "", "", "property overrides file")
	command.Flags().StringVarP(&c.Instance.Port, "port", "p", "", "defines the port for exposing the instance")

	return command
}
