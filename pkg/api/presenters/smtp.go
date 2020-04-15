package presenters

import (
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/api"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/api/openapi"
)

func ConvertSMTP(smtp openapi.Smtp) *api.SMTP {
	return &api.SMTP{
		Meta: api.Meta{
			ID:        smtp.Id,
		},
		SendGridID: smtp.SendGridID,
		Host: smtp.Host,
		Port: smtp.Port,
		TLS: smtp.Tls,
		Username: smtp.Username,
		Password: smtp.Password,
	}
}

func PresentSMTP(smtp *api.SMTP) openapi.Smtp {
	reference := PresentReference(smtp.ID, smtp)
	return openapi.Smtp{
		Id:        reference.Id,
		Kind:      reference.Kind,
		Href:      reference.Href,
		SendGridID: smtp.SendGridID,
		Host:      smtp.Host,
		Port:      smtp.Port,
		Tls:       smtp.TLS,
		Username:  smtp.Username,
		Password:  smtp.Password,
		CreatedAt: smtp.CreatedAt,
		UpdatedAt: smtp.UpdatedAt,
	}
}