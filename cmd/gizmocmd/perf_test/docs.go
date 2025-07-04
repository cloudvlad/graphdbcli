package perf_test

var shortDescription = `Tests the performance on given data by executing tests`

var longDescription = `Tests the performance on given data that is preseeded
beforehand. Those are tipically Insert SPARQL quriest, that are send for processing.
Tests are tipically some kind of Selects quriest. The resposne for each querty is stored
and added to the statitics for each iterations and more general idea is created.

Statistics are printed in the terminal for the prepared data and for the test data. The main
difference is that only tests could be executed multiple times, while seeing the instance are done once.`

var examples = `
# Runs the insert queries from the specified load directory in the specified instance address
# and repository and executes the test queries from the specified directory
perf_test -a http://localhost:7200 -r test -l ./data_insert_queries -t ./test_queries
# Executes the each test two times before moving to the next one
perf_test -a http://localhost:7200 -r test -l ./data_insert_queries -t ./test_queries --runs 2
# Stores the results in timestamped markdown file (table formatted) 
perf_test -a http://localhost:7200 -r test -l ./data_insert_queries -t ./test_queries --md
`
