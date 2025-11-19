---
created: 2025-11-19
last-updated: 2025-11-19
---
## 19/11/2025
Planned first sub-command of the super-command is `create` which is used to create a GraphDB repositories based on the repositories API.
Planned properties and types:

| Flag Name                                 | Type      |
|--------------------------------------------|-----------|
| eidSize40                                 | bool      |
| desc                                      | string    |
| readOnly                                  | bool      |
| ruleset                                   | string    |
| disableOwlSameAs                          | bool      |
| enableConsistencyChecks                   | bool      |
| enableShaclValidation                     | bool      |
| enableCacheSelectNodes                    | bool      |
| enableLogValidationPlans                  | bool      |
| enableParallelValidation                  | bool      |
| enablePerformanceLogging                  | bool      |
| enableDashDataShapes                      | bool      |
| enableLogValidationViolations             | bool      |
| enableGlobalLogValidationExecution        | bool      |
| enableEclipseRdf4jShaclExtensions         | bool      |
| validationResultsLimitTotal               | uint32    |
| validationResultsLimitPerConstraint       | uint32    |
| shapesGraph                               | string    |
| enableContextIndex                        | bool      |
| enablePredicateList                       | bool      |
| enableFtsIndex                            | bool      |
| entityIndexSize                           | uint32    |
| ftsIndexes                                | []string  |
| ftsStringLiteralsIndex                    | string    |
| ftsIrisIndex                              | string    |
| throwQueryEvaluationExceptionOnTimeout    | bool      |
| queryTimeout                              | uint32    |
| queryLimitResults                         | uint32    |
