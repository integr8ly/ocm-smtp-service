package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
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
				Kind:  "SMTPList",
				Page:  int32(paging.Page),
				Size:  int32(paging.Size),
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
	var clusterMeta openapi.ClusterMeta
	cfg := &handlerConfig{
		// cluster id
		MarshalInto: &clusterMeta,
		// validate the body has the cluster id
		Validate: []validate{
			validateEmpty(&clusterMeta.Id, "id"),
			validateNotEmpty(&clusterMeta.ClusterID, "clusterID"),
		},
		Action: func() (interface{}, *errors.ServiceError) {
			ctx := r.Context()
			convertedMeta := presenters.ConvertClusterMeta(clusterMeta)
			created, err := s.service.Create(ctx, convertedMeta)
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
	cfg := &handlerConfig{
		// todo add validation
		Action: func() (interface{}, *errors.ServiceError) {
			ctx := r.Context()
			id := mux.Vars(r)["id"]
			found, err := s.service.Get(ctx, id)
			if err != nil {
				return nil, err
			}
			if err := s.service.Delete(ctx, found); err != nil {
				return nil, err
			}
			return fmt.Sprintf("%s deleted", id), nil
		},
		ErrorHandler: handleError,
	}
	handleDelete(w, r, cfg, http.StatusAccepted)
}
