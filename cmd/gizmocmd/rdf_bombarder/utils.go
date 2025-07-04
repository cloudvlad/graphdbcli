package rdf_bombarder

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var instanceAddress string
var repoName string

var numberOfNamedGraphs int
var numberOfStatementsPerNamedGraph int
var numberOfThreads int
var numberOfStatementsPerRequest int
var interconnection float64

func bombard() {
	rand.Seed(time.Now().UnixNano())
	endpoint := fmt.Sprintf("%s/repositories/%s/statements", instanceAddress, repoName)
	client := &http.Client{Timeout: 10 * time.Second}
	jobs := make(chan int, numberOfNamedGraphs)
	done := make(chan struct{})

	   worker := func() {
		   for i := range jobs {
			   stmts := numberOfStatementsPerNamedGraph
			   batch := numberOfStatementsPerRequest
			   if batch < 1 {
				   batch = stmts
			   }
			   sent := 0
			   for sent < stmts {
				   n := batch
				   if sent+n > stmts {
					   n = stmts - sent
				   }
				   sparql := randomInsertQuery(i, sent, n, true)
				   form := url.Values{}
				   form.Set("update", sparql)

				   req, err := http.NewRequest("POST", endpoint, strings.NewReader(form.Encode()))
				   if err != nil {
					   fmt.Println("Error creating request:", err)
					   continue
				   }
				   req.Header.Set("accept", "*/*")
				   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				   resp, err := client.Do(req)
				   if err != nil {
					   fmt.Println("Error sending request:", err)
					   continue
				   }
				   if resp.StatusCode < 200 || resp.StatusCode >= 300 {
					   fmt.Printf("Error: insert returned status %d\n", resp.StatusCode)
				   } else {
					   fmt.Printf("[BOMBARD] Sent SPARQL insert for graph %d, status %d, stmts %d\n", i, resp.StatusCode, n)
				   }
				   resp.Body.Close()
				   sent += n
			   }
		   }
		   done <- struct{}{}
	   }

	// Start workers
	for t := 0; t < numberOfThreads; t++ {
		go worker()
	}

	// Send jobs
	for i := 0; i < numberOfNamedGraphs; i++ {
		jobs <- i
	}
	close(jobs)

	// Wait for all workers to finish
	for t := 0; t < numberOfThreads; t++ {
		<-done
	}
}

// randomInsertQuery generates a random SPARQL INSERT DATA query for a batch
func randomInsertQuery(graphIdx int, offset int, count int, useNamed bool) string {
   var b strings.Builder
   b.WriteString(fmt.Sprintf("INSERT DATA { GRAPH <urn:graph:%d> {\n", graphIdx))

   poolSize := offset + count
   subjects := make([]string, poolSize)
   objects := make([]string, poolSize)
   for i := 0; i < poolSize; i++ {
	   subjects[i] = fmt.Sprintf("<urn:e%d>", i+graphIdx*poolSize)
	   objects[i] = fmt.Sprintf("<urn:e%d>", i+graphIdx*poolSize)
   }
   predicates := []string{"<urn:p1>", "<urn:p2>", "<urn:p3>", "<urn:p4>", "<urn:p5>"}

   if interconnection <= 0.0 {
	   // Each entity gets one statement (subject, pred, object=subject)
	   for i := offset; i < offset+count; i++ {
		   subj := subjects[i]
		   pred := predicates[rand.Intn(len(predicates))]
		   obj := objects[i]
		   b.WriteString(fmt.Sprintf("  %s %s %s .\n", subj, pred, obj))
	   }
   } else if interconnection >= 1.0 {
	   // Every subject to every object (full mesh)
	   for i := offset; i < offset+count; i++ {
		   for j := 0; j < poolSize; j++ {
			   if i == j {
				   continue // skip self-link if desired
			   }
			   subj := subjects[i]
			   pred := predicates[rand.Intn(len(predicates))]
			   obj := objects[j]
			   b.WriteString(fmt.Sprintf("  %s %s %s .\n", subj, pred, obj))
		   }
	   }
   } else {
	   // Partial interconnection: for each subject, connect to a fraction of objects
	   linksPerSubject := int(float64(poolSize) * interconnection)
	   if linksPerSubject < 1 {
		   linksPerSubject = 1
	   }
	   for i := offset; i < offset+count; i++ {
		   used := map[int]bool{}
		   for l := 0; l < linksPerSubject; l++ {
			   var j int
			   for {
				   j = rand.Intn(poolSize)
				   if j != i && !used[j] {
					   used[j] = true
					   break
				   }
			   }
			   subj := subjects[i]
			   pred := predicates[rand.Intn(len(predicates))]
			   obj := objects[j]
			   b.WriteString(fmt.Sprintf("  %s %s %s .\n", subj, pred, obj))
		   }
	   }
   }
   b.WriteString("}\n}\n")
   return b.String()
}
