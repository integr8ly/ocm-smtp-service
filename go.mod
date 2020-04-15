module gitlab.cee.redhat.com/service/ocm-smtp-service

go 1.12

require (
	github.com/Masterminds/squirrel v1.1.0
	github.com/antihax/optional v1.0.0
	github.com/auth0/go-jwt-middleware v0.0.0-20190805220309-36081240882b
	github.com/bxcodec/faker/v3 v3.2.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/go-healthcheck v0.1.0
	github.com/getsentry/sentry-go v0.3.1
	github.com/ghodss/yaml v1.0.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/google/uuid v1.1.1
	github.com/gorilla/handlers v1.4.0
	github.com/gorilla/mux v1.7.3
	github.com/integr8ly/smtp-service v0.2.1
	github.com/jinzhu/gorm v1.9.8
	github.com/lib/pq v1.1.1
	github.com/mendsley/gojwk v0.0.0-20141217222730-4d5ec6e58103
	github.com/onsi/gomega v1.7.1
	github.com/openshift-online/ocm-sdk-go v0.1.59
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v0.9.3
	github.com/segmentio/ksuid v1.0.2
	github.com/sendgrid/rest v2.4.1+incompatible
	github.com/sendgrid/sendgrid-go v3.5.0+incompatible
	github.com/sethvargo/go-password v0.1.3
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/yaacov/tree-search-language v0.0.0-20190923184055-1c2dad2e354b
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	gopkg.in/gormigrate.v1 v1.6.0
	k8s.io/api v0.18.1
	k8s.io/apimachinery v0.18.1
)
