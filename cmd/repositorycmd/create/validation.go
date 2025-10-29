package create

func validateBaseUrl(url string) error {
	return nil
}

// The number of namespaces must be equal to the number of schema files from the imports parameter
func validateDefaultNS(namespace []string) error {
	// check if the passed namespace is follwoing the standard for URIs
	// if not trhow error(the provided defaultns is not a alid URI or somethin similar
	// eg corrct uri http://www.w3.org/2002/07/owl#

	return nil
}

func validateEntityIdSize(size int) error {
	return nil
}

func validateEntityIndexSize(size int) error {
	return nil
}

func validateFtsIndexes(indexes []string) error {
	return nil
}

func validateFtsIrisIndex(index string) error {
	return nil
}

func validateFtsStringLiteralsIndex(index string) error {
	return nil
}

func validateImports(imports []string) error {
	return nil
}

func validateQueryLimitResults(limit int) error {
	return nil
}

func validateQueryTimeout(timeout int) error {
	return nil
}

func validateRepositoryType(repoType string) error {
	return nil
}

func validateRuleset(ruleset string) error {
	return nil
}

func validateStorageFolder(folder string) error {
	return nil
}
