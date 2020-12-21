package console

import (
	"net/http"
	"strings"

	"github.com/go-logr/logr"
	"github.com/gorilla/mux"
	"github.com/infinispan/infinispan-operator/pkg/ui/console/rest"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func StartServer(client crclient.Client, watchNamespace string, log logr.Logger) error {
	route := mux.NewRouter()
	cl := rest.ClientHandler{
		Client:          client,
		WatchNamespaces: strings.Split(watchNamespace, ","),
		Log:             log,
	}

	route.HandleFunc("/", cl.HomePage)
	route.HandleFunc("/rest/clusters", cl.GetClusters).Methods(http.MethodGet)

	return http.ListenAndServe(":8080", route)
}
