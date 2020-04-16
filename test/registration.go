package test

import (
	"gitlab.cee.redhat.com/service/ocm-smtp-service/pkg/api/openapi"
	"testing"

	gm "github.com/onsi/gomega"
)

// Register a test
// This should be run before every integration test
func RegisterIntegration(t *testing.T) (*Helper, *openapi.APIClient) {
	// Register the test with gomega
	gm.RegisterTestingT(t)
	// Create a new helper
	helper := NewHelper(t)
	// Reset the database to a seeded blank state
	helper.ResetDB()
	// Create an api client
	client := helper.NewApiClient()
	return helper, client
}
