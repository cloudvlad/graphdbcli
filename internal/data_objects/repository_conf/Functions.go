package repository_conf

// NewDefaultRepositoryConfig returns a RepositoryConfig populated with the
// default values from the provided RDF/Turtle snippet.
func NewDefaultRepositoryConfig(id string) *RepositoryConfig {
	return &RepositoryConfig{
		RepositoryID: id,
		RepositoryImpl: RepositoryImpl{
			RepositoryType: "graphdb:SailRepository",
			SailImpl: SailImpl{
				BaseURL:                                "http://example.org/owlim#",
				CheckForInconsistencies:                false,
				DefaultNS:                              "",
				DisableSameAs:                          true,
				EnableContextIndex:                     false,
				EnableFtsIndex:                         false,
				EnableLiteralIndex:                     true,
				EnablePredicateList:                    true,
				EntityIDSize:                           32,
				EntityIndexSize:                        10000000,
				FtsIndexes:                             []string{"default", "iri"},
				FtsIrisIndex:                           "none",
				FtsStringLiteralsIndex:                 "default",
				Imports:                                []string{""},
				InMemoryLiteralProperties:              true,
				QueryLimitResults:                      0,
				QueryTimeout:                           0,
				ReadOnly:                               false,
				RepositoryType:                         "file-repository",
				Ruleset:                                "rdfsplus-optimized",
				StorageFolder:                          "storage",
				ThrowQueryEvaluationExceptionOnTimeout: false,
				SailType:                               "graphdb:Sail",
			},
		},
		Label: "",
	}
}
