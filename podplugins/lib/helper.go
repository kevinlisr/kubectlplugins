package lib

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"sigs.k8s.io/yaml"

	v1 "k8s.io/api/core/v1"
	"log"
)

func checkError(msg string, err error) {
	if err != nil{
		errMsg:= fmt.Sprintf("%s:%s\n", msg, err.Error())
		log.Fatalln(errMsg)
	}
}

func InitHeader(table *tablewriter.Table) []string {
	return nil
}
func FilterListByJSON(list *v1.PodList)  {}

// huo qu pod xiang xi xin xi
func getPodDetail(args []string) {
	if len(args) == 0{
		log.Println("pod name is required")
	}
	podName := args[0]
	pod, err := fact.Core().V1().Pods().Lister().
		Pods("default").Get(podName)
	if err!=nil{
		log.Println(err)
		return
	}
	b, err := yaml.Marshal(pod)
	if err != nil{
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}