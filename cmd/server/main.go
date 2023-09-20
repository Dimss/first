package main

import (
	"github.com/Dimss/first/cmd/server/cmd"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/component-base/cli"
	"os"
)

func main() {
	os.Exit(
		cli.Run(
			cmd.NewCommandStartFirstServer(
				cmd.NewFirstServerOptions(os.Stdout, os.Stderr),
				genericapiserver.SetupSignalHandler(),
			),
		),
	)
}
