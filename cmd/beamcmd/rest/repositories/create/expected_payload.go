package create

type Param struct {
	Name     string      `json:"name"`
	Label    string      `json:"label"`
	Value    interface{} `json:"value"`
	IsNumber bool        `json:"isNumber,omitempty"`
}

type ExpectedPayload struct {
	ID       string           `json:"id"`
	Params   map[string]Param `json:"params"`
	Title    string           `json:"title"`
	Type     string           `json:"type"`
	Location string           `json:"location"`
}

// ExampleExpectedPayload returns an example ExpectedPayload populated with sample data.
func ExampleExpectedPayload() ExpectedPayload {
	return ExpectedPayload{
		ID:       "",
		Title:    "",
		Type:     "graphdb",
		Location: "",
		Params: map[string]Param{
			"queryTimeout":                           {Name: "queryTimeout", Label: "Query timeout (seconds)", Value: 0, IsNumber: true},
			"cacheSelectNodes":                       {Name: "cacheSelectNodes", Label: "Cache select nodes", Value: "true"},
			"rdfsSubClassReasoning":                  {Name: "rdfsSubClassReasoning", Label: "RDFS subClass reasoning", Value: "true"},
			"validationEnabled":                      {Name: "validationEnabled", Label: "Enable the SHACL validation", Value: "true"},
			"ftsStringLiteralsIndex":                 {Name: "ftsStringLiteralsIndex", Label: "FTS index for xsd:string literals", Value: "default"},
			"shapesGraph":                            {Name: "shapesGraph", Label: "Named graphs for SHACL shapes", Value: "http://rdf4j.org/schema/rdf4j#SHACLShapeGraph"},
			"parallelValidation":                     {Name: "parallelValidation", Label: "Run parallel validation", Value: "true"},
			"title":                                  {Name: "title", Label: "Repository description", Value: ""},
			"checkForInconsistencies":                {Name: "checkForInconsistencies", Label: "Enable consistency checks", Value: "false"},
			"performanceLogging":                     {Name: "performanceLogging", Label: "Log the execution time per shape", Value: "false"},
			"disableSameAs":                          {Name: "disableSameAs", Label: "Disable owl:sameAs", Value: "true"},
			"ftsIrisIndex":                           {Name: "ftsIrisIndex", Label: "FTS index for full-text indexing of IRIs", Value: "none"},
			"entityIndexSize":                        {Name: "entityIndexSize", Label: "Entity index size", Value: 10000000, IsNumber: true},
			"dashDataShapes":                         {Name: "dashDataShapes", Label: "DASH data shapes extensions", Value: "true"},
			"queryLimitResults":                      {Name: "queryLimitResults", Label: "Limit query results", Value: 0, IsNumber: true},
			"throwQueryEvaluationExceptionOnTimeout": {Name: "throwQueryEvaluationExceptionOnTimeout", Label: "Throw exception on query timeout", Value: "false"},
			"id":                                     {Name: "id", Label: "Repository ID", Value: "repo-test"},
			"storageFolder":                          {Name: "storageFolder", Label: "Storage folder", Value: "storage"},
			"validationResultsLimitPerConstraint":    {Name: "validationResultsLimitPerConstraint", Label: "Validation results limit per constraint", Value: 1000, IsNumber: true},
			"enablePredicateList":                    {Name: "enablePredicateList", Label: "Enable predicate list index", Value: "true"},
			"transactionalValidationLimit":           {Name: "transactionalValidationLimit", Label: "Transactional validation limit", Value: "500000"},
			"ftsIndexes":                             {Name: "ftsIndexes", Label: "FTS indexes to build (comma delimited)", Value: "default, iri"},
			"logValidationPlans":                     {Name: "logValidationPlans", Label: "Log the executed validation plans", Value: "false"},
			"imports":                                {Name: "imports", Label: "Imported RDF files(';' delimited)", Value: ""},
			"inMemoryLiteralProperties":              {Name: "inMemoryLiteralProperties", Label: "Cache literal language tags", Value: "true"},
			"isShacl":                                {Name: "isShacl", Label: "Enable SHACL validation", Value: "false"},
			"ruleset":                                {Name: "ruleset", Label: "Ruleset", Value: "rdfsplus-optimized"},
			"readOnly":                               {Name: "readOnly", Label: "Read-only", Value: "false"},
			"enableFtsIndex":                         {Name: "enableFtsIndex", Label: "Enable full-text search (FTS) index", Value: "false"},
			"enableLiteralIndex":                     {Name: "enableLiteralIndex", Label: "Enable literal index", Value: "true"},
			"enableContextIndex":                     {Name: "enableContextIndex", Label: "Enable context index", Value: "false"},
			"defaultNS":                              {Name: "defaultNS", Label: "Default namespaces for imports(';' delimited)", Value: ""},
			"baseURL":                                {Name: "baseURL", Label: "Base URL", Value: "http://example.org/owlim#"},
			"logValidationViolations":                {Name: "logValidationViolations", Label: "Log validation violations", Value: "false"},
			"globalLogValidationExecution":           {Name: "globalLogValidationExecution", Label: "Log every execution step of the SHACL validation", Value: "false"},
			"entityIdSize":                           {Name: "entityIdSize", Label: "Entity ID size", Value: "32"},
			"repositoryType":                         {Name: "repositoryType", Label: "Repository type", Value: "file-repository"},
			"eclipseRdf4jShaclExtensions":            {Name: "eclipseRdf4jShaclExtensions", Label: "RDF4J SHACL extensions", Value: "true"},
			"validationResultsLimitTotal":            {Name: "validationResultsLimitTotal", Label: "Validation results limit total", Value: 1000000, IsNumber: true},
		},
	}
}
