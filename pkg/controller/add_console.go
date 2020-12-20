package controller

import "github.com/infinispan/infinispan-operator/pkg/controller/ui"

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, ui.Add)
}
