package statements_stats

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Stats struct {
	NumberOfStatements         int
	NumberOfExplicitStatements int
	NumberOfEntities           int
}

func parseOwlimProperties(path string) (Stats, error) {
	file, err := os.Open(path)
	if err != nil {
		return Stats{}, err
	}
	defer file.Close()

	stats := Stats{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "NumberOfStatements=") {
			val := strings.TrimPrefix(line, "NumberOfStatements=")
			stats.NumberOfStatements, _ = strconv.Atoi(val)
		} else if strings.HasPrefix(line, "NumberOfExplicitStatements=") {
			val := strings.TrimPrefix(line, "NumberOfExplicitStatements=")
			stats.NumberOfExplicitStatements, _ = strconv.Atoi(val)
		} else if strings.HasPrefix(line, "NumberOfEntities=") {
			val := strings.TrimPrefix(line, "NumberOfEntities=")
			stats.NumberOfEntities, _ = strconv.Atoi(val)
		}
	}
	return stats, scanner.Err()
}

func PrintRDFStatementsStats(root string) {
	entries, err := os.ReadDir(root)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	total := Stats{}
	fmt.Printf("%-50s %12s %20s %15s\n", "Repository", "Statements", "ExplicitStatements", "Entities")
	fmt.Println(strings.Repeat("-", 100))

	for _, entry := range entries {
		if entry.IsDir() {
			dirPath := filepath.Join(root, entry.Name())
			owlimPath := filepath.Join(dirPath, "owlim.properties")
			if _, err := os.Stat(owlimPath); os.IsNotExist(err) {
				fmt.Printf("%-50s %12s %20s %15s\n", entry.Name(), "SKIPPED", "SKIPPED", "SKIPPED")
				continue
			}
			stats, err := parseOwlimProperties(owlimPath)
			if err != nil {
				fmt.Printf("%-50s %12s %20s %15s\n", entry.Name(), "N/A", "N/A", "N/A")
				continue
			}
			fmt.Printf("%-50s %12s %20s %15s\n", entry.Name(),
				formatWithApostrophe(stats.NumberOfStatements),
				formatWithApostrophe(stats.NumberOfExplicitStatements),
				formatWithApostrophe(stats.NumberOfEntities))
			total.NumberOfStatements += stats.NumberOfStatements
			total.NumberOfExplicitStatements += stats.NumberOfExplicitStatements
			total.NumberOfEntities += stats.NumberOfEntities
		}
	}

	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("%-50s %12s %20s %15s\n", "TOTAL",
		formatWithApostrophe(total.NumberOfStatements),
		formatWithApostrophe(total.NumberOfExplicitStatements),
		formatWithApostrophe(total.NumberOfEntities))
}

// formatWithApostrophe formats an integer with apostrophes as thousands separators (e.g., 12'345'678)
func formatWithApostrophe(n int) string {
	s := strconv.Itoa(n)
	nLen := len(s)
	if nLen <= 3 {
		return s
	}
	var out []byte
	count := 0
	for i := nLen - 1; i >= 0; i-- {
		out = append([]byte{s[i]}, out...)
		count++
		if count%3 == 0 && i != 0 {
			out = append([]byte{'\''}, out...)
		}
	}
	return string(out)
}
