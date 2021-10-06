//package cronrun defines the cronrunner and its methods
package cronrun

import (
	"fmt"

	"github.com/jhoelzel/cronor/pkg/demo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Run executes the main tasks for our program
func (r *Runner) Run() (err error) {
	err = r.prepareNamespace()
	if err != nil {
		err = fmt.Errorf("failed to prepare the namespace: %v", err)
		return
	}
	err = r.prepareDeployment()
	if err != nil {
		err = fmt.Errorf("failed to prepare the deployment: %v", err)
		return
	}
	err = r.prepareService()
	if err != nil {
		err = fmt.Errorf("failed to prepare the service: %v", err)
	}
	return
}

//prepareDeployment checks if the current deployment exists or deploys a new instance.
//it also contains different resource limits for daily and nightly activities
func (r *Runner) prepareDeployment() (err error) {
	options := &metav1.GetOptions{}
	err = r.GetDeployment(*options)
	if err != nil {
		// we don have a deployment with that name in our namespace so lets cerate one
		createOptions := &metav1.CreateOptions{}
		//reset the deployment because its not found
		r.demo = demo.NewDemoDeployment(r.demo.Name, r.namespace)
		err = r.CreateDeployment(*createOptions)
		if err != nil {
			err = fmt.Errorf("failed to create the deployment: %v", err)
			return
		}
	}

	if inTimeSpan(r.ActiveStart, r.ActiveEnd) {
		r.demo.Dayli()
	} else {
		r.demo.Nightly()
	}
	updateOptions := &metav1.UpdateOptions{}
	err = r.UpdateDeployment(*updateOptions)
	if err != nil {
		err = fmt.Errorf("failed to update the deployment: %v", err)
	}
	return
}

//prepareService makes sure the service for our deployment is running
func (r *Runner) prepareService() (err error) {
	options := &metav1.GetOptions{}
	err = r.GetService(*options)
	if err != nil {
		//our service has not been found so we create a new one
		//we reset the deployment with our serach so reinitialize its service
		r.demo = demo.SetDemoDeployment(r.demo.Name, r.demo.Namespace, r.demo.Deployment)
		createOptoions := &metav1.CreateOptions{}
		err = r.CreateService(*createOptoions)
		if err != nil {
			err = fmt.Errorf("failed to create the service: %v", err)
		}
	}
	return
}

//prepareService makes sure the service for our namespace is running
func (r *Runner) prepareNamespace() (err error) {
	options := &metav1.GetOptions{}
	err = r.GetNamespace(*options)
	if err != nil {
		//our namespace has not been found so we create a new one
		createOptoions := &metav1.CreateOptions{}
		err = r.CreateNamespace(*createOptoions)
		if err != nil {
			err = fmt.Errorf("failed to create the namespace: %v", err)
		}
	}
	return
}
