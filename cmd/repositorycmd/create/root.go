package create

import (
	"context"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

var (
	baseURL                                string
	checkForInconsistencies                bool
	defaultNS                              string
	disableSameAs                          bool
	enableContextIndex                     bool
	enableFtsIndex                         bool
	enableLiteralIndex                     bool
	enablePredicateList                    bool
	entityIdSize                           int
	entityIndexSize                        int
	ftsIndexes                             []string
	ftsIrisIndex                           string
	ftsStringLiteralsIndex                 string
	imports                                []string
	inMemoryLiteralProperties              bool
	queryLimitResults                      int
	queryTimeout                           int
	readOnly                               bool
	repositoryType                         string
	ruleset                                string
	storageFolder                          string
	throwQueryEvaluationExceptionOnTimeout bool
)

func Create(ctx context.Context) *cobra.Command {
	var command = &cobra.Command{
		Use:     "create <repository-name>",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		Aliases: []string{"c"},
		RunE: func(cmd *cobra.Command, args []string) error {
			// ...existing code...
			return nil
		},
	}

	command.Flags().StringVar(&baseURL, "base-url", "http://example.org/owlim#", "Specify the default namespace for the main persistence file")
	command.Flags().BoolVar(&checkForInconsistencies, "check-for-inconsistencies", false, "Enable the mechanism for consistency checking")
	command.Flags().StringVar(&defaultNS, "default-ns", "", "Specify the default namespaces corresponding to each imported schema file, separated by semicolons")
	command.Flags().BoolVar(&disableSameAs, "disable-sameas", true, "Disable the owl:sameAs optimization")
	command.Flags().BoolVar(&enableContextIndex, "enable-context-index", false, "Enable the context index")
	command.Flags().BoolVar(&enableFtsIndex, "enable-fts-index", false, "Disable the full-text search index")
	command.Flags().BoolVar(&enableLiteralIndex, "enable-literal-index", true, "Disable the literal index")
	command.Flags().BoolVar(&enablePredicateList, "enable-predicate-list", true, "Disable the predicate list")
	command.Flags().IntVar(&entityIdSize, "entity-id-size", 32, "Define the bit size of internal IDs used to index entities")
	command.Flags().IntVar(&entityIndexSize, "entity-index-size", 10000000, "Define the initial size of the entity hash table index entries.")
	command.Flags().StringSliceVar(&ftsIndexes, "fts-indexes", []string{"default", "iri"}, "Comma-delimited list of languages that should have a specific index with an appropriate analyzer for full-text search.")
	command.Flags().StringVar(&ftsIrisIndex, "fts-iris-index", "none", "Specify the index in which the IRIs are indexed for full-text search")
	command.Flags().StringVar(&ftsStringLiteralsIndex, "fts-string-literals-index", "default", "Specify the index in which the string literals are indexed")
	command.Flags().StringSliceVar(&imports, "imports", []string{""}, "Specify the list of schema files that will be imported at startup")
	command.Flags().BoolVar(&inMemoryLiteralProperties, "in-memory-literal-properties", true, "Disables caching of the literal languages and data types")
	command.Flags().IntVar(&queryLimitResults, "query-limit-results", 0, "Set the maximum number of results returned from a query")
	command.Flags().IntVar(&queryTimeout, "query-timeout", 0, "Set the number of seconds after which the evaluation of a query will be terminated")
	command.Flags().BoolVar(&readOnly, "read-only", false, "Sets the repository in read-only mode")
	command.Flags().StringVar(&repositoryType, "repository-type", "file-repository", "Define the repository type")
	command.Flags().StringVar(&ruleset, "ruleset", "rdfsplus-optimized", "Set of axiomatic triples, consistency checks and entailment rules, which determine the applied semantics")
	command.Flags().StringVar(&storageFolder, "storage-folder", "storage", "Specifies the folder where the index files will be stored")
	command.Flags().BoolVar(&throwQueryEvaluationExceptionOnTimeout, "throw-query-evaluation-exception-on-timeout", false, "")

	command.MarkFlagsRequiredTogether("enable-fts-index", "fts-indexes", "fts-string-literals-index", "fts-iris-index")
	command.MarkFlagsRequiredTogether("default-ns", "imports")

	return command
}
