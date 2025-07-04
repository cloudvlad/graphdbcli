// Package add provides the command for adding licenses.
package add

import (
	"fmt"
	"graphdbcli/internal/tool_configurations/logging"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:   "add <path>",
		Short: "Add a license",
		Example: "add ./graphdb.license\n" +
			"add ./another.license --note 'Expiration date 2024-01-01'\n" +
			"add /home/mark/license.file --note 'GraphDB 11 License'\n" +
			"a ./license.file --note 'GraphDB 11 License'\n" +
			"a ./license.file",
		Aliases: []string{"a"},
		RunE: func(cmd *cobra.Command, args []string) error {
			var LicenseFilepath string

			if len(args) == 0 {
				fmt.Println(logging.ErrorMessages[001].External)
				logging.LOGGER.Fatal(logging.ErrorMessages[001].Internal)
			}
			LicenseFilepath = args[0]

			StoreLicenseFile(LicenseFilepath)

			return nil
		},
	}

	command.Flags().StringVarP(&LicenseFileNoteContent, "note", "n", "", "note to add the license")
	command.Flags().StringVarP(&NewLicenseFileName, "name", "", "", "set the new name of the license file")

	return command
}
