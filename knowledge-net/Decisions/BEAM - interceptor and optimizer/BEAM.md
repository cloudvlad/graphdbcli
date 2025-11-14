## Objective
Develop an optimized API that is used as interceptor of the request that are meant for GraphDB.
The interceptor is supposed to be able to handle improved and clean version of the API payload and
retransmit it toward to GraphDB API in the format it requires.

This is needed because the current format is full with fields, that are either not needed, not set by the user
or there are fields, that are just required just because. Some parameters, that are showed in TODO LINK PROPERTIES TABLE
are never passed, and some fields are rather unneeded when sending user data, like label and so on. Also each property is an object
on its own, and most of the properites are present as strings.

## Decision
- Create lightweight version of the API
- Reduce the number of required parametes where possible
- MAke it compatible with the GraphDB API
