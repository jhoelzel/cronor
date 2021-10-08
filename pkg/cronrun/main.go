//package cronrun defines the cronrunner and its methods
package cronrun

import (
	"fmt"
	"os"
	"time"

	"github.com/jhoelzel/cronor/pkg/demo"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Runner struct {
	ActiveStart time.Time
	ActiveEnd   time.Time
	Now         time.Time
	config      *rest.Config
	clientset   *kubernetes.Clientset
	namespace   string
	demo        *demo.Demodeployment
}

//NewRunner returns a new Runner with namespace and config set
func NewRunner(kubeconfig, namespace, deploymentName string) (r Runner, err error) {
	if namespace == "" {
		namespace = apiv1.NamespaceDefault
	}
	r.namespace = namespace
	//first init to get our defaults
	r.demo = demo.NewDemoDeployment(deploymentName, r.namespace)

	r.config, err = r.getConfig(kubeconfig)
	if err != nil {
		err = fmt.Errorf("failed to get the kubeconfig: %v", err)
		return
	}

	r.clientset, err = kubernetes.NewForConfig(r.config)
	if err != nil {
		err = fmt.Errorf("failed to create the kubernetes clientset: %v", err)
		return
	}

	return
}

//getConfig either loads the kubernetes config from disk or uses the inclusterconfig
func (r *Runner) getConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if os.IsNotExist(err) {
			cfg, err = rest.InClusterConfig()
			if err != nil {
				return nil, err
			}
		}
		if err != nil {
			err = fmt.Errorf("problems loading the kubeconfig: %v", err)
			return nil, err
		}
		return cfg, nil
	}
	cfg, err := rest.InClusterConfig()
	if err != nil {
		err = fmt.Errorf("failed set InClusterConfig: %v", err)
		return nil, err
	}
	return cfg, nil
}
