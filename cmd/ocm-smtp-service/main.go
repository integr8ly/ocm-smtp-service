package main

import (
	"flag"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/cmd/ocm-smtp-service/migrate"
	"gitlab.cee.redhat.com/service/ocm-smtp-service/cmd/ocm-smtp-service/servecmd"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

//nolint
//go:generate go-bindata -o ../../data/generated/openapi/openapi.go -pkg openapi -prefix ../../openapi/ ../../openapi
func init() {
}

func main() {
	// This is needed to make `glog` believe that the flags have already been parsed, otherwise
	// every log messages is prefixed by an error message stating the the flags haven't been
	// parsed.
	_ = flag.CommandLine.Parse([]string{})

	//pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	// Always log to stderr by default
	if err := flag.Set("logtostderr", "true"); err != nil {
		glog.Infof("Unable to set logtostderr to true")
	}

	rootCmd := &cobra.Command{
		Use:  "ocm-smtp-service",
		Long: "ocm-smtp-service serves as an example service template for new microservices",
	}

	// All subcommands under root
	migrateCmd := migrate.NewMigrateCommand()
	serveCmd := servecmd.NewServeCommand()

	// Add subcommand(s)
	rootCmd.AddCommand(migrateCmd, serveCmd)

	if err := rootCmd.Execute(); err != nil {
		glog.Fatalf("error running command: %v", err)
	}
}
