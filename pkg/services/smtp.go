package services

import (
	"context"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/api"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/db"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/errors"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/logger"
)

type SMTPService interface {
	List(ctx context.Context, listArgs *ListArguments) (api.SMTPList, *api.PagingMeta, *errors.ServiceError)
	Create(ctx context.Context, smtp *api.SMTP) (*api.SMTP, *errors.ServiceError)
}

func NewSMTPService(connectionFactory *db.ConnectionFactory) SMTPService {
	return &sqlSMTPService{
		connectionFactory: connectionFactory,
	}
}

var _ SMTPService = &sqlSMTPService{}

type sqlSMTPService struct {
	connectionFactory *db.ConnectionFactory
}

func (s sqlSMTPService) List(ctx context.Context, listArgs *ListArguments) (api.SMTPList, *api.PagingMeta, *errors.ServiceError) {
	gorm := s.connectionFactory.New()
	ulog := logger.NewUHCLogger(ctx)
	pagingMeta := api.PagingMeta{
		Page:  listArgs.Page,
	}

	// Unbounded list operations should be discouraged, as they can result in very long API operations
	if listArgs.Size < 0 {
		ulog.Warningf("A query with an unbound size was requested.")
	}

	// Get the total number of records
	gorm.Model(api.SMTP{}).Count(&pagingMeta.Total)

	// Set the order by arguments
	for _, orderByArg := range listArgs.OrderBy {
		gorm = gorm.Order(orderByArg)
	}

	// TODO Search

	// Get the full list, using page/size to limit the result set
	var smtp api.SMTPList
	if err := gorm.Offset((listArgs.Page - 1) * listArgs.Size).Limit(listArgs.Size).Find(&smtp).Error; err != nil {
		return smtp, &pagingMeta, errors.GeneralError("Unable to list smtp details: %s", err)
	}

	// Set the proper size, as the result set total may be less than the requested size
	pagingMeta.Size = len(smtp)

	return smtp, &pagingMeta, nil
}

func (s sqlSMTPService) Create(ctx context.Context, smtp *api.SMTP) (*api.SMTP, *errors.ServiceError) {
	gorm := s.connectionFactory.New()

	if err := gorm.Save(smtp).Error; err != nil {
		db.MarkForRollback(ctx, err)
		return nil, handleUpdateError("SMTP", err)
	}
	return smtp, nil
}




