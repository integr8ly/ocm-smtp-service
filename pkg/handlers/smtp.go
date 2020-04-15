package handlers

import (
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/api/openapi"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/api/presenters"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/errors"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/services"
	"net/http"
)

var _ RestHandler = smtpHandler{}

type smtpHandler struct {
	service services.SMTPService
}

func NewSMTPHandler(service services.SMTPService) *smtpHandler {
	return &smtpHandler{
		service: service,
	}
}

func (s smtpHandler) List(w http.ResponseWriter, r *http.Request) {
	cfg := &handlerConfig{
		Action: func() (interface{}, *errors.ServiceError) {
			ctx := r.Context()

			listArgs := services.NewListArguments(r.URL.Query())
			smtps, paging, err := s.service.List(ctx, listArgs)
			if err != nil {
				return nil, err
			}
			smtpList := openapi.SmtpList{
				Kind: "SMTPList",
				Page: int32(paging.Page),
				Size: int32(paging.Size),
				Total: int32(paging.Total),
				Items: []openapi.Smtp{},
			}
			for _, smtp := range smtps {
				converted := presenters.PresentSMTP(smtp)
				smtpList.Items = append(smtpList.Items, converted)
			}
			return smtpList, nil
		},
	}

	handleList(w, r, cfg)
}

func (s smtpHandler) Get(w http.ResponseWriter, r *http.Request) {
	handleError(r.Context(), w, errors.NotImplemented("get"))
}

func (s smtpHandler) Create(w http.ResponseWriter, r *http.Request) {
	var smtp openapi.Smtp
	cfg := &handlerConfig{
		MarshalInto:  &smtp,
		Validate:     []validate{
			validateEmpty(&smtp.Id, "id"),
			validateNotEmpty(&smtp.Username, "username"),
			validateNotEmpty(&smtp.Password, "password"),
			validateNotEmpty(&smtp.Port, "port"),
			validateNotEmpty(&smtp.Tls, "tls"),
			validateNotEmpty(&smtp.Host, "host"),
		},
		Action: func() (interface{}, *errors.ServiceError) {
			ctx := r.Context()
			converted := presenters.ConvertSMTP(smtp)
			created, err := s.service.Create(ctx, converted)
			if err != nil {
				return nil, err
			}
			return presenters.PresentSMTP(created), nil
		},
		ErrorHandler: handleError,
	}
	handle(w, r, cfg, http.StatusCreated)
}

func (s smtpHandler) Patch(w http.ResponseWriter, r *http.Request) {
	handleError(r.Context(), w, errors.NotImplemented("patch"))
}

func (s smtpHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handleError(r.Context(), w, errors.NotImplemented("delete"))
}

