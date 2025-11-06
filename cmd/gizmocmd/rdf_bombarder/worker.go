package rdf_bombarder

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// makeWorker returns a worker function for bombard
func makeWorker(client *http.Client, endpoint string, numberOfStatementsPerNamedGraph, numberOfStatementsPerRequest int, interconnection float64, done chan struct{}, jobs chan int) func() {
	return func() {
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
				sparql := randomInsertQuery(i, sent, n, interconnection)
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
					fmt.Printf("[BOMBARD] Sent SPARQL insert for Graph %d, Status %d, Number of statements send %d\n", i, resp.StatusCode, n)
				}
				resp.Body.Close()
				sent += n
			}
		}
		done <- struct{}{}
	}
}
