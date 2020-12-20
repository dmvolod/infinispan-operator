package ui

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/operator-framework/operator-sdk/pkg/k8sutil"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const ControllerName = "console-controller"

type Controller struct {
	crclient.Client
	Log               logr.Logger
	Scheme            *runtime.Scheme
	OperatorName      string
	OperatorNamespace string
}

func Add(mgr manager.Manager) error {
	return CreateController(ControllerName, mgr)
}

func CreateController(name string, mgr manager.Manager) error {
	log := logf.Log.WithName(name)

	// Validates, that operator is running in non-local mode
	operatorName, operatorNamespace, err := GetOperator()
	if err != nil {
		log.Info("Controller not started. The operator seems to be starting in local mode", "error", err.Error())
		return nil
	}

	client, err := crclient.New(mgr.GetConfig(), crclient.Options{})
	if err != nil {
		log.Info("Controller not started. Unable to create client", "error", err.Error())
		return nil
	}
	if err := DeployConsoleUI(client, mgr.GetScheme(), operatorName, operatorNamespace, log); err != nil {
		return nil
	}

	r := &Controller{
		Client:            mgr.GetClient(),
		Log:               log,
		Scheme:            mgr.GetScheme(),
		OperatorName:      operatorName,
		OperatorNamespace: operatorNamespace,
	}

	// Create a new controller
	c, err := controller.New(name, mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	deploymentPredicate := predicate.Funcs{
		DeleteFunc: func(e event.DeleteEvent) bool {
			return false
		},
		CreateFunc: func(e event.CreateEvent) bool {
			return false
		},
		UpdateFunc: func(e event.UpdateEvent) bool {
			return false
		},
	}

	// Watch for changes to primary Infinispan resource
	err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForObject{}, deploymentPredicate)
	if err != nil {
		return err
	}

	secondaryTypePredicate := predicate.Funcs{
		CreateFunc: func(e event.CreateEvent) bool {
			return false
		},
		UpdateFunc: func(e event.UpdateEvent) bool {
			return false
		},
	}

	// Watch for changes to secondary configured resources
	for _, watchType := range []runtime.Object{&corev1.Pod{}, &corev1.Service{}} {
		if err = c.Watch(&source.Kind{Type: watchType}, &handler.EnqueueRequestForOwner{
			IsController: true,
			OwnerType:    &appsv1.Deployment{},
		}, secondaryTypePredicate); err != nil {
			return err
		}
	}

	return nil
}

func (r *Controller) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	if request.Name == r.OperatorName && request.Namespace == r.OperatorNamespace {
		return reconcile.Result{}, DeployConsoleUI(r.Client, r.Scheme, r.OperatorName, r.OperatorNamespace, r.Log)
	}
	return reconcile.Result{}, nil
}

func DeployConsoleUI(client crclient.Client, scheme *runtime.Scheme, operatorName, operatorNamespace string, logger logr.Logger) error {
	operatorDeployment := &appsv1.Deployment{}
	if err := client.Get(context.TODO(), types.NamespacedName{Namespace: operatorNamespace, Name: operatorName}, operatorDeployment); err != nil {
		logger.Info("Unable to get operator Deployment", "error", err.Error())
		return err
	}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ConsoleName(operatorName),
			Namespace: operatorNamespace,
			Labels:    map[string]string{"name" : ConsoleName(operatorName)},
		},
		Spec:       corev1.PodSpec{
			Containers: []corev1.Container{{
				Name:            "nginx-hello",
				Image:           "nginxdemos/nginx-hello",
				ImagePullPolicy: corev1.PullIfNotPresent,
			}},
		},
	}
	controllerutil.SetControllerReference(operatorDeployment, pod, scheme)
	if err := client.Create(context.Background(), pod); err != nil && !errors.IsAlreadyExists(err) {
		logger.Info("Unable to create operator console Pod", "error", err.Error())
		return err
	}

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ConsoleName(operatorName),
			Namespace: operatorNamespace,
			Labels:    map[string]string{"name" : ConsoleName(operatorName)},
		},
		Spec:       corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name: "console-port",
					Port: 8080,
				},
			},
			Selector: map[string]string{"name" : ConsoleName(operatorName)},
			Type:     corev1.ServiceTypeClusterIP,
		},
	}
	controllerutil.SetControllerReference(operatorDeployment, service, scheme)
	if err := client.Create(context.Background(), service); err != nil && !errors.IsAlreadyExists(err) {
		logger.Info("Unable to create operator console Service", "error", err.Error())
		return err
	}
	return nil
}

func ConsoleName(name string) string {
	return fmt.Sprintf("%v-console", name)
}

func GetOperator() (string, string, error) {
	operatorNamespace, err := k8sutil.GetOperatorNamespace()
	if err != nil {
		return "", "", err
	}
	operatorName, err := k8sutil.GetOperatorName()
	if err != nil {
		return "", "", err
	}
	return operatorName, operatorNamespace, nil
}
