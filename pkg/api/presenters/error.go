package presenters

import (
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/api/openapi"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/errors"
)

func PresentError(err *errors.ServiceError) openapi.Error {
	return err.AsOpenapiError("")
}
