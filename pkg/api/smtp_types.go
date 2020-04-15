package api

import "github.com/jinzhu/gorm"

type SMTP struct {
	Meta
	SendGridID string
	Host string
	Port string
	TLS string
	Username string
	Password string
}

// SMTPClient to create SMTP details for an OpenShift cluster by its ID
type SMTPClient interface {
	Create(id string) (*SMTP, error)
	Get(id string) (*SMTP, error)
	Refresh(id string) (*SMTP, error)
	Delete(id string) error
}

type SMTPList []*SMTP
type SMTPIndex map[string]*SMTP

func (l SMTPList) Index() SMTPIndex {
	index := SMTPIndex{}
	for _, o := range l {
		index[o.ID] = o
	}
	return index
}

func (org *SMTP) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", NewID())
}
