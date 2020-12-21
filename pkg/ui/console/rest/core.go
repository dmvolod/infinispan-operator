package rest

import (
	"fmt"
	"net/http"

	"github.com/go-logr/logr"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
)

type ClientHandler struct {
	crclient.Client
	WatchNamespaces []string
	Log             logr.Logger
}

func (c ClientHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	c.Log.Info("Endpoint Hit: homePage")
}
