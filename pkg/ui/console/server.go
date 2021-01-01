package console

import (
	"net/http"
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/go-logr/logr"
	"github.com/infinispan/infinispan-operator/pkg/ui/console/rest"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var consoleRoutes = []string{
	"/support",
	"/settings/general",
	"/settings/profile",
}

func StartServer(client crclient.Client, watchNamespace string, log logr.Logger) error {
	cl := rest.ClientHandler{
		Client:          client,
		WatchNamespaces: strings.Split(watchNamespace, ","),
		Log:             log,
	}

	appBox := rice.MustFindBox("../../../ui/console/dist")
	assetHandler := http.FileServer(appBox.HTTPBox())
	index, err := appBox.Open("index.html")
	if err != nil {
		log.Error(err, "Unable to find index.html file")
	}

	http.Handle("/", assetHandler)
	for _, consoleRoute := range consoleRoutes {
		http.Handle(consoleRoute, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			stat, _ := index.Stat()
			http.ServeContent(w, r, consoleRoute, stat.ModTime(), index)
		}))
	}

	http.Handle("/rest/clusters", http.HandlerFunc(cl.GetClusters))

	return http.ListenAndServe(":8080", nil)
}
