package create

import (
	repo "graphdbcli/internal/data_objects/repository"
	"reflect"
	"strings"
)

// ProcessConfigToExpectedPayload converts a Config object to an ExpectedPayload object, only setting fields that are non-zero/non-empty in the input.
func ProcessConfigToExpectedPayload(cfg repo.Config, defaultPayload ExpectedPayload) ExpectedPayload {
	params := make(map[string]Param)
	// Use reflection to check zero values
	cfgVal := reflect.ValueOf(cfg.Params)
	cfgType := reflect.TypeOf(cfg.Params)

	for k, def := range defaultPayload.Params {
		field, found := cfgType.FieldByNameFunc(func(fieldName string) bool {
			return strings.EqualFold(fieldName, k)
		})
		if found {
			val := cfgVal.FieldByName(field.Name)
			zero := reflect.Zero(val.Type()).Interface()
			if !reflect.DeepEqual(val.Interface(), zero) {
				params[k] = Param{
					Name:     def.Name,
					Label:    def.Label,
					Value:    val.Interface(),
					IsNumber: def.IsNumber,
				}
				continue
			}
		}

		params[k] = def
	}

	return ExpectedPayload{
		ID:       ifThenElse(cfg.ID != "", cfg.ID, defaultPayload.ID),
		Title:    ifThenElse(cfg.Title != "", cfg.Title, defaultPayload.Title),
		Type:     defaultPayload.Type,
		Location: ifThenElse(cfg.Location != "", cfg.Location, defaultPayload.Location),
		Params:   params,
	}
}

// Helper for inline ternary
func ifThenElse(cond bool, a, b string) string {
	if cond {
		return a
	}
	return b
}
