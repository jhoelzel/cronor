# Cronor

Cronor is a kubernetes cron job image with one task: change our deployment depending if it's day or night.
This is a showcase of how easily the kubernetes api can be implemented directly into your code in multiple ways.
I chose the cronjob because it is perfect for our scenario: once an hour check if our deployment is correctly scaled and deployed, but the packages could easily be implemented in other applications like webapplications, GRPC servers and more.

## What does it do?

Between 05:00 and 18:00 our fictional cluster is under a lot of stress. Therefore we need to increase the memory limits and of course the replicacount for our application to still be highly available.
This is exactly what this program does:

- Check if our Namespace exist. If not create it.
- Check if our Deployment exists. If not create it.
- If our deployment exists: do we need to scale it up because its during the day?
- If our deployment exists: do we need to scale it down because its during the night?
- Along with everything: make sure our namespace atually exists.

## What is it build with

- Golang 1.16
- Distroless container image for the cronjob. (the entire container is only 31mb and those are mostly dependencies)
- In the makefile I am linking my local microk8s container registry, you will have to replace it if you want to build the image.
- the kubeconfig is either given as a flag, simply in your homedir or even incluster configuration can be used, so this simply executes "in cloud".
- versioning for your containers
- versioning for your app applied when its build through ldflags
- Makefile for testing and deployment locally, which includes a lot of goodies like kube-apply, docker image building and much more:

```
Usage:
  make <target>

General
  help             Display this help.

Development
  fmt              Run go fmt against code.
  vet              Run go vet against code.
  test             Run tests.
  run              Run.
  commithistory    create the commithistory in a nice format

Build
  clean            remove previous binaries
  build            build a version of the app, pass Buildversion, Comit and projectname as build arguments
  docker-build     Build the docker image and tag it with the current version and :latest
  docker-run       Build the docker image and tag it and run it in docker
  docker-push      push your image to the docker hub

Kubernetes
  kube-manifests   generated the kubernetes manifests and replaces variables in them
  kube-clean       removes release manifests
  kube-apply       apply kube manifests
  kube-remove      remove kube manifests
  kube-ns          create desired namespace
  kube-removens    remove the desired namespace
  kube-renew       build, docker-build, remove existing deployment, deploy  
  ```

## What are we deploying?

The cronjob deploys the simpleapp container that I use for trainings. You can find it here: <https://github.com/jhoelzel/simpleapp>.

## How do we access the deployment?

With:
kubectl get service -n "{{.APP_NAME}}-loadbalancer-service"
You will reveal the load balancer that this project shedules.
If you have used the defaults without changing anything you will find the service as "my-supercool-deployment-loadblancer-service" and the deployment as "my-supercool-deployment"

## What assumptions have I made?

- The project should be able to compile offline without the use of the internet, thats why I commited the vendor files
- The tool can either be run locally directly, in the cloud, or simply in docker. Therefore i provided multiple ways to feed the kubeconfig to the client.
- It currently will not need RBAC because this is a demo project. Of course RBAC rules should be given to the cronjob which in turn executes the container.
- Testing the kubernetes API is not really in my realm, it resides with the go-client. Therefore I skipped test cases. There is the fake-client which could be used, but in our scenario a quick k get pods will be sufficient.

## What are you running this on?

Locally I am developing in one of my alpine dev containers.
My kubernetes distribution of choice for local development is microk8s because its simple to use and with metallb plays very well with my OpenWRT router.
I dont like to run workloads directly on my surfacebook therefore I have a seperate Server in my network where everything runs on. With a Wireguard tunnel on OpenWRT I have no problems when I am in the green office either ;)

Versions:
```
Client Version: version.Info{Major:"1", Minor:"21", GitVersion:"v1.21.4", GitCommit:"3cce4a82b44f032d0cd1a1790e6d2f5a55d20aae", GitTreeState:"clean", BuildDate:"2021-08-11T18:16:05Z", GoVersion:"go1.16.7", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"21+", GitVersion:"v1.21.5-3+83e2bb7ee39726", GitCommit:"83e2bb7ee3972654beca02a12a94777da22d6669", GitTreeState:"clean", BuildDate:"2021-09-28T15:36:44Z", GoVersion:"go1.16.8", Compiler:"gc", Platform:"linux/amd64"}                                                                                                                                                                                                   /0.1s
   ```
