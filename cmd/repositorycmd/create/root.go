package create

import (
	"context"
	repo "graphdbcli/internal/data_objects/repository"
	pf "graphdbcli/internal/flags/repositorycmd"
	"graphdbcli/internal/tui/common_components"

	"github.com/spf13/cobra"
)

var eidSize40 bool

var rc = repo.Config{
	ID:     "",
	Title:  "",
	Params: repo.Params{},
}

func Create(ctx context.Context, ctxCancel context.CancelFunc) *cobra.Command {
	var command = &cobra.Command{
		Use:     "create --repository <repository-name>",
		Short:   shortDescription,
		Long:    longDescription,
		Example: common_components.PadExamples(examples),
		Aliases: []string{"c"},
		Args: func(cmd *cobra.Command, args []string) error {
			rc.SetDefaults()

			rc.ID = pf.Repository
			rc.Params.EntityIdSize = 32

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			validationStatus := validation(rc)
			if validationStatus != nil {
				return validationStatus
			}

			if eidSize40 {
				rc.Params.EntityIdSize = 40
			}

			createRepository(ctx, ctxCancel, rc, pf.GraphdbAddress)

			return nil
		},
	}

	command.Flags().BoolVar(&eidSize40, "eidSize40", false, "Use 40-bit entity ID size (default is 32)")

	command.Flags().StringVar(&rc.Title, "desc", "", "Repository description")
	command.Flags().BoolVar(&rc.Params.ReadOnly, "readOnly", false, "Set repository as read-only")
	command.Flags().StringVar(&rc.Params.Ruleset, "ruleset", "rdfsplus-optimized", "Inference ruleset")
	command.Flags().BoolVar(&rc.Params.DisableSameAs, "disableOwlSameAs", true, "Disable owl:sameAs reasoning")
	command.Flags().BoolVar(&rc.Params.CheckForInconsistencies, "enableConsistencyChecks", false, "Enable consistency checks")
	command.Flags().BoolVar(&rc.Params.IsShacl, "enableShaclValidation", false, "Enable SHACL validation")
	command.Flags().BoolVar(&rc.Params.CacheSelectNodes, "enableCacheSelectNodes", true, "Enable cache select nodes")
	command.Flags().BoolVar(&rc.Params.LogValidationPlans, "enableLogValidationPlans", false, "Log validation plans")
	command.Flags().BoolVar(&rc.Params.ParallelValidation, "enableParallelValidation", true, "Enable parallel validation")
	command.Flags().BoolVar(&rc.Params.PerformanceLogging, "enablePerformanceLogging", false, "Enable performance logging")
	command.Flags().BoolVar(&rc.Params.DashDataShapes, "enableDashDataShapes", true, "Enable DASH data shapes")
	command.Flags().BoolVar(&rc.Params.LogValidationViolations, "enableLogValidationViolations", false, "Log validation violations")
	command.Flags().BoolVar(&rc.Params.GlobalLogValidationExecution, "enableGlobalLogValidationExecution", false, "Global log validation execution")
	command.Flags().BoolVar(&rc.Params.EclipseRdf4jShaclExtensions, "enableEclipseRdf4jShaclExtensions", true, "Enable Eclipse RDF4J SHACL extensions")
	command.Flags().Uint32Var(&rc.Params.ValidationResultsLimitTotal, "validationResultsLimitTotal", 1000000, "Validation results limit total")
	command.Flags().Uint32Var(&rc.Params.ValidationResultsLimitPerConstraint, "validationResultsLimitPerConstraint", 1000, "Validation results limit per constraint")
	command.Flags().StringVar(&rc.Params.ShapesGraph, "shapesGraph", "http://rdf4j.org/schema/rdf4j#SHACLShapeGraph", "Shapes graph URI")
	command.Flags().BoolVar(&rc.Params.EnableContextIndex, "enableContextIndex", false, "Enable context index")
	command.Flags().BoolVar(&rc.Params.EnablePredicateList, "enablePredicateList", true, "Enable predicate list")
	command.Flags().BoolVar(&rc.Params.EnableFtsIndex, "enableFtsIndex", false, "Enable FTS index")
	command.Flags().Uint32Var(&rc.Params.EntityIndexSize, "entityIndexSize", 10000000, "Entity index size")
	command.Flags().StringSliceVar(&rc.Params.FtsIndexes, "ftsIndexes", []string{"default", "iri"}, "FTS indexes")
	command.Flags().StringVar(&rc.Params.FtsStringLiteralsIndex, "ftsStringLiteralsIndex", "default", "FTS string literals index")
	command.Flags().StringVar(&rc.Params.FtsIrisIndex, "ftsIrisIndex", "none", "FTS IRIs index")
	command.Flags().BoolVar(&rc.Params.ThrowQueryEvaluationExceptionOnTimeout, "throwQueryEvaluationExceptionOnTimeout", false, "Throw query evaluation exception on timeout")
	command.Flags().Uint32Var(&rc.Params.QueryTimeout, "queryTimeout", 0, "Query timeout in seconds")
	command.Flags().Uint32Var(&rc.Params.QueryLimitResults, "queryLimitResults", 0, "Query limit results")

	return command
}
