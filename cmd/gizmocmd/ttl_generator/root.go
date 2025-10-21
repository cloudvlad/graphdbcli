package ttl_generator

import (
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var command = &cobra.Command{
		Use:     "ttl-generator <output-dir>",
		Short:   "Generates RDF files in the Turtle format",
		Example: common_components.PadExamples(examples),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			generateTTLFiles(args[0])

			return nil
		},
	}

	command.Flags().UintVarP(&numberOfEntitiesPerFile, "numberOfEntitiesPerFile", "e", 1, "number of entities per file")
	command.Flags().UintVarP(&numberOfFiles, "numberOfFiles", "f", 1, "number of files to generate")
	command.Flags().UintVarP(&numberOfTripletsPerEntity, "numberOfTripletsPerEntity", "t", 1, "number of triplets per entity")
	return command
}
