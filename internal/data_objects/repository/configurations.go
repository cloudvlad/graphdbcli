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
	// Never used, but required
	EnableLiteralIndex           bool   `json:"enableLiteralIndex"`
	RdfsSubClassReasoning        bool   `json:"rdfsSubClassReasoning"`
	ValidationEnabled            bool   `json:"validationEnabled"`
	StorageFolder                string `json:"storageFolder"`
	TransactionalValidationLimit string `json:"transactionalValidationLimit"`
	Imports                      string `json:"imports"`
	InMemoryLiteralProperties    bool   `json:"inMemoryLiteralProperties"`
	DefaultNS                    string `json:"defaultNS"`
	BaseURL                      string `json:"baseURL"`
	RepositoryType               string `json:"repositoryType"`
}

type Config struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Params   Params `json:"params"`
	Type     string `json:"type"`
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

// SetDefaults sets the default values for fields that are not typically configured.
func (c *Config) SetDefaults() {
	c.Type = "graphdb"
	p := c.Params
	p.EnablePredicateList = true
	p.RdfsSubClassReasoning = true
	p.ValidationEnabled = true
	p.StorageFolder = "storage"
	p.TransactionalValidationLimit = "500000"
	p.Imports = ""
	p.InMemoryLiteralProperties = true
	p.DefaultNS = ""
	p.BaseURL = "http://example.org/owlim#"
	p.RepositoryType = "file-repository"
}
