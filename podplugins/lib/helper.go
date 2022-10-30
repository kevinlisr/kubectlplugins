package lib

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"k8s.io/apimachinery/pkg/util/json"
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
		return
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

func getPodDetailByJSON(podName, path string, cmd *cobra.Command) {
	ns, err := cmd.Flags().GetString("namespace")
	if err != nil {
		log.Println("error ns param")
		return
	}
	if ns == "" {ns = "default"}
	pod,err := fact.Core().V1().Pods().Lister().Pods(ns).Get(podName)
	if err != nil {
		log.Println(err)
		return
	}
	jsonStr, _ := json.Marshal(pod)
	ret := gjson.Get(string(jsonStr),path)
	if !ret.Exists(){
		log.Println("No corresponding content was found "+path)
		return
	}
	if !ret.IsObject() && !ret.IsArray(){
		fmt.Println(ret.Raw)
		return
	}
	tempMap := make(map[string]interface{})
	err = yaml.Unmarshal([]byte(ret.Raw),&tempMap)
	if err != nil {
		log.Println(err)
		return
	}
	b,_ := yaml.Marshal(tempMap)
	fmt.Println(string(b))
}