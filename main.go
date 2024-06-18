package main

import (
	"context"
	"flag"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	// want to talk to kube api server using client libraries
	// must specify where kube config file otherwise we wont be
	// able to talk to cluster since its secured by authentication
	// and kube config contains these details. This way the apis will
	// be able to figure out which part of cluster to talk to and what
	// perms they have

	// either provide a flag or else use default location
	kubeconfig := flag.String("kubeconfig", "/Users/grantmeyer/.kube/config", "path to kubeconfig file")

	//import client-go library to talk to k8s api server
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("[LOCAL] ignore if running in cluster - error %s,  \n", err.Error())
		//if there is an error, meaning if we are running inside of cluster, use this method from client-go/rest
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("[CLUSTER] error %s getting in cluster config\n", err.Error())
		}

	}

	// create a clientset to talk to k8s api server using kubeconfig
	// a set of clients that can be used to interact with resources from different api versions
	// ex can call pod from corev1 to get get, list, create, delete
	// ex deployment from appsv1 to get, list, create, delete
	// "it gets us clients of all the groups and api versions" but not CRDs and api services
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {

		fmt.Printf("error %s creating clientset from config", err.Error())
	}

	// clientset.CoreV1().Pods("").Get()

	// get all pods from defualt namespace
	// list expects context and list options
	// list options are from api machinery

	//.List returns a list of pods
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		//handle error
		fmt.Printf("error %s getting pods from default namespace\n", err.Error())
	}

	fmt.Println("Pods from default namespace")
	for _, pod := range pods.Items {
		fmt.Printf("%s\n", pod.Name)
	}

	// get all deployments from default namespace
	/*
		deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
		if err != nil {
			fmt.Printf("error %s getting deployments from default namespace\n", err.Error())
		}

		fmt.Println("Deployments from default namespace")
		for _, deployment := range deployments.Items {
			fmt.Printf("%s\n", deployment.Name)
		}
	*/

}
