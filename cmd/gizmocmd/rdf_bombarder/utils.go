package rdf_bombarder

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func bombard(instanceAddress, repoName string, numberOfNamedGraphs, numberOfStatementsPerNamedGraph, numberOfThreads, numberOfStatementsPerRequest int, interconnection float64) {
	addr := strings.TrimRight(instanceAddress, "/")
	endpoint := fmt.Sprintf("%s/repositories/%s/statements", addr, repoName)
	client := &http.Client{Timeout: 10 * time.Second}
	jobs := make(chan int, numberOfNamedGraphs)
	done := make(chan struct{})

	// Make worker instances
	worker := makeWorker(client, endpoint, numberOfStatementsPerNamedGraph, numberOfStatementsPerRequest, interconnection, done, jobs)

	// Start workers
	for t := 0; t < numberOfThreads; t++ {
		go worker()
	}

	// Send jobs to the jobs channel (mapped to the named graphs).
	// Each worker consumes one job at a time.
	for i := 0; i < numberOfNamedGraphs; i++ {
		jobs <- i
	}

	close(jobs)

	// Wait for all workers to finish
	// Waiting to consume from the channel the exact number of
	// threads. Publishing in the channel happens only when a thread is done
	// with its job.
	for t := 0; t < numberOfThreads; t++ {
		<-done
	}
}
