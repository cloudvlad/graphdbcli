package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"graphdbcli/internal/tool_configurations/logging"
	"io"
	"net/http"

	"go.uber.org/zap"
)

// HandleCreate handles requests that are supposed to create repositories.
func HandleCreate(req *http.Request) {
	// Read the request body
	defer req.Body.Close()
	bodyBefore, _ := io.ReadAll(req.Body)
	var input map[string]interface{}
	if err := json.Unmarshal(bodyBefore, &input); err != nil {
		// handle error (could log or return)
		logging.LOGGER.Error("unable to unmarshal body", zap.Error(err))
		input = map[string]interface{}{}
	}

	// Log before conversion
	logging.LOGGER.Debug("body before conversion", zap.String("body", string(bodyBefore)))

	// Convert to old format
	oldFormat := FillExpectedFormat(input)

	bodyAfter, err := json.Marshal(oldFormat)
	if err == nil {
		// Log after conversion
		logging.LOGGER.Debug("body after conversion", zap.String("body", string(bodyAfter)))
		req.Body = io.NopCloser(bytes.NewReader(bodyAfter))
		req.ContentLength = int64(len(bodyAfter))
	}
}

// FillExpectedFormat takes a new API Payload format (even partial) and returns the fully populated old format.
func FillExpectedFormat(input map[string]interface{}) map[string]interface{} {
	paramTemplate := map[string]map[string]interface{}{
		"queryTimeout":                           {"name": "queryTimeout", "label": "Query timeout (seconds)", "value": "0"},
		"cacheSelectNodes":                       {"name": "cacheSelectNodes", "label": "Cache select nodes", "value": "true"},
		"rdfsSubClassReasoning":                  {"name": "rdfsSubClassReasoning", "label": "RDFS subClass reasoning", "value": "true"},
		"validationEnabled":                      {"name": "validationEnabled", "label": "Enable the SHACL validation", "value": "true"},
		"ftsStringLiteralsIndex":                 {"name": "ftsStringLiteralsIndex", "label": "FTS index for xsd:string literals", "value": "default"},
		"shapesGraph":                            {"name": "shapesGraph", "label": "Named graphs for SHACL shapes", "value": "http://rdf4j.org/schema/rdf4j#SHACLShapeGraph"},
		"parallelValidation":                     {"name": "parallelValidation", "label": "Run parallel validation", "value": "true"},
		"checkForInconsistencies":                {"name": "checkForInconsistencies", "label": "Enable consistency checks", "value": "false"},
		"performanceLogging":                     {"name": "performanceLogging", "label": "Log the execution time per shape", "value": "false"},
		"disableSameAs":                          {"name": "disableSameAs", "label": "Disable owl:sameAs", "value": "true"},
		"ftsIrisIndex":                           {"name": "ftsIrisIndex", "label": "FTS index for full-text indexing of IRIs", "value": "en"},
		"entityIndexSize":                        {"name": "entityIndexSize", "label": "Entity index size", "value": "10000000"},
		"dashDataShapes":                         {"name": "dashDataShapes", "label": "DASH data shapes extensions", "value": "true"},
		"queryLimitResults":                      {"name": "queryLimitResults", "label": "Limit query results", "value": "0"},
		"throwQueryEvaluationExceptionOnTimeout": {"name": "throwQueryEvaluationExceptionOnTimeout", "label": "Throw exception on query timeout", "value": "false"},
		"member":                                 {"name": "member", "label": "FedX repo members", "value": []interface{}{}},
		"storageFolder":                          {"name": "storageFolder", "label": "Storage folder", "value": "storage"},
		"validationResultsLimitPerConstraint":    {"name": "validationResultsLimitPerConstraint", "label": "Validation results limit per constraint", "value": "1000"},
		"enablePredicateList":                    {"name": "enablePredicateList", "label": "Enable predicate list index", "value": "true"},
		"transactionalValidationLimit":           {"name": "transactionalValidationLimit", "label": "Transactional validation limit", "value": "500000"},
		"ftsIndexes":                             {"name": "ftsIndexes", "label": "FTS indexes to build (comma delimited)", "value": "default, iri, en"},
		"logValidationPlans":                     {"name": "logValidationPlans", "label": "Log the executed validation plans", "value": "false"},
		"imports":                                {"name": "imports", "label": "Imported RDF files(';' delimited)", "value": ""},
		"isShacl":                                {"name": "isShacl", "label": "Enable SHACL validation", "value": "false"},
		"inMemoryLiteralProperties":              {"name": "inMemoryLiteralProperties", "label": "Cache literal language tags", "value": "true"},
		"ruleset":                                {"name": "ruleset", "label": "Ruleset", "value": "rdfsplus-optimized"},
		"readOnly":                               {"name": "readOnly", "label": "Read-only", "value": "false"},
		"enableLiteralIndex":                     {"name": "enableLiteralIndex", "label": "Enable literal index", "value": "true"},
		"enableFtsIndex":                         {"name": "enableFtsIndex", "label": "Enable full-text search (FTS) index", "value": "false"},
		"defaultNS":                              {"name": "defaultNS", "label": "Default namespaces for imports(';' delimited)", "value": ""},
		"enableContextIndex":                     {"name": "enableContextIndex", "label": "Enable context index", "value": "false"},
		"baseURL":                                {"name": "baseURL", "label": "Base URL", "value": "http://example.org/owlim#"},
		"logValidationViolations":                {"name": "logValidationViolations", "label": "Log validation violations", "value": "false"},
		"globalLogValidationExecution":           {"name": "globalLogValidationExecution", "label": "Log every execution step of the SHACL validation", "value": "false"},
		"entityIdSize":                           {"name": "entityIdSize", "label": "Entity ID size", "value": "32"},
		"repositoryType":                         {"name": "repositoryType", "label": "Repository type", "value": "file-repository"},
		"eclipseRdf4jShaclExtensions":            {"name": "eclipseRdf4jShaclExtensions", "label": "RDF4J SHACL extensions", "value": "true"},
		"validationResultsLimitTotal":            {"name": "validationResultsLimitTotal", "label": "Validation results limit total", "value": "1000000"},
	}

	// Start with the template for the top-level object
	oldTemplate := map[string]interface{}{
		"id":         "",
		"location":   "",
		"title":      "",
		"type":       "graphdb",
		"sesameType": "graphdb:SailRepository",
	}

	// Copy top-level fields from input if present
	for _, key := range []string{"id", "location", "title", "type"} {
		if v, ok := input[key]; ok {
			oldTemplate[key] = v
		}
	}

	// Always set sesameType
	oldTemplate["sesameType"] = "graphdb:SailRepository"

	// Merge params
	params := map[string]map[string]interface{}{}
	toString := func(val interface{}) string {
		if val == nil {
			return ""
		}
		return fmt.Sprintf("%v", val)
	}

	if inputParams, ok := input["params"].(map[string]interface{}); ok {
		for k, v := range paramTemplate {
			params[k] = map[string]interface{}{
				"name":  v["name"],
				"label": v["label"],
				"value": toString(v["value"]),
			}
			if _, exists := inputParams[k]; exists {
				if k == "member" {
					arr, ok := inputParams[k].([]interface{})
					if ok {
						params[k]["value"] = arr
					} else {
						params[k]["value"] = []interface{}{}
					}
				} else {
					params[k]["value"] = toString(inputParams[k])
				}
			}
		}
	} else {
		// No input params, use all defaults
		for k, v := range paramTemplate {
			// `member` is the only JSON list value, and it requires special handling
			if k == "member" {
				params[k] = map[string]interface{}{
					"name":  v["name"],
					"label": v["label"],
					"value": []interface{}{},
				}
			} else {
				params[k] = map[string]interface{}{
					"name":  v["name"],
					"label": v["label"],
					"value": toString(v["value"]),
				}
			}
		}
	}

	oldTemplate["params"] = params

	return oldTemplate
}
