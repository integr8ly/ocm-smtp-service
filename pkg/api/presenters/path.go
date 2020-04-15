package presenters

import (
	"fmt"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/api"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/errors"
)

const (
	BasePath = "/api/ocm-smtp-service/v1"
)

func ObjectPath(id string, obj interface{}) string {
	return fmt.Sprintf("%s/%s/%s", BasePath, path(obj), id)
}

func path(i interface{}) string {
	switch i.(type) {
	case api.SMTP, *api.SMTP:
		return "smtp"
	case errors.ServiceError, *errors.ServiceError:
		return "errors"
	default:
		return ""
	}
}
