package repository

import (
	"fmt"
	"strings"
)

// ToAcceptableJson converts the Config struct to a map[string]interface{} in the API-accepted format.
// Please delete this method and use the ToJSON when fixed.
func (c *Config) ToAcceptableJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":       c.ID,
		"title":    c.Title,
		"type":     c.Type,
		"location": c.Location,
	}
	params := map[string]interface{}{}

	toStr := func(v interface{}) string {
		switch val := v.(type) {
		case string:
			return val
		case uint8:
			return fmt.Sprintf("%d", val)
		case uint32:
			return fmt.Sprintf("%d", val)
		case uint64:
			return fmt.Sprintf("%d", val)
		case bool:
			if val {
				return "true"
			}
			return "false"
		default:
			return fmt.Sprintf("%v", val)
		}
	}

	toCSV := func(v []string) string {
		return strings.Join(v, ", ")
	}

	p := c.Params
	params["queryTimeout"] = map[string]interface{}{ "name": "queryTimeout", "label": "Query timeout (seconds)", "value": toStr(p.QueryTimeout) }
	params["cacheSelectNodes"] = map[string]interface{}{ "name": "cacheSelectNodes", "label": "Cache select nodes", "value": toStr(p.CacheSelectNodes) }
	params["rdfsSubClassReasoning"] = map[string]interface{}{ "name": "rdfsSubClassReasoning", "label": "RDFS subClass reasoning", "value": toStr(p.RdfsSubClassReasoning) }
	params["validationEnabled"] = map[string]interface{}{ "name": "validationEnabled", "label": "Enable the SHACL validation", "value": toStr(p.ValidationEnabled) }
	params["ftsStringLiteralsIndex"] = map[string]interface{}{ "name": "ftsStringLiteralsIndex", "label": "FTS index for xsd:string literals", "value": p.FtsStringLiteralsIndex }
	params["shapesGraph"] = map[string]interface{}{ "name": "shapesGraph", "label": "Named graphs for SHACL shapes", "value": p.ShapesGraph }
	params["parallelValidation"] = map[string]interface{}{ "name": "parallelValidation", "label": "Run parallel validation", "value": toStr(p.ParallelValidation) }
	params["checkForInconsistencies"] = map[string]interface{}{ "name": "checkForInconsistencies", "label": "Enable consistency checks", "value": toStr(p.CheckForInconsistencies) }
	params["performanceLogging"] = map[string]interface{}{ "name": "performanceLogging", "label": "Log the execution time per shape", "value": toStr(p.PerformanceLogging) }
	params["disableSameAs"] = map[string]interface{}{ "name": "disableSameAs", "label": "Disable owl:sameAs", "value": toStr(p.DisableSameAs) }
	params["ftsIrisIndex"] = map[string]interface{}{ "name": "ftsIrisIndex", "label": "FTS index for full-text indexing of IRIs", "value": p.FtsIrisIndex }
	params["entityIndexSize"] = map[string]interface{}{ "name": "entityIndexSize", "label": "Entity index size", "value": toStr(p.EntityIndexSize) }
	params["dashDataShapes"] = map[string]interface{}{ "name": "dashDataShapes", "label": "DASH data shapes extensions", "value": toStr(p.DashDataShapes) }
	params["queryLimitResults"] = map[string]interface{}{ "name": "queryLimitResults", "label": "Limit query results", "value": toStr(p.QueryLimitResults) }
	params["throwQueryEvaluationExceptionOnTimeout"] = map[string]interface{}{ "name": "throwQueryEvaluationExceptionOnTimeout", "label": "Throw exception on query timeout", "value": toStr(p.ThrowQueryEvaluationExceptionOnTimeout) }
	params["storageFolder"] = map[string]interface{}{ "name": "storageFolder", "label": "Storage folder", "value": p.StorageFolder }
	params["validationResultsLimitPerConstraint"] = map[string]interface{}{ "name": "validationResultsLimitPerConstraint", "label": "Validation results limit per constraint", "value": toStr(p.ValidationResultsLimitPerConstraint) }
	params["enablePredicateList"] = map[string]interface{}{ "name": "enablePredicateList", "label": "Enable predicate list index", "value": toStr(p.EnablePredicateList) }
	params["transactionalValidationLimit"] = map[string]interface{}{ "name": "transactionalValidationLimit", "label": "Transactional validation limit", "value": p.TransactionalValidationLimit }
	params["ftsIndexes"] = map[string]interface{}{ "name": "ftsIndexes", "label": "FTS indexes to build (comma delimited)", "value": toCSV(p.FtsIndexes) }
	params["logValidationPlans"] = map[string]interface{}{ "name": "logValidationPlans", "label": "Log the executed validation plans", "value": toStr(p.LogValidationPlans) }
	params["imports"] = map[string]interface{}{ "name": "imports", "label": "Imported RDF files(';' delimited)", "value": p.Imports }
	params["isShacl"] = map[string]interface{}{ "name": "isShacl", "label": "Enable SHACL validation", "value": toStr(p.IsShacl) }
	params["inMemoryLiteralProperties"] = map[string]interface{}{ "name": "inMemoryLiteralProperties", "label": "Cache literal language tags", "value": toStr(p.InMemoryLiteralProperties) }
	params["ruleset"] = map[string]interface{}{ "name": "ruleset", "label": "Ruleset", "value": p.Ruleset }
	params["readOnly"] = map[string]interface{}{ "name": "readOnly", "label": "Read-only", "value": toStr(p.ReadOnly) }
	params["enableLiteralIndex"] = map[string]interface{}{ "name": "enableLiteralIndex", "label": "Enable literal index", "value": toStr(p.EnableLiteralIndex) }
	params["enableFtsIndex"] = map[string]interface{}{ "name": "enableFtsIndex", "label": "Enable full-text search (FTS) index", "value": toStr(p.EnableFtsIndex) }
	params["defaultNS"] = map[string]interface{}{ "name": "defaultNS", "label": "Default namespaces for imports(';' delimited)", "value": p.DefaultNS }
	params["enableContextIndex"] = map[string]interface{}{ "name": "enableContextIndex", "label": "Enable context index", "value": toStr(p.EnableContextIndex) }
	params["baseURL"] = map[string]interface{}{ "name": "baseURL", "label": "Base URL", "value": p.BaseURL }
	params["logValidationViolations"] = map[string]interface{}{ "name": "logValidationViolations", "label": "Log validation violations", "value": toStr(p.LogValidationViolations) }
	params["globalLogValidationExecution"] = map[string]interface{}{ "name": "globalLogValidationExecution", "label": "Log every execution step of the SHACL validation", "value": toStr(p.GlobalLogValidationExecution) }
	params["entityIdSize"] = map[string]interface{}{ "name": "entityIdSize", "label": "Entity ID size", "value": toStr(p.EntityIdSize) }
	params["repositoryType"] = map[string]interface{}{ "name": "repositoryType", "label": "Repository type", "value": p.RepositoryType }
	params["eclipseRdf4jShaclExtensions"] = map[string]interface{}{ "name": "eclipseRdf4jShaclExtensions", "label": "RDF4J SHACL extensions", "value": toStr(p.EclipseRdf4jShaclExtensions) }
	params["validationResultsLimitTotal"] = map[string]interface{}{ "name": "validationResultsLimitTotal", "label": "Validation results limit total", "value": toStr(p.ValidationResultsLimitTotal) }
	payload["params"] = params

	return payload
}
