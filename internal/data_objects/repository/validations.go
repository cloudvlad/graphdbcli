package repository

import (
	"fmt"
	"strings"
)

// Supported rulesets for validation
var SupportedRulesets = []string{"rdfsplus-optimized"}

// ValidateRuleset checks if the ruleset is one of the supported default rulesets
func (p Params) ValidateRuleset() error {
	for _, s := range SupportedRulesets {
		if p.Ruleset == s {
			return nil
		}
	}
	return fmt.Errorf("unsupported ruleset: %s", p.Ruleset)
}

// ValidateShapesGraph checks if the shapesGraph is a valid URI (basic check)
func (p Params) ValidateShapesGraph() error {
	if !strings.HasPrefix(p.ShapesGraph, "http://") && !strings.HasPrefix(p.ShapesGraph, "https://") {
		return fmt.Errorf("shapesGraph must be a valid URI")
	}
	return nil
}

// ValidateEntityIndexSize checks for a reasonable range
func (p Params) ValidateEntityIndexSize() error {
	if p.EntityIndexSize < 1000 || p.EntityIndexSize > 100000000 {
		return fmt.Errorf("entityIndexSize out of range")
	}
	return nil
}

// ValidateFtsIndexes checks that at least one index is provided
func (p Params) ValidateFtsIndexes() error {
	if len(p.FtsIndexes) == 0 {
		return fmt.Errorf("ftsIndexes must have at least one value")
	}
	return nil
}

// ValidateFtsStringLiteralsIndex checks for non-empty value
func (p Params) ValidateFtsStringLiteralsIndex() error {
	if p.FtsStringLiteralsIndex == "" {
		return fmt.Errorf("ftsStringLiteralsIndex cannot be empty")
	}
	return nil
}

// ValidateFtsIrisIndex checks for non-empty value
func (p Params) ValidateFtsIrisIndex() error {
	if p.FtsIrisIndex == "" {
		return fmt.Errorf("ftsIrisIndex cannot be empty")
	}
	return nil
}

// ValidateValidationResultsLimitTotal checks for reasonable range
func (p Params) ValidateValidationResultsLimitTotal() error {
	if p.ValidationResultsLimitTotal < 1 || p.ValidationResultsLimitTotal > 10000000 {
		return fmt.Errorf("validationResultsLimitTotal out of range")
	}
	return nil
}

// ValidateValidationResultsLimitPerConstraint checks for reasonable range
func (p Params) ValidateValidationResultsLimitPerConstraint() error {
	low := uint32(1)
	high := uint32(1000000)
	if p.ValidationResultsLimitPerConstraint < low || p.ValidationResultsLimitPerConstraint > high {
		println(p.ValidationResultsLimitPerConstraint)
		return fmt.Errorf("validationResultsLimitPerConstraint out of range %v - %v", low, high)
	}
	return nil
}
