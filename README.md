OCM SMTP Service 
---

This project is an SMTP credential management service for RHMI

## *** NOTICE ***
This Github repository has been archived and migrated to [Gitlab](https://gitlab.cee.redhat.com/service/ocm-sendgrid-service)

## SMTP Management OpenAPI 
[SMTP Management OpenAPI Documentation](./pkg/api/openapi/README.md)

## Running the Service Locally
### Setup Postgres
An instance of Postgres is required to run this service locally, the following steps will install and setup a postgres locally for you.
```
# install and enable postgresql
sudo dnf install postgresql
sudo dnf install postgresql-server
sudo -i -u postgres
initdb --auth-local=trust
logout
sudo systemctl enable --now postgresql
systemctl status postgresql
# add your username to posgresql
sudo -i -u postgres
createuser --interactive --pwprompt
  Enter name of role to add: <your local username>
  Enter password for new role: <******>
  Enter it again: <******>
  Shall the new role be a superuser? (y/n) y
logout
createdb --owner=<your local username> <your local username>
sudo systemctl restart postgresql
# one-time setup for ocm-smtp-service
cd $GOPATH/src/gitlab.cee.redhat.com/service/ocm-smtp-service/
./local_db_setup.sh
```
### Run the Service
```
$ make run
```

### Interacting with the API
The service uses the same authentication and API infrastructure as other OCM services, so to test it you can use the OCM command line tool. First you will need to log-in:

Browse to [cloud.redhat.com/beta/openshift/token](https://cloud.redhat.com/beta/openshift/token) and copy the offline access token
```
$ ocm login --url=http://localhost:8000 --token=<<your offline token>> --insecure
```
To list SMTP credentials
```
$ ocm get /api/ocm-smtp-service/v1/smtp 
```
## Generating the OpenAPI
After updating `openapi/ocm-smtp-service.yaml` run the following to generate the OpenAPI
```
$ make generate
```
Check `pkg/api/openapi` directory for generated files. 

**NOTE** If there is a `go.mod` file in this directory, remove it before running the service. 

## Openshift Deployment
[Documentation](./templates/README.md) for deploying a service to Openshift