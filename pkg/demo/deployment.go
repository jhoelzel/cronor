//demo defines all the different objects types for this deployment
package demo

import (
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

//getDeploymentDef returns the default deployment definition
func (d *Demodeployment) getDeploymentDef() {
	d.Deployment = &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      d.Name,
			Namespace: d.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "demo",
					"name": d.Name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "demo",
						"name": d.Name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "microk8s:32000/simpleapp:latest",
							Env:   []apiv1.EnvVar{},
							Resources: apiv1.ResourceRequirements{
								Limits: apiv1.ResourceList{
									"cpu":    resource.MustParse("50m"),
									"memory": resource.MustParse("50M"),
								},
								Requests: apiv1.ResourceList{
									"cpu":    resource.MustParse("25m"),
									"memory": resource.MustParse("25M"),
								},
							},
							ReadinessProbe: &apiv1.Probe{
								Handler:             apiv1.Handler{HTTPGet: &apiv1.HTTPGetAction{Path: "/status/readyz", Port: intstr.FromInt(8080), Scheme: "HTTP"}},
								InitialDelaySeconds: 5,
								TimeoutSeconds:      5,
								PeriodSeconds:       5,
								SuccessThreshold:    5,
								FailureThreshold:    5,
							},
						},
					},
				},
			},
		},
	}
}

//Nightly manages the night settings for this project
func (d *Demodeployment) Nightly() {
	d.Deployment.Spec.Replicas = int32Ptr(1) // reduce replica count
	d.Deployment.Spec.Template.Spec.Containers[0].Resources = apiv1.ResourceRequirements{
		Limits: apiv1.ResourceList{
			"cpu":    resource.MustParse("50m"),
			"memory": resource.MustParse("50M"),
		},
		Requests: apiv1.ResourceList{
			"cpu":    resource.MustParse("25m"),
			"memory": resource.MustParse("25M"),
		},
	} // decrease request and limit for use during the night
}

//Dayli manages the day settings for this project
func (d *Demodeployment) Dayli() {
	d.Deployment.Spec.Replicas = int32Ptr(5) // reduce replica count
	d.Deployment.Spec.Template.Spec.Containers[0].Resources = apiv1.ResourceRequirements{
		Limits: apiv1.ResourceList{
			"cpu":    resource.MustParse("150m"),
			"memory": resource.MustParse("150M"),
		},
		Requests: apiv1.ResourceList{
			"cpu":    resource.MustParse("75m"),
			"memory": resource.MustParse("75M"),
		},
	} // increase request and limit for use during the day
}
