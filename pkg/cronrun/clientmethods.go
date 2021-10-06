//package cronrun defines the cronrunner and its methods
package cronrun

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

//GetDeployment checks if a deployment exists in the namespace
func (r *Runner) GetDeployment(options metav1.GetOptions) (err error) {
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		r.demo.Deployment, err = r.clientset.AppsV1().Deployments(r.namespace).Get(context.TODO(), r.demo.Name, options)
		if err != nil {
			err = fmt.Errorf("failed to get the deployment from the kubernetes api: %v", err)
		}
		return err
	})
	return
}

//CreateDeployment creates deployment in the namespace
func (r *Runner) CreateDeployment(options metav1.CreateOptions) (err error) {
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		r.demo.Deployment, err = r.clientset.AppsV1().Deployments(r.namespace).Create(context.TODO(), r.demo.Deployment, options)
		if err != nil {
			err = fmt.Errorf("failed to create the deployment with the kubernetes api: %v", err)
		}
		return err
	})
	return
}

//updateDeployment updates and existing deployment in the namespace
func (r *Runner) UpdateDeployment(updateOptions metav1.UpdateOptions) (err error) {
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		r.demo.Deployment, err = r.clientset.AppsV1().Deployments(r.namespace).Update(context.TODO(), r.demo.Deployment, updateOptions)
		if err != nil {
			err = fmt.Errorf("failed to update the deployment with the kubernetes api: %v", err)
		}
		return err
	})
	return
}

//GetService checks if a service exists in the namespace
func (r *Runner) GetService(options metav1.GetOptions) (err error) {
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		r.demo.Service, err = r.clientset.CoreV1().Services(r.namespace).Get(context.TODO(), r.demo.Name+"-loadblancer-service", options)
		if err != nil {
			err = fmt.Errorf("failed to get the service with from kubernetes api: %v", err)
		}
		return err
	})
	return
}

//UpdateService updates a service
func (r *Runner) UpdateService(options metav1.UpdateOptions) (err error) {
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		r.demo.Service, err = r.clientset.CoreV1().Services(r.namespace).Update(context.TODO(), r.demo.Service, options)
		if err != nil {
			err = fmt.Errorf("failed to updae the service with the kubernetes api: %v", err)
		}
		return err
	})
	return
}

//CreateService creates a service
func (r *Runner) CreateService(options metav1.CreateOptions) (err error) {
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		r.demo.Service, err = r.clientset.CoreV1().Services(r.namespace).Create(context.TODO(), r.demo.Service, options)
		if err != nil {
			err = fmt.Errorf("failed to create the service with the kubernetes api: %v", err)
		}
		return err
	})
	return
}

//GetNamespace checks if a service exists in the Namespace
func (r *Runner) GetNamespace(options metav1.GetOptions) (err error) {
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		_, err = r.clientset.CoreV1().Namespaces().Get(context.TODO(), r.namespace, options)
		if err != nil {
			err = fmt.Errorf("failed to get the service with from kubernetes api: %v", err)
		}
		return err
	})
	return
}

//CreateNamespace creates a namespace
func (r *Runner) CreateNamespace(options metav1.CreateOptions) (err error) {
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		r.demo.NamespaceObj, err = r.clientset.CoreV1().Namespaces().Create(context.TODO(), r.demo.NamespaceObj, options)
		if err != nil {
			err = fmt.Errorf("failed to create the service with the kubernetes api: %v", err)
		}
		return err
	})
	return
}
