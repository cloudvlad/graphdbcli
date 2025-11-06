package rdf_bombarder

import (
	"fmt"
	"math/rand"
	"strings"
)

// randomInsertQuery generates a random SPARQL INSERT DATA query for a batch
func randomInsertQuery(graphIdx, offset, count int, interconnection float64) string {
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
