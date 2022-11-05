package lib

import (
	"context"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/json"
	"os"
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
		Pods(args[1]).Get(podName)
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

func setTable(table *tablewriter.Table) {
	// she zhi TABLE de yang shi , bu zhong yao, kankan jiu hao
}

var eventHeaders = []string{"EVENT", "REASON", "RESOURCE", "MESSAGES"}

func printEvent(events []*v1.Event)  {
	table := tablewriter.NewWriter(os.Stdout)
	// set header
	table.SetHeader(eventHeaders)
	for _,e := range events {
		podRow := []string{e.Type, e.Reason,
			fmt.Sprintf("%s/%s",e.InvolvedObject.Kind,e.InvolvedObject.Name),e.Message}
		table.Append(podRow)
	}
	setTable(table)
	// jin xing xuan ran
	table.Render()
}

func getPodDetailByJSON(podName, path ,nameSpace string, cmd *cobra.Command) {
	//ns, err := cmd.Flags().GetString("namespace")
	//if err != nil {
	//	log.Println("error ns param")
	//	return
	//}
	if nameSpace == "" {nameSpace = "default"}
	pod,err := fact.Core().V1().Pods().Lister().Pods(nameSpace).Get(podName)
	if err != nil {
		log.Println(err)
		return
	}

	// get resource events
	if path == PodEventType {
		eventList, err := fact.Core().V1().Events().Lister().List(labels.Everything())
		if err != nil {
			log.Println(err)
			return
		}
		podEvents := []*v1.Event{}
		for _, e := range eventList{
			if e.InvolvedObject.UID == pod.UID{
				podEvents = append(podEvents, e)
			}
		}
		printEvent(podEvents)
		// dao zhe jiu bu xu yao wang xia zou le ,zhi jie return
		return
	}

	// get resource logs
	if path == PodLogType{
		req := client.CoreV1().Pods(nameSpace).GetLogs(pod.Name,&v1.PodLogOptions{})
		//Pod, err := client.CoreV1().Pods(nameSpace).Get(context.Background(), pod.Name, metav1.GetOptions{})
		//fmt.Println("huo qu pod",Pod)

		ret := req.Do(context.Background())
		fmt.Println(ret.Get())
		b, err := ret.Raw()
		if err != nil {
			fmt.Println("get logs error")
			log.Println(err)
			return
		}
		fmt.Println(string(b))
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