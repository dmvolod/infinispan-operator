package launcher

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/go-logr/logr"
	"github.com/infinispan/infinispan-operator/pkg/apis"
	"github.com/infinispan/infinispan-operator/pkg/ui/console"
	"github.com/infinispan/infinispan-operator/version"
	"github.com/operator-framework/operator-sdk/pkg/k8sutil"
	"github.com/operator-framework/operator-sdk/pkg/log/zap"
	sdkv "github.com/operator-framework/operator-sdk/version"
	"github.com/spf13/pflag"
	"k8s.io/client-go/kubernetes/scheme"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

func printVersion(log logr.Logger) {
	log.Info(fmt.Sprintf("Operator Console Version: %s", version.Version))
	log.Info(fmt.Sprintf("Go Version: %s", runtime.Version()))
	log.Info(fmt.Sprintf("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH))
	log.Info(fmt.Sprintf("Operator SDK Version: %v", sdkv.Version))
}

var log = logf.Log.WithName("cmd")

// Parameters represent operator launcher parameters
type Parameters struct {
	StopChannel chan struct{}
}

// Launch launches operator
func Launch(params Parameters) {
	pflag.CommandLine.AddFlagSet(zap.FlagSet())
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.Parse()

	logf.SetLogger(logf.ZapLogger(false))

	printVersion(log)

	watchNamespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	// Get a config to talk to the API Server
	cfg, err := config.GetConfig()
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	log.Info("Registering Components.")
	// Setup Scheme for all resources
	if err := apis.AddToScheme(scheme.Scheme); err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	client, err := crclient.New(cfg, crclient.Options{})
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	log.Info("Starting WebServer...")
	if err := console.StartServer(client, watchNamespace, log); err != nil {
		log.Error(err, "")
		os.Exit(1)
	}
}
