/*
Copyright 2021 Johannes Hölzel.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"path/filepath"
	"time"

	"github.com/jhoelzel/cronor/pkg/cronrun"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		//this will set the kubeconfig automatically if no path is given and the file exists
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		//otherwise we prepare for incluser load
		kubeconfig = flag.String("kubeconfig", "", "(optional) absolute path to the kubeconfig file")
	}
	namespace := flag.String("namespace", "default", "namespace for the kubernetes deployments. example: default")
	deploymentName := flag.String("name", "my-supercool-deployment", "Name for the deployment. example: my-supercool-deployment")
	strLocation := flag.String("location", "Europe/Berlin", "location for the timezone setting. example: Europe/Berlin")
	flag.Parse()
	//prepare the timezone for correct timing
	location, err := time.LoadLocation(*strLocation)
	if err != nil {
		panic(err.Error())
	}
	runner, err := cronrun.NewRunner(*kubeconfig, *namespace, *deploymentName)
	if err != nil {
		panic(err.Error())
	}
	t := time.Now().In(location)
	runner.ActiveStart = time.Date(t.Year(), t.Month(), t.Day(), 5, 0, 0, 0, t.Location())
	runner.ActiveEnd = time.Date(t.Year(), t.Month(), t.Day(), 18, 0, 0, 0, t.Location())
	runner.Now = t
	err = runner.Run()
	if err != nil {
		panic(err.Error())
	}

}
