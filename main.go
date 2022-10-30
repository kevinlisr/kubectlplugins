package main

import (
	"context"
	"fmt"
	"github.com/kevinlisr/config"
	"github.com/kevinlisr/lib"
	"github.com/spf13/cobra"

	"k8s.io/client-go/kubernetes"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"log"
)
var KubernetesConfigFlags *genericclioptions.ConfigFlags
func InitClient() *kubernetes.Clientset {
	cfgFlags := genericclioptions.NewConfigFlags(true)
	config, err := cfgFlags.ToRawKubeConfigLoader().ClientConfig()
	if err != nil {
		log.Fatalln(err)
	}
	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	return c
}


//func MergeFlags(cmd *cobra.Command)  {
//	var Namespace = "namespace"
//	cfgFlags := genericclioptions.NewConfigFlags(true)
//	cfgFlags.AddFlags(cmd.PersistentFlags())
//	cmd.Flags().StringVar(&Namespace, "labels", "", "kubectl pods --labels=\"app=nginx\"")
//}



func run(c *cobra.Command,args []string) error {
	client := config.NewK8sConfig().InitClient()
		ns,err := c.Flags().GetString("namespace")
		fmt.Println(ns)
		if err != nil {
			log.Fatalln(err)
		}
		if ns==""{ns="default"}

		list, err := client.CoreV1().Pods(ns).List(context.Background(),v1.ListOptions{})
		if err != nil {
			log.Fatalln(err,"huo qu pod list")
		}
		for _,pod :=range list.Items{
			fmt.Println(pod.Name)
		}
		return nil
}





func main() {

	lib.RunCmd(run)
}
