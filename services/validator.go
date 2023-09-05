package services

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func validateSlugName(name string) bool {
	pattern, err := regexp.Compile(`^[A-Z0-9]+(?:_[A-Z0-9]+)*$`)
	if err != nil {
		return false
	}

	return pattern.MatchString((name))
}

func validateNumeric(numb string) bool {
	pattern, err := regexp.Compile(`^[0-9]+$`)
	if err != nil {
		return false
	}

	return pattern.MatchString(numb)
}

func validateSlugsListToAdd(field validator.FieldLevel) bool {
	inter := field.Field()

	slice, ok := inter.Interface().([]map[string]string)
	if !ok {
		return false
	}

	for _, val := range slice {
		switch keys := len(val); {
		case (keys == 0) || (keys > 2):
			return false
		case keys == 1:
			if _, exists := val["slug_name"]; !exists {
				return false
			}
		case keys == 2:
			if _, slug := val["slug_name"]; !slug {
				return false
			} else if res := validateSlugName(val["slug_name"]); !res {
				return false
			}

			if _, days := val["days"]; !days {
				return false
			} else if res := validateNumeric(val["days"]); !res {
				return false
			}
		default:
			return false
		}
	}

	return true
}

func validateSlugsListToDel(field validator.FieldLevel) bool {
	inter := field.Field()

	slice, ok := inter.Interface().([]string)
	if !ok {
		return false
	}

	for _, val := range slice {
		if res := validateSlugName(val); !res {
			return false
		}
	}

	return true
}
