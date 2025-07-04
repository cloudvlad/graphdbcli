package join

import "strings"

var electionMinTimeout uint
var electionRangeTimeout uint
var heartbeatInterval uint
var messageSizeKB uint
var transactionLogMaximumSizeGB uint
var verificationTimeout uint
var instances []string

func joinInstances() {
	checkInstancesHealth()
	formCluster()
}

func checkInstancesHealth() {
	// add cehck for each instance in the instances string array
	//it shoul check the address + /protocol endpoint
	//if returned 200 proceed
	//if not, wait
}

func formCluster() {
	// form cluster using the curl based approach from the docs
}

func setInstances(instancesArg string) {
	instances = strings.Split(instancesArg, ",")
}
