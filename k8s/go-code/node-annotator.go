// a node annotator that updates the available values upon pod events
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func main() {

	// creates the in-cluster config
	fmt.Println("Let's try in-cluster config first!")
	config, err := rest.InClusterConfig()

	if err != nil {
		// panic(err.Error())
		fmt.Printf("Failed to use in-cluster config, use out-cluster config instead!")
		home := homeDir()

		var kubeconfig *string
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		flag.Parse()

		// use the current context in kubeconfig
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	event_watcher, err := clientset.CoreV1().Pods("").Watch(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for {
		e := <-event_watcher.ResultChan()
		if e.Object == nil {
			return
		}
		pod, ok := e.Object.(*apiv1.Pod)
		if !ok {
			continue
		}

		podnode := pod.Spec.NodeName

		node, err := clientset.CoreV1().Nodes().Get(podnode, metav1.GetOptions{})
		if err != nil {
			log.Panicf("Node "+podnode+" cannot be found!", err.Error())
			continue
		}

		cpu_avail := node.Status.Allocatable.Cpu().MilliValue()
		mem_avail := node.Status.Allocatable.Memory().Value()
		gpu_avail := node.Status.Allocatable.NvidiaGPU().Value()
		podcnt_avail := node.Status.Allocatable.Pods().Value()

		all_pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{FieldSelector: "spec.nodeName=" + podnode})
		if err != nil {
			log.Panicf("All pods on "+podnode+" cannot be found!", err.Error())
			continue
		}

		for i := 0; i < len(all_pods.Items); i += 1 {
			all_containers := all_pods.Items[i].Spec.Containers

			pod_cpu_request := int64(0)
			pod_mem_request := int64(0)
			pod_gpu_request := int64(0)
			for j := 0; j < len(all_containers); j += 1 {
				pod_cpu_request += all_containers[j].Resources.Requests.Cpu().MilliValue()
				pod_mem_request += all_containers[j].Resources.Requests.Memory().Value()
				pod_gpu_request += all_containers[j].Resources.Requests.NvidiaGPU().Value()
			}

			cpu_avail = cpu_avail - pod_cpu_request
			mem_avail = mem_avail - pod_mem_request
			gpu_avail = gpu_avail - pod_gpu_request
			podcnt_avail = podcnt_avail - 1
		}

		node.Annotations["cpu_avail"] = strconv.FormatInt(cpu_avail, 10)
		node.Annotations["mem_avail"] = strconv.FormatInt(mem_avail, 10)
		node.Annotations["gpu_avail"] = strconv.FormatInt(gpu_avail, 10)
		node.Annotations["podcnt_avail"] = strconv.FormatInt(podcnt_avail, 10)

		clientset.CoreV1().Nodes().Update(node)
	}
}
