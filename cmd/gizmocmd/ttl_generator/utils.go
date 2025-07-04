package ttl_generator

import (
	"fmt"
	"os"
	"path/filepath"
)

var numberOfStatementsPerFile uint
var numberOfFiles uint
var numberOfTripletsPerEntity uint

// generateTTLFiles creates TTL files and writes them to separate files according to the properties.
func generateTTLFiles(outputDir string) error {
	if numberOfFiles == 0 || numberOfStatementsPerFile == 0 || numberOfTripletsPerEntity == 0 {
		return fmt.Errorf("All properties must be set and greater than zero")
	}
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		return err
	}

	for i := uint(0); i < numberOfFiles; i++ {
		fileName := filepath.Join(outputDir, fmt.Sprintf("file_%d.ttl", i+1))
		f, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer f.Close()
		for j := uint(0); j < numberOfStatementsPerFile; j++ {
			entityID := i*numberOfStatementsPerFile + j + 1
			triplets := generateEntity(entityID)
			for _, t := range triplets {
				_, err := f.WriteString(t + "\n")
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// generateEntity generates TTL triplets for a single entity.
func generateEntity(entityID uint) []string {
	triplets := make([]string, 0, numberOfTripletsPerEntity)
	for k := uint(0); k < numberOfTripletsPerEntity; k++ {
		triplet := fmt.Sprintf("<http://example.org/entity/%d> <http://example.org/property/%d> \"value%d\" .", entityID, k+1, k+1)
		triplets = append(triplets, triplet)
	}
	return triplets
}
