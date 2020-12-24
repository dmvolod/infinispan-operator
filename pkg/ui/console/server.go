package console

import (
	"net/http"
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/go-logr/logr"
	"github.com/infinispan/infinispan-operator/pkg/ui/console/rest"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func StartServer(client crclient.Client, watchNamespace string, log logr.Logger) error {
	cl := rest.ClientHandler{
		Client:          client,
		WatchNamespaces: strings.Split(watchNamespace, ","),
		Log:             log,
	}

	appBox := rice.MustFindBox("../../../ui/console/dist")
	http.Handle("/", http.FileServer(appBox.HTTPBox()))
	http.Handle("/rest/clusters", http.HandlerFunc(cl.GetClusters))

	return http.ListenAndServe(":8080", nil)
}
