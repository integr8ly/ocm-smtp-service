package handlers

import (
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/errors"
	"strings"
)

func validateNotEmpty(value *string, field string) validate {
	return func() *errors.ServiceError {
		if value == nil || len(*value) == 0 {
			return errors.Validation("%s is required", field)
		}
		return nil
	}
}

func validateIntNotEmpty(value *int, field string) validate {
	return func() *errors.ServiceError {
		if value == nil {
			return errors.Validation("%s is required", field)
		}
		return nil
	}
}

func validateEmpty(value *string, field string) validate {
	return func() *errors.ServiceError {
		if len(*value) > 0 {
			return errors.Validation("%s must be empty", field)
		}
		return nil
	}
}

func validateInclusionIn(value *string, list []string) validate {
	return func() *errors.ServiceError {
		for _, item := range list {
			if strings.EqualFold(*value, item) {
				return nil
			}
		}
		return errors.Validation("%s is not a valid value", *value)
	}
}

func validateNonNegative(value *int32, field string) validate {
	return func() *errors.ServiceError {
		if (*value) < 0 {
			return errors.Validation("%s must be non-negative", field)
		}
		return nil
	}
}
