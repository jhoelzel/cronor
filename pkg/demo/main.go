//demo defines all the different objects types for this deployment
package demo

import (
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
)

type Demodeployment struct {
	Deployment   *appsv1.Deployment
	Name         string
	Namespace    string
	Service      *apiv1.Service
	NamespaceObj *apiv1.Namespace
}

//NewDemoDeployment creates a new demo deployment for us
func NewDemoDeployment(name, namespace string) (depl *Demodeployment) {
	depl = &Demodeployment{}
	depl.Name = name
	depl.Namespace = namespace
	depl.getDeploymentDef()
	depl.getLoadbalancerServiceDef()
	depl.getNamespaceDef()
	return
}

//SetDemodeployment sets an existing deploymen as the current deployment NOTE: We dont update the services or namespaces diretly its not needed for our scenario
func SetDemoDeployment(name, namespace string, deployment *appsv1.Deployment) (depl *Demodeployment) {
	depl = &Demodeployment{}
	depl.Name = name
	depl.Namespace = namespace
	depl.Deployment = deployment
	depl.getLoadbalancerServiceDef()
	depl.getNamespaceDef()
	return
}
func int32Ptr(i int32) *int32 { return &i }
