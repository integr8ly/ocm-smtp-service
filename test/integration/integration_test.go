package integration

import (
	"flag"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/test"
	"os"
	"runtime"
	"testing"

	"github.com/golang/glog"
)

func TestMain(m *testing.M) {
	flag.Parse()
	glog.Infof("Starting integration test using go version %s", runtime.Version())
	helper := test.NewHelper(&testing.T{})
	helper.ResetDB()
	exitCode := m.Run()
	helper.Teardown()
	os.Exit(exitCode)
}
