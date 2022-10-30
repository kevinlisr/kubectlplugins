package config

import (
	"flag"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
)

type K8sConfig struct {

}
func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

func (*K8sConfig) K8sRestConfig() *rest.Config {
	// 使用当前上下文环境
	var cliKubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		cliKubeconfig = flag.String("cliKubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		cliKubeconfig = flag.String("cliKubeconfig", "", "absolute path to the kubeconfig file")

	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *cliKubeconfig)
	if err != nil {
		panic(err.Error())
	}
	//// 根据指定的 config 创建一个新的 clientSet
	//clientSet, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	panic(err.Error())
	//}
	//return clientSet
	return config
}
// 要创建两个获取config函数，否则报错：tmp/go-build849215055/b001/exe/test flag redefined: kubeconfig
// panic: /tmp/go-build849215055/b001/exe/test flag redefined: kubeconfig

func (*K8sConfig) K8sDcConfig() *rest.Config {
	// 使用当前上下文环境
	var dcKubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		dcKubeconfig = flag.String("dcKubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		dcKubeconfig = flag.String("dcKubeconfig", "", "absolute path to the kubeconfig file")

	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *dcKubeconfig)
	if err != nil {
		panic(err.Error())
	}
	//// 根据指定的 config 创建一个新的 clientSet
	//clientSet, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	panic(err.Error())
	//}
	//return clientSet
	return config
}


func (this *K8sConfig) InitClient() *kubernetes.Clientset {
	c, err := kubernetes.NewForConfig(this.K8sRestConfig())
	if err != nil{
		log.Fatal(err)
	}
	return c
}

func (this *K8sConfig) InitDynamicClient() dynamic.Interface {
	client, err := dynamic.NewForConfig(this.K8sDcConfig())
	if err != nil{
		log.Fatal(err)
	}
	return client
}