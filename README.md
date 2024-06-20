# Pod and Deployment Lister App

an app that uses client go to talk to kube api server to do things

## Running Locally
We can run this program locally

In <i>main.go</i> find the where its locating your kubeconfig file and change the path.


`go build`

`./lister`

It should list the pods and deployments running in the cluster.

## Containerization
To run this app as a deployment in our cluster:
Because kubernetes is container orchestrator, we have to containerize our app

Make sure to compile the binary to run on Linux since thats what our container runs on.

`GOOS=linux go build -o listerLinux`

`docker build -t <imagename>`

Lets have a look into the yaml. This command will generate a yaml based on our container image. Dont spend too much time on this, it simply creates a deployment with a pod running our container.

`k create deployment lister --image <imagename> --dry-run=client -oyaml > lister.yaml`


Locally, our app uses <i>kubeconfig</i> to authenticate with the kluster. <i>kubeconfig</i> doesn't exist in cluster.
There is a method provided by client-go/rest that can get configuration while running inside cluster.

InClusterConfig()

How?
We know that pods gets mounted default service account to provide identity to authenticate to api server. We can use this service account to talk to server. InClusterConfig will look at the path where service account gets mounted

... reads the details from the path where kubelet mounts the default service account

uses tokenFile to authenticate against k8s cluster

Now, we gotta create RBAC  since default service account not allowed to list pods and deployments.

`k create role poddepl --resource pods,deployments --verb list`

`k create rolebinding poddepl --role poddepl --serviceaccount default:default`

To test this:

First, deploy the example manifest:

`k apply -f manifests/example.yaml`

All this does it create a deployment with some empty pods on the cluster.

Now you can deploy the manifest from above and check the logs of the pod it created to see it running on the cluster.

`k apply -f lister.yaml`

<u>Summary</U>

We use InClusterConfig() to configure app so it can run inside kubernetes cluster as a pod


