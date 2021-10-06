//demo defines all the different objects types for this deployment
package demo

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

//getLoadbalancerServiceDef returns the default loadbalancer definition
func (d *Demodeployment) getLoadbalancerServiceDef() {
	d.Service = &apiv1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      d.Name + "-loadblancer-service",
			Namespace: d.Namespace,
			Labels: map[string]string{
				"app":  "demo",
				"name": d.Name,
			},
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{{
				Name:       "demo-app-http",
				Port:       80,
				TargetPort: intstr.FromInt(8080),
			},
			},
			Selector: map[string]string{
				"app":  "demo",
				"name": d.Name,
			},
			Type: "LoadBalancer",
		},
	}
}
