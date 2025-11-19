package repository

import (
	"encoding/json"
)

type Params struct {
	ReadOnly                               bool     `json:"readOnly"`
	Ruleset                                string   `json:"ruleset"`
	DisableSameAs                          bool     `json:"disableSameAs"`
	CheckForInconsistencies                bool     `json:"checkForInconsistencies"`
	IsShacl                                bool     `json:"isShacl"`
	CacheSelectNodes                       bool     `json:"cacheSelectNodes"`
	LogValidationPlans                     bool     `json:"logValidationPlans"`
	ParallelValidation                     bool     `json:"parallelValidation"`
	PerformanceLogging                     bool     `json:"performanceLogging"`
	DashDataShapes                         bool     `json:"dashDataShapes"`
	LogValidationViolations                bool     `json:"logValidationViolations"`
	GlobalLogValidationExecution           bool     `json:"globalLogValidationExecution"`
	EclipseRdf4jShaclExtensions            bool     `json:"eclipseRdf4jShaclExtensions"`
	ValidationResultsLimitTotal            uint32   `json:"validationResultsLimitTotal"`
	ValidationResultsLimitPerConstraint    uint32   `json:"validationResultsLimitPerConstraint"`
	ShapesGraph                            string   `json:"shapesGraph"`
	EntityIdSize                           uint8    `json:"entityIdSize"`
	EnableContextIndex                     bool     `json:"enableContextIndex"`
	EnablePredicateList                    bool     `json:"enablePredicateList"`
	EnableFtsIndex                         bool     `json:"enableFtsIndex"`
	EntityIndexSize                        uint32   `json:"entityIndexSize"`
	FtsIndexes                             []string `json:"ftsIndexes"`
	FtsStringLiteralsIndex                 string   `json:"ftsStringLiteralsIndex"`
	FtsIrisIndex                           string   `json:"ftsIrisIndex"`
	ThrowQueryEvaluationExceptionOnTimeout bool     `json:"throwQueryEvaluationExceptionOnTimeout"`
	QueryTimeout                           uint32   `json:"queryTimeout"`
	QueryLimitResults                      uint32   `json:"queryLimitResults"`
}

type Config struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Params   Params `json:"params"`
	Location string `json:"location"`
}

// ToJSON returns the JSON representation of the Config struct as a string
func (c *Config) ToJSON() (string, error) {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
