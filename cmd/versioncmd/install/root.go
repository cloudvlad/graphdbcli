// Package install provides the command for installing a specified version of the Platform Independent Distribution File.
package install

import (
	"context"
	"fmt"
	"graphdbcli/internal/tool_configurations/logging"
	"graphdbcli/internal/tool_configurations/statics"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var longDescription = `Installs a specified Platform Independent Distribution File by using a version.
The file is fetched from the Inter-Planetary Filesystem by using the corresponding CID,
Which is a hashed version of the PID file itself.`

func Command(ctx context.Context) *cobra.Command {
	var command = &cobra.Command{
		Use:     "install",
		Short:   "Installs a specified version of the Platform Independent Distribution File",
		Long:    longDescription,
		Example: "install 11.0.1",
		RunE: func(cmd *cobra.Command, args []string) error {
			var selectedVersion string

			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}

			if len(args) > 1 {
				return fmt.Errorf("too many arguments specified")
			}

			versionIndex := statics.GetVersionIndex(args[0])
			if versionIndex < 0 {
				fmt.Printf("version %s not found. The latest available version: %s\n", selectedVersion, statics.Versions[0].Version)
				logging.LOGGER.Fatal("unavailable version was specified",
					zap.String("version", selectedVersion))
			}
			version := statics.Versions[versionIndex]

			installSelectedVersion(version)
			integrityCheck(version)

			return nil
		},
	}

	command.Flags().BoolVarP(&isIntegrityCheckNeeded, "integrity-check", "c", false, "do an integrity check on the version that is being installed")

	return command
}
