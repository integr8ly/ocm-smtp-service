package services

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/api"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/db"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/errors"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/logger"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/sendgrid"
	"os"
)

type SMTPService interface {
	Get(ctx context.Context, id string) (*api.SMTP, *errors.ServiceError)
	List(ctx context.Context, listArgs *ListArguments) (api.SMTPList, *api.PagingMeta, *errors.ServiceError)
	Create(ctx context.Context, cluster *api.ClusterMeta) (*api.SMTP, *errors.ServiceError)
	Delete(ctx context.Context, smtp *api.SMTP) *errors.ServiceError
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

var log = logrus.NewEntry(&logrus.Logger{
	Out:          os.Stderr,
	Formatter:    &logrus.TextFormatter{},
	ReportCaller: false,
	Level:        logrus.FatalLevel,
})

func (s sqlSMTPService) Get(ctx context.Context, id string) (*api.SMTP, *errors.ServiceError) {
	gorm := s.connectionFactory.New()
	var smtp api.SMTP
	if err := gorm.First(&smtp, "cluster_id =?", id).Error; err != nil {
		return nil, handleGetError("SMTP", "id", id, err)
	}
	return &smtp, nil
}

func (s sqlSMTPService) List(ctx context.Context, listArgs *ListArguments) (api.SMTPList, *api.PagingMeta, *errors.ServiceError) {
	gorm := s.connectionFactory.New()
	olog := logger.NewOCMLogger(ctx)
	pagingMeta := api.PagingMeta{
		Page: listArgs.Page,
	}

	// Unbounded list operations should be discouraged, as they can result in very long API operations
	if listArgs.Size < 0 {
		olog.Warningf("A query with an unbound size was requested.")
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

func (s sqlSMTPService) Create(ctx context.Context, cluster *api.ClusterMeta) (*api.SMTP, *errors.ServiceError) {
	sendgridClient, err := setupSMTPDetailsClient(log)
	if err != nil {
		return nil, errors.GeneralError(err.Error(), "failed to create sendgrid client")
	}

	smtpDetails, err := sendgridClient.Create(cluster.ClusterID)
	if err != nil {
		if api.IsAlreadyExistsError(err) {
			return nil, errors.GeneralError(err.Error(), fmt.Sprintf("api key for cluster %s already exists", cluster.ClusterID))
		}
		return nil, errors.GeneralError(err.Error(), "error returned from sendgrid")
	}

	gorm := s.connectionFactory.New()

	if err := gorm.Save(smtpDetails).Error; err != nil {
		db.MarkForRollback(ctx, err)
		return nil, handleUpdateError("SMTP", err)
	}
	return smtpDetails, nil
}

func (s sqlSMTPService) Delete(ctx context.Context, smtp *api.SMTP) *errors.ServiceError {
	sendgridClient, err := setupSMTPDetailsClient(log)
	if err != nil {
		return errors.GeneralError(err.Error(), "failed to create sendgrid client")
	}

	if err := sendgridClient.Delete(smtp.ClusterID); err != nil {
		if api.IsNotExistError(err) {
			return errors.GeneralError(err.Error(), fmt.Sprintf("api key for cluster %s does not exist", smtp.ClusterID))
		}
		return errors.GeneralError(err.Error(), "failed to delete api key")
	}

	gorm := s.connectionFactory.New()
	if err := gorm.Delete(smtp).Error; err != nil {
		db.MarkForRollback(ctx, err)
		return handleDeleteError("SMTP", err)
	}

	return nil
}

func setupSMTPDetailsClient(logger *logrus.Entry) (*sendgrid.Client, error) {
	smtpdetailsClient, err := sendgrid.NewDefaultClient(logger)
	if err != nil {
		logger.Fatalf("failed to create sendgrid details client: %v", err)
		return nil, errors.GeneralError(err.Error(), "failed to setup sendgrid smtp details client")
	}
	return smtpdetailsClient, nil
}
