package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	ispnv1 "github.com/infinispan/infinispan-operator/pkg/apis/infinispan/v1"
	consts "github.com/infinispan/infinispan-operator/pkg/controller/constants"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const StatusUnknown = "UNKNOWN"

type Cluster struct {
	ClusterName      string             `json:"name"`
	ClusterNamespace string             `json:"namespace"`
	ServiceType      ispnv1.ServiceType `json:"type"`
	Status           string             `json:"status"`
	Console          string             `json:"console"`
}

func (c ClientHandler) GetClusters(w http.ResponseWriter, r *http.Request) {
	var clusters = make([]Cluster, 0)
	for _, ns := range c.WatchNamespaces {
		infinspanList := &ispnv1.InfinispanList{}
		listOps := &client.ListOptions{Namespace: ns}
		err := c.List(context.TODO(), infinspanList, listOps)
		if err != nil {
			c.Log.Error(err, "Unable to list Infinispan CR's")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		for _, infinispan := range infinspanList.Items {
			clusters = append(clusters, Cluster{ClusterName: infinispan.GetName(), ClusterNamespace: ns, ServiceType: infinispan.Spec.Service.Type, Status: c.GetClusterStatus(&infinispan), Console: c.GetConsoleURL(&infinispan)})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clusters)
}

func (c ClientHandler) GetClusterStatus(ispn *ispnv1.Infinispan) string {
	if ispn.IsWellFormed() {
		response, err := http.Get(fmt.Sprintf("%s://%s.%s.svc.cluster.local/%s", ispn.GetEndpointScheme(), ispn.Name, ispn.Namespace, consts.ServerHTTPHealthStatusPath))
		if err != nil {
			c.Log.Error(err, "Unable to get cluster health status", "name", ispn.Name, "namespace", ispn.Namespace)
			return StatusUnknown
		}
		responseStatus, err := ioutil.ReadAll(response.Body)
		if err != nil {
			c.Log.Error(err, "Unable to parse cluster health status", "name", ispn.Name, "namespace", ispn.Namespace)
			return StatusUnknown
		}
		return string(responseStatus)
	}
	return StatusUnknown
}

func (c ClientHandler) GetConsoleURL(ispn *ispnv1.Infinispan) string {
	if ispn.IsExposed() {
		switch ispn.GetExposeType() {
		case ispnv1.ExposeTypeRoute:
			if ispn.Spec.Expose.Host != "" {
				return fmt.Sprintf("%s://%s", ispn.GetEndpointScheme(), ispn.Spec.Expose.Host)
			}
		case ispnv1.ExposeTypeLoadBalancer:
			if ispn.Spec.Expose.Host != "" {
				return fmt.Sprintf("%s://%s:%d", ispn.GetEndpointScheme(), ispn.Spec.Expose.Host, consts.InfinispanPort)
			}
		}
	}
	return ""
}
