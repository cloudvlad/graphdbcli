---
tags:
  - api
  - repository
version: 11.1.2
---
## Objective
Find out what are the properties that are send from the workbench to the backend when creating a repository and what are their constraints/limits.

## Method
1. Create GraphDB instance
2. Check default properties and their values returned from  [/rest/repositories/default-config/graphdb](http://localhost:7200/rest/repositories/default-config/graphdb)
3. Map the Workbench GraphDB repository creation flags/arguments to the one from the Step 2.

## Observations

| Property name                          | Default value                                 | Type             | Constraints     | Configurable / Used |
| -------------------------------------- | --------------------------------------------- | ---------------- | --------------- | ------------------- |
| id                                     | -                                             | string           | Required        |                     |
| title                                  |                                               | string           |                 |                     |
| readOnly                               | false                                         | boolean          |                 |                     |
| ruleset                                | rdfsplus-optimized                            | string/file-path |                 |                     |
| disableSameAs                          | true                                          | boolean          |                 |                     |
| checkForInconsistencies                | false                                         | boolean          |                 |                     |
| isShacl                                | false                                         | boolean          |                 |                     |
| cacheSelectNodes                       | true                                          | boolean          |                 |                     |
| logValidationPlans                     | false                                         | boolean          |                 |                     |
| parallelValidation                     | true                                          | boolean          |                 |                     |
| performanceLogging                     | false                                         | boolean          |                 |                     |
| dashDataShapes                         | true                                          | boolean          |                 |                     |
| logValidationViolations                | false                                         | boolean          |                 |                     |
| globalLogValidationExecution           | false                                         | boolean          |                 |                     |
| eclipseRdf4jShaclExtensions            | true                                          | boolean          |                 |                     |
| validationResultsLimitTotal            | 1000000                                       | uint             |                 |                     |
| validationResultsLimitPerConstraint    | 1000                                          | uint             |                 |                     |
| shapesGraph                            | http://rdf4j.org/schema/rdf4j#SHACLShapeGraph | IRI              |                 |                     |
| entityIdSize                           | 32                                            | uint             | Either 32 or 40 |                     |
| enableContextIndex                     | false                                         | boolean          |                 |                     |
| enablePredicateList                    | true                                          | boolean          |                 |                     |
| enableFtsIndex                         | false                                         | boolean          |                 |                     |
| entityIndexSize                        | 10000000                                      | uint             |                 |                     |
| ftsIndexes                             | default, iri                                  | list             |                 |                     |
| ftsStringLiteralsIndex                 | default                                       | string           |                 |                     |
| ftsIrisIndex                           | none                                          | string           |                 |                     |
| throwQueryEvaluationExceptionOnTimeout | false                                         | boolean          |                 |                     |
| queryTimeout                           | 0                                             | uint             |                 |                     |
| queryLimitResults                      | 0                                             | uint             |                 |                     |
| enableLiteralIndex                     | true                                          | bool             |                 | no                  |
| rdfsSubClassReasoning                  | true                                          | bool             |                 | no                  |
| validationEnabled                      | true                                          | bool             |                 | no                  |
| storageFolder                          | "storage"                                     | string           |                 | no                  |
| transactionalValidationLimit           | 500000                                        | string           |                 | no                  |
| imports                                | ""                                            | string           |                 | no                  |
| inMemoryLiteralProperties              | true                                          | bool             |                 | no                  |
| defaultNS                              | ""                                            | string           |                 | no                  |
| baseURL                                | http://example.org/owlim#                     | IRI              |                 | no                  |
| repositoryType                         | file-repository                               | string           |                 | no                  |
| type                                   | graphdb                                       | string           |                 | no                  |
| location                               | ""                                            | string           |                 | no                  |

There was a property - `sesameType` , found [here](https://graphdb.ontotext.com/documentation/11.1/manage-repos-with-restapi.html#edit-a-repository-s-configuration) was is not part of the repository configurations from [here](https://graphdb.ontotext.com/documentation/11.1/configuring-a-repository.html#configuration-parameters).

^gdb-11-1-2-properties
## Conclusion
There are many properties, that are either not exposed for override, or not used internally.
