//demo defines all the different objects types for this deployment
package demo

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//getNamespaceDef returns the default namespace definition
func (d *Demodeployment) getNamespaceDef() {
	d.NamespaceObj = &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: d.Namespace,
		},
	}
}
