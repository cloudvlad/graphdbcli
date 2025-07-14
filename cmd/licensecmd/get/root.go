// Package get provides the command for getting the license content in base64 format
package get

import (
	"encoding/base64"
	"fmt"
	types "graphdbcli/internal/data_objects/license"
	"graphdbcli/internal/tool_configurations/logging"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:     "get <license name>",
		Short:   "Get a license in Base64 encoded format",
		Example: "get graphdb.license\n",
		Aliases: []string{"g"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				fmt.Println("Please provide a license name")
				logging.LOGGER.Fatal("License name was not provided")
			}
			encodedContent := base64.StdEncoding.EncodeToString(types.GetLicenseContent(args[0]))
			fmt.Println(encodedContent)

			return nil
		},
	}

	return command
}
