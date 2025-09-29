package perf_test

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var instanceAddress string
var loaderQueriesDir string
var testQueriesDir string
var repoName string
var runs int
var markdownResultTable bool

type ResultStats struct {
	TestName string
	Times    []time.Duration
}

var loaderResults []ResultStats
var testResults []ResultStats

func ContextIndexTests() {
	if loaderQueriesDir != "" {
		executeLoaders()
	}

	executeTests()
	printResults()

	if markdownResultTable {
		now := time.Now().Format("2006-01-02_15-04-05")
		filename := fmt.Sprintf("results_%s.md", now)
		content := MarkdownResultsTableFromResults()
		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Println("Error writing markdown results file:", err)
		} else {
			fmt.Println("Markdown results table written to:", filename)
		}
	}
}

func executeLoaders() {
	files, err := os.ReadDir(loaderQueriesDir)
	if err != nil {
		fmt.Println("Error reading insert queries dir:", err)
		return
	}
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		queryPath := filepath.Join(loaderQueriesDir, file.Name())
		queryBytes, err := os.ReadFile(queryPath)
		if err != nil {
			fmt.Println("Error reading query file:", queryPath, err)
			continue
		}
		url := fmt.Sprintf("%s/repositories/%s/statements", instanceAddress, repoName)
		req, err := http.NewRequest("POST", url, strings.NewReader(string(queryBytes)))
		if err != nil {
			fmt.Println("Error creating request:", err)
			continue
		}
		req.Header.Set("accept", "*/*")
		req.Header.Set("Content-Type", "application/sparql-update")
		start := time.Now()
		resp, err := http.DefaultClient.Do(req)
		duration := time.Since(start)
		if err != nil {
			fmt.Println("Error executing insert:", err)
			continue
		}

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			fmt.Printf("Error: insert for %s returned status %d\n", file.Name(), resp.StatusCode)
			resp.Body.Close()
			continue
		}
		defer resp.Body.Close()
		fmt.Printf("[LOAD] %s - %v\n", file.Name(), duration)
		loaderResults = append(loaderResults, ResultStats{TestName: file.Name(), Times: []time.Duration{duration}})
	}
}

func executeTests() {
	files, err := os.ReadDir(testQueriesDir)
	if err != nil {
		fmt.Println("Error reading select queries dir:", err)
		return
	}
	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		queryPath := filepath.Join(testQueriesDir, file.Name())
		queryBytes, err := os.ReadFile(queryPath)
		if err != nil {
			fmt.Println("Error reading query file:", queryPath, err)
			continue
		}
		times := make([]time.Duration, 0, runs)
		for i := 0; i < runs; i++ {
			url := fmt.Sprintf("%s/repositories/%s", instanceAddress, repoName)
			req, err := http.NewRequest("POST", url, strings.NewReader(string(queryBytes)))
			if err != nil {
				fmt.Println("Error creating request:", err)
				continue
			}
			req.Header.Set("accept", "application/sparql-results+xml")
			req.Header.Set("Content-Type", "application/sparql-query")
			start := time.Now()
			resp, err := http.DefaultClient.Do(req)
			duration := time.Since(start)
			if err != nil {
				fmt.Println("Error executing select:", err)
				continue
			}
			resp.Body.Close()
			times = append(times, duration)
			fmt.Printf("[TEST] %s (run %d/%d) - %v\n", file.Name(), i+1, runs, duration)
		}
		testResults = append(testResults, ResultStats{TestName: file.Name(), Times: times})
	}
}

