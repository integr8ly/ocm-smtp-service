# Go API client for openapi

Manages SMTP credentials

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.openshift.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**ApiOcmSmtpServiceV1SmtpDelete**](docs/DefaultApi.md#apiocmsmtpservicev1smtpdelete) | **Delete** /api/ocm-smtp-service/v1/smtp | Delete an SMTP credential for a cluster
*DefaultApi* | [**ApiOcmSmtpServiceV1SmtpGet**](docs/DefaultApi.md#apiocmsmtpservicev1smtpget) | **Get** /api/ocm-smtp-service/v1/smtp | Returns a list of SMTP
*DefaultApi* | [**ApiOcmSmtpServiceV1SmtpPost**](docs/DefaultApi.md#apiocmsmtpservicev1smtppost) | **Post** /api/ocm-smtp-service/v1/smtp | Create a new SMTP credential


## Documentation For Models

 - [ClusterMeta](docs/ClusterMeta.md)
 - [ClusterMetaAllOf](docs/ClusterMetaAllOf.md)
 - [Error](docs/Error.md)
 - [ErrorAllOf](docs/ErrorAllOf.md)
 - [ErrorList](docs/ErrorList.md)
 - [ErrorListAllOf](docs/ErrorListAllOf.md)
 - [List](docs/List.md)
 - [ObjectReference](docs/ObjectReference.md)
 - [Smtp](docs/Smtp.md)
 - [SmtpAllOf](docs/SmtpAllOf.md)
 - [SmtpDeleteRequest](docs/SmtpDeleteRequest.md)
 - [SmtpList](docs/SmtpList.md)
 - [SmtpListAllOf](docs/SmtpListAllOf.md)


## Documentation For Authorization



## Bearer

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```



## Author



