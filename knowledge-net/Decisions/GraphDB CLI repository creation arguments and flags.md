---
tags:
  - cmd/repository
---
## Objective
Define the arguments and flags that will be used in the CLI for creating a repository.

## Decision
1. Check what were the currently send repositories [[Properties send from the GraphDB Workbench to the Backend when creating repository]]
2. Define naming, constraints, dependencies and so on
	- ID will be argument and required
	- Entity ID Size will be replaced with will be implemented by one flag, specifying if  should be 40-bits (the default will be 32-bit)
	- Boolean flags will be mapped directly as boolean local cobra flags
	- Each non-boolean field should have validation, evaluating the constraints
	- Ruleset is a "special" property. For initial version supporting the default rulesets only is fine (later we should add the custom ones as well)
	- Log warnings if some properties are overriden, but the feature is not enabled
		- SHACL
		- FTS
3. Required and unused parameters should be set, so we are comliant with the current API version