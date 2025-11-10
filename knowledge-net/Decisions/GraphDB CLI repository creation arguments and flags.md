---
tags:
  - ecosystem/graphdbcli
  - product/graphdb
  - dt/internal
---
## Objective
Define the arguments and flags that will be used in the CLI for creating a repository.

## Method
1. Check what were the currently send repositories

[[Properties send from the GraphDB Workbench to the Backend when creating repository]]

2. Define naming, constraints, dependencies and so on
	- ID will be argument and required
	- Entity ID Size will be replaced with will be implemented by one flag, specifying if the repository should be
	- Boolean flags will be mapped directly as boolean
	- Each non-boolean field should have validation, evaluating the constraints
	- Ruleset is a "special" properti. For initial version supporting the default rulesets only is fine (later we should add the custom ones as well)
	- Log warnings if some properties are overriden, but the feature is not enabled
		- SHACL
		- FTS
## Observations

## Conclusion

## Derived Axioms