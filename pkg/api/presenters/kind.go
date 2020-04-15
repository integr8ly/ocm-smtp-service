package presenters

import (
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/api"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/errors"
)

func ObjectKind(i interface{}) string {
	switch i.(type) {
	case api.SMTP, *api.SMTP:
		return "SMTP"
	case errors.ServiceError, *errors.ServiceError:
		return "Error"
	default:
		return ""
	}
}
