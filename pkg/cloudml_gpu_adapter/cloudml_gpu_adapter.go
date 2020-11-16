package cloudml_gpu_adapter

import (
	"log"
	"os"
	"context"

	"k8s.io/client-go/rest"
	k8s "k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	gpuCardLabel = "xiaomi/gpu-card"
	gpuMemLabel = "xiaomi/gpu-mem"
	resourcePrefix = "cloudml.gpu/"
	resourceName = "" //default
	resourceSock = "cloudml.sock"
)

func genResourceName() bool{
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Printf(err.Error())
		return false
	}
	client, err := k8s.NewForConfig(config)
	if err != nil {
		log.Printf(err.Error())
		return false
	}
	nodeName := os.Getenv("NODE_NAME")
	if nodeName == ""{
		log.Printf("Get env node name error")
		return false
	}
	nodeInf, err := client.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
	if err != nil {
		log.Printf("Get node info error %s", err.Error())
		return false
	}

	gpuCard, ok := nodeInf.Labels[gpuCardLabel]
	if !ok{
		log.Printf("Get node label error")
		return false
	}
	gpuMem, ok := nodeInf.Labels[gpuMemLabel]
	if !ok{
		log.Printf("Get node label error")
	}
	resourceName = resourcePrefix + gpuCard + "-" + gpuMem
	if resourceName == ""{
		log.Printf("Generate resource name error")
		return false
	}
	return true
}

func GetGpuResourceName() string  {
	return resourceName
}

func GetGpuResourceSock() string {
	return resourceSock
}

func init()  {
	log.Println("Generate resource name")
	if !genResourceName(){
		log.Println("Generate resource name error")
		os.Exit(1)
	}
}
