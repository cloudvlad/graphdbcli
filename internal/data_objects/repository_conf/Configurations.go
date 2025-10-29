package repository_conf

// Structures representing repository configuration similar to the RDF/Turtle
// snippet provided by the user. Fields are exported and include `json` tags
// for convenient marshaling.
//
// For more information:
// https://graphdb.ontotext.com/documentation/11.1/configuring-a-repository.html#configuration-parameters

// SailImpl contains the GraphDB/Sail specific configuration options.
type SailImpl struct {
	BaseURL                                string   `json:"baseURL"`
	CheckForInconsistencies                bool     `json:"checkForInconsistencies"`
	DefaultNS                              string   `json:"defaultNS"`
	DisableSameAs                          bool     `json:"disableSameAs"`
	EnableContextIndex                     bool     `json:"enableContextIndex"`
	EnableFtsIndex                         bool     `json:"enableFtsIndex"`
	EnableLiteralIndex                     bool     `json:"enableLiteralIndex"`
	EnablePredicateList                    bool     `json:"enablePredicateList"`
	EntityIDSize                           int      `json:"entityIdSize"`
	EntityIndexSize                        int      `json:"entityIndexSize"`
	FtsIndexes                             []string `json:"ftsIndexes"`
	FtsIrisIndex                           string   `json:"ftsIrisIndex"`
	FtsStringLiteralsIndex                 string   `json:"ftsStringLiteralsIndex"`
	Imports                                []string `json:"imports"`
	InMemoryLiteralProperties              bool     `json:"inMemoryLiteralProperties"`
	QueryLimitResults                      int      `json:"queryLimitResults"`
	QueryTimeout                           int      `json:"queryTimeout"`
	ReadOnly                               bool     `json:"readOnly"`
	RepositoryType                         string   `json:"repositoryType"`
	Ruleset                                string   `json:"ruleset"`
	StorageFolder                          string   `json:"storageFolder"`
	ThrowQueryEvaluationExceptionOnTimeout bool     `json:"throwQueryEvaluationExceptionOnTimeout"`
	SailType                               string   `json:"sailType"`
}

// RepositoryImpl models the inner repository implementation that references a Sail implementation.
type RepositoryImpl struct {
	RepositoryType string   `json:"repositoryType"`
	SailImpl       SailImpl `json:"sailImpl"`
}

// RepositoryConfig is the top-level structure
type RepositoryConfig struct {
	RepositoryID   string         `json:"repositoryID"`
	RepositoryImpl RepositoryImpl `json:"repositoryImpl"`
	Label          string         `json:"label"`
}
