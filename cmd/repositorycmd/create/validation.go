package create

import (
	"graphdbcli/internal/data_objects/repository"
)

func validation(config repository.Config) error {
	var err error

	err = config.Params.ValidateRuleset()
	if err != nil {
		return err
	}

	err = config.Params.ValidateShapesGraph()
	if err != nil {
		return err
	}

	err = config.Params.ValidateEntityIndexSize()
	if err != nil {
		return err
	}

	err = config.Params.ValidateFtsIndexes()
	if err != nil {
		return err
	}

	err = config.Params.ValidateFtsStringLiteralsIndex()
	if err != nil {
		return err
	}

	err = config.Params.ValidateFtsIrisIndex()
	if err != nil {
		return err
	}

	err = config.Params.ValidateFtsIrisIndex()
	if err != nil {
		return err
	}

	err = config.Params.ValidateValidationResultsLimitTotal()
	if err != nil {
		return err
	}

	err = config.Params.ValidateValidationResultsLimitPerConstraint()
	if err != nil {
		return err
	}

	return nil
}