func printResults() {
	// Helper to format duration as seconds with 4 digits before and 3 after the decimal, right-aligned
	formatDuration := func(d time.Duration) string {
		sec := float64(d.Nanoseconds()) / 1e9
		// 4 digits before, 3 after, right-aligned in 8 chars, e.g. "  1.234 s"
		return fmt.Sprintf("%8.3f s", sec)
	}

	fmt.Println("\nInsert query timings:")
	fmt.Printf("%-40s %15s %15s %15s\n", "Test Name", "Min", "Avg", "Max")
	fmt.Println(strings.Repeat("-", 90))
	for _, r := range loaderResults {
		min, max, sum := r.Times[0], r.Times[0], time.Duration(0)
		for _, t := range r.Times {
			if t < min {
				min = t
			}
			if t > max {
				max = t
			}
			sum += t
		}
		avg := sum / time.Duration(len(r.Times))
		fmt.Printf("%-40s %15s %15s %15s\n", r.TestName, formatDuration(min), formatDuration(avg), formatDuration(max))
	}
	fmt.Println(strings.Repeat("-", 90))

	fmt.Println("\nSelect query timings:")
	fmt.Printf("%-40s %15s %15s %15s\n", "Test Name", "Min", "Avg", "Max")
	fmt.Println(strings.Repeat("-", 90))
	for _, r := range testResults {
		min, max, sum := r.Times[0], r.Times[0], time.Duration(0)
		for _, t := range r.Times {
			if t < min {
				min = t
			}
			if t > max {
				max = t
			}
			sum += t
		}
		avg := sum / time.Duration(len(r.Times))
		fmt.Printf("%-40s %15s %15s %15s\n", r.TestName, formatDuration(min), formatDuration(avg), formatDuration(max))
	}
	fmt.Println(strings.Repeat("-", 90))
}

// MarkdownResultsTableFromResults generates a markdown table from your test results with proper column widths.
func MarkdownResultsTableFromResults() string {
	columns := []string{"Test Name", "Min (s)", "Avg (s)", "Max (s)"}
	// Calculate column widths
	colWidths := make([]int, len(columns))
	for i, col := range columns {
		colWidths[i] = len(col)
	}
	// Check data for width
	for _, r := range testResults {
		if len(r.TestName) > colWidths[0] {
			colWidths[0] = len(r.TestName)
		}
		// Min, Avg, Max
		min, max, sum := r.Times[0], r.Times[0], time.Duration(0)
		for _, t := range r.Times {
			if t < min {
				min = t
			}
			if t > max {
				max = t
			}
			sum += t
		}
		avg := sum.Seconds() / float64(len(r.Times))
		minStr := fmt.Sprintf("%.3f", min.Seconds())
		avgStr := fmt.Sprintf("%.3f", avg)
		maxStr := fmt.Sprintf("%.3f", max.Seconds())
		if len(minStr) > colWidths[1] {
			colWidths[1] = len(minStr)
		}
		if len(avgStr) > colWidths[2] {
			colWidths[2] = len(avgStr)
		}
		if len(maxStr) > colWidths[3] {
			colWidths[3] = len(maxStr)
		}
	}
	// Build header
	var sb strings.Builder
	sb.WriteString("|")
	for i, col := range columns {
		sb.WriteString(" " + fmt.Sprintf("%-*s", colWidths[i], col) + " |")
	}
	sb.WriteString("\n|")
	for _, w := range colWidths {
		sb.WriteString(" " + strings.Repeat("-", w) + " |")
	}
	sb.WriteString("\n")
	// Build rows
	for _, r := range testResults {
		min, max, sum := r.Times[0], r.Times[0], time.Duration(0)
		for _, t := range r.Times {
			if t < min {
				min = t
			}
			if t > max {
				max = t
			}
			sum += t
		}
		avg := sum.Seconds() / float64(len(r.Times))
		minStr := fmt.Sprintf("%.3f", min.Seconds())
		avgStr := fmt.Sprintf("%.3f", avg)
		maxStr := fmt.Sprintf("%.3f", max.Seconds())
		sb.WriteString("|")
		sb.WriteString(" " + fmt.Sprintf("%-*s", colWidths[0], r.TestName) + " |")
		sb.WriteString(" " + fmt.Sprintf("%-*s", colWidths[1], minStr) + " |")
		sb.WriteString(" " + fmt.Sprintf("%-*s", colWidths[2], avgStr) + " |")
		sb.WriteString(" " + fmt.Sprintf("%-*s", colWidths[3], maxStr) + " |\n")
	}
	return sb.String()
}
