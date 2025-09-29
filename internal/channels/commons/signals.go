// Package commons provides channels that are used for synchronization for many components and execution stages across the CLI.
package commons

// Success is used for notifying that a stage/operation was executed successfully.
var Success = make(chan bool, 1)

// Failure is used for notifying that a stage/operation failed.
var Failure = make(chan bool, 1)
