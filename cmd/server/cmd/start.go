package cmd

import (
	"fmt"
	"github.com/Dimss/first/apis/metafirst/v1alpha1"
	"github.com/Dimss/first/pkg/apiserver"
	metafirstopenapi "github.com/Dimss/first/pkg/generated/openapi"
	"github.com/spf13/cobra"
	"io"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apiserver/pkg/endpoints/openapi"
	"k8s.io/apiserver/pkg/features"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/informers"
	netutils "k8s.io/utils/net"
	"net"
	"os"
)

const defaultEtcdPathPrefix = "/registry/metafirst.cnvrg.io"

type FirstServerOptions struct {
	RecommendedOptions    *genericoptions.RecommendedOptions
	SharedInformerFactory informers.SharedInformerFactory
	StdOut                io.Writer
	StdErr                io.Writer

	AlternateDNS []string
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func (o *FirstServerOptions) Validate(args []string) error {
	var errors []error
	errors = append(errors, o.RecommendedOptions.Validate()...)
	return utilerrors.NewAggregate(errors)
}

// Complete fills in fields required to have valid data
func (o *FirstServerOptions) Complete() error {
	o.RecommendedOptions.Etcd.StorageConfig.Paging = true

	o.RecommendedOptions.Etcd.StorageConfig.Transport.ServerList = []string{"127.0.0.1"}
	return nil
}

func (o *FirstServerOptions) Config() (*apiserver.Config, error) {

	if err := o.
		RecommendedOptions.
		SecureServing.
		MaybeDefaultWithSelfSignedCerts(
			"localhost",
			o.AlternateDNS,
			[]net.IP{netutils.ParseIPSloppy("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	serverConfig := genericapiserver.NewRecommendedConfig(apiserver.Codecs)

	serverConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(metafirstopenapi.GetOpenAPIDefinitions, openapi.NewDefinitionNamer(apiserver.Scheme))
	serverConfig.OpenAPIConfig.Info.Title = "First"
	serverConfig.OpenAPIConfig.Info.Version = "0.1"

	if utilfeature.DefaultFeatureGate.Enabled(features.OpenAPIV3) {
		serverConfig.OpenAPIV3Config = genericapiserver.DefaultOpenAPIV3Config(metafirstopenapi.GetOpenAPIDefinitions, openapi.NewDefinitionNamer(apiserver.Scheme))
		serverConfig.OpenAPIV3Config.Info.Title = "First"
		serverConfig.OpenAPIV3Config.Info.Version = "0.1"
	}

	if err := o.RecommendedOptions.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	config := &apiserver.Config{
		GenericConfig: serverConfig,
		ExtraConfig:   apiserver.ExtraConfig{},
	}
	return config, nil
}

func (o *FirstServerOptions) RunFirstServer(stopCh <-chan struct{}) error {
	config, err := o.Config()
	if err != nil {
		return err
	}

	server, err := config.Complete().New()
	if err != nil {
		return err
	}

	server.GenericAPIServer.AddPostStartHookOrDie("start-first-server-informers", func(context genericapiserver.PostStartHookContext) error {
		config.GenericConfig.SharedInformerFactory.Start(context.StopCh)
		o.SharedInformerFactory.Start(context.StopCh)
		return nil
	})

	return server.GenericAPIServer.PrepareRun().Run(stopCh)
}

func NewFirstServerOptions(out, errOut io.Writer) *FirstServerOptions {
	o := &FirstServerOptions{
		RecommendedOptions: genericoptions.NewRecommendedOptions(
			defaultEtcdPathPrefix,
			apiserver.Codecs.LegacyCodec(v1alpha1.SchemeGroupVersion),
		),

		StdOut: out,
		StdErr: errOut,
	}

	o.RecommendedOptions.Etcd.StorageConfig.EncodeVersioner = runtime.NewMultiGroupVersioner(v1alpha1.SchemeGroupVersion, schema.GroupKind{Group: v1alpha1.GroupName})

	return o
}

var startCmd = &cobra.Command{
	Use:   "start First K8s AA Server",
	Short: "start the aa server",
	RunE: func(cmd *cobra.Command, args []string) error {
		stopCh := genericapiserver.SetupSignalHandler()
		options := NewFirstServerOptions(os.Stdout, os.Stderr)

		flags := cmd.Flags()
		options.RecommendedOptions.AddFlags(flags)
		utilfeature.DefaultMutableFeatureGate.AddFlag(flags)

		if err := options.Complete(); err != nil {
			return err
		}
		if err := options.Validate(args); err != nil {
			return err
		}
		if err := options.RunFirstServer(stopCh); err != nil {
			return err
		}

		return nil
	},
}
