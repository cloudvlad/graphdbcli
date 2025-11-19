---
version:
  - 11.1.2
  - 11.1.3
---
## Objective

Develop an optimized API to serve as a proxy for requests intended for GraphDB. The proxy should be able to intercept an improved, cleaner version of the API payload, transform it, and retransmit it to the GraphDB API in the format it requires.

This is needed because the current payload format contains many fields that are unnecessary, unused, or required only due to legacy reasons. Some parameters listed below are never passed, while other fields, such as labels, are not relevant when sending user data, but are required for parsing the requests. In addition, each property is currently represented as a separate object, even though most properties are essentially simple strings.

![[Properties send from the GraphDB Workbench to the Backend when creating repository#^gdb-11-1-2-properties]]

## Decision
- Create a syntactically payload-lightweight version of the API.
- Reduce the number of required parameters where possible.
- Ensure full compatibility with the GraphDB API.