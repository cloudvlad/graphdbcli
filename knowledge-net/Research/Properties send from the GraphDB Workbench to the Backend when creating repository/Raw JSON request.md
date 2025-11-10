```json
{
  "id": "vlads-repo",
  "params": {
    "queryTimeout": {
      "name": "queryTimeout",
      "label": "Query timeout (seconds)",
      "value": 0,
      "isNumber": true
    },
    "cacheSelectNodes": {
      "name": "cacheSelectNodes",
      "label": "Cache select nodes",
      "value": "true"
    },
    "rdfsSubClassReasoning": {
      "name": "rdfsSubClassReasoning",
      "label": "RDFS subClass reasoning",
      "value": "true"
    },
    "validationEnabled": {
      "name": "validationEnabled",
      "label": "Enable the SHACL validation",
      "value": "true"
    },
    "ftsStringLiteralsIndex": {
      "name": "ftsStringLiteralsIndex",
      "label": "FTS index for xsd:string literals",
      "value": "default"
    },
    "shapesGraph": {
      "name": "shapesGraph",
      "label": "Named graphs for SHACL shapes",
      "value": "http://rdf4j.org/schema/rdf4j#SHACLShapeGraph"
    },
    "parallelValidation": {
      "name": "parallelValidation",
      "label": "Run parallel validation",
      "value": "true"
    },
    "title": {
      "name": "title",
      "label": "Repository description",
      "value": ""
    },
    "checkForInconsistencies": {
      "name": "checkForInconsistencies",
      "label": "Enable consistency checks",
      "value": "false"
    },
    "performanceLogging": {
      "name": "performanceLogging",
      "label": "Log the execution time per shape",
      "value": "true"
    },
    "disableSameAs": {
      "name": "disableSameAs",
      "label": "Disable owl:sameAs",
      "value": "true"
    },
    "ftsIrisIndex": {
      "name": "ftsIrisIndex",
      "label": "FTS index for full-text indexing of IRIs",
      "value": "none"
    },
    "entityIndexSize": {
      "name": "entityIndexSize",
      "label": "Entity index size",
      "value": 10000000,
      "isNumber": true
    },
    "dashDataShapes": {
      "name": "dashDataShapes",
      "label": "DASH data shapes extensions",
      "value": "true"
    },
    "queryLimitResults": {
      "name": "queryLimitResults",
      "label": "Limit query results",
      "value": 0,
      "isNumber": true
    },
    "throwQueryEvaluationExceptionOnTimeout": {
      "name": "throwQueryEvaluationExceptionOnTimeout",
      "label": "Throw exception on query timeout",
      "value": "false"
    },
    "id": {
      "name": "id",
      "label": "Repository ID",
      "value": "repo-test"
    },
    "storageFolder": {
      "name": "storageFolder",
      "label": "Storage folder",
      "value": "storage"
    },
    "validationResultsLimitPerConstraint": {
      "name": "validationResultsLimitPerConstraint",
      "label": "Validation results limit per constraint",
      "value": 1000,
      "isNumber": true
    },
    "enablePredicateList": {
      "name": "enablePredicateList",
      "label": "Enable predicate list index",
      "value": "true"
    },
    "transactionalValidationLimit": {
      "name": "transactionalValidationLimit",
      "label": "Transactional validation limit",
      "value": "500000"
    },
    "ftsIndexes": {
      "name": "ftsIndexes",
      "label": "FTS indexes to build (comma delimited)",
      "value": "default, iri"
    },
    "logValidationPlans": {
      "name": "logValidationPlans",
      "label": "Log the executed validation plans",
      "value": "true"
    },
    "imports": {
      "name": "imports",
      "label": "Imported RDF files(';' delimited)",
      "value": ""
    },
    "inMemoryLiteralProperties": {
      "name": "inMemoryLiteralProperties",
      "label": "Cache literal language tags",
      "value": "true"
    },
    "isShacl": {
      "name": "isShacl",
      "label": "Enable SHACL validation",
      "value": "true"
    },
    "ruleset": {
      "name": "ruleset",
      "label": "Ruleset",
      "value": "/home/vlad/.gdb/instances/jolly_hertz/work/tmp/graphdb/builtin_owl2-qltmp1761583708338.pie"
    },
    "readOnly": {
      "name": "readOnly",
      "label": "Read-only",
      "value": "false"
    },
    "enableFtsIndex": {
      "name": "enableFtsIndex",
      "label": "Enable full-text search (FTS) index",
      "value": "false"
    },
    "enableLiteralIndex": {
      "name": "enableLiteralIndex",
      "label": "Enable literal index",
      "value": "true"
    },
    "enableContextIndex": {
      "name": "enableContextIndex",
      "label": "Enable context index",
      "value": "false"
    },
    "defaultNS": {
      "name": "defaultNS",
      "label": "Default namespaces for imports(';' delimited)",
      "value": ""
    },
    "baseURL": {
      "name": "baseURL",
      "label": "Base URL",
      "value": "http://example.org/owlim#"
    },
    "logValidationViolations": {
      "name": "logValidationViolations",
      "label": "Log validation violations",
      "value": "true"
    },
    "globalLogValidationExecution": {
      "name": "globalLogValidationExecution",
      "label": "Log every execution step of the SHACL validation",
      "value": "true"
    },
    "entityIdSize": {
      "name": "entityIdSize",
      "label": "Entity ID size",
      "value": "32"
    },
    "repositoryType": {
      "name": "repositoryType",
      "label": "Repository type",
      "value": "file-repository"
    },
    "eclipseRdf4jShaclExtensions": {
      "name": "eclipseRdf4jShaclExtensions",
      "label": "RDF4J SHACL extensions",
      "value": "true"
    },
    "validationResultsLimitTotal": {
      "name": "validationResultsLimitTotal",
      "label": "Validation results limit total",
      "value": 1000000,
      "isNumber": true
    }
  },
  "title": "",
  "type": "graphdb",
  "location": ""
}
```


curl -X PUT <base_url>/rest/repositories/<repo_id> -H 'Accept: application/json' -H 'Content-Type: application/json' -d  ''
