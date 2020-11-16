package main

import (
	"cloudml_adapter/pkg/cloudml_gpu_adapter"
	"fmt"
	_ "cloudml_adapter/pkg/cloudml_gpu_adapter"
)

var d = cloudml_gpu_adapter.GetGpuResourceSock()

func main()  {
	fmt.Println(d)
	fmt.Println(cloudml_gpu_adapter.GetGpuResourceName())
}
