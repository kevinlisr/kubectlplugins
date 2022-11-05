package lib

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/json"
	"log"
	"os"
	"regexp"
)

var cacheCmd = &cobra.Command{
	Use: "cache",
	Short: "pods by cache",
	Hidden: true,
	RunE: func(c *cobra.Command, args []string) error {
		//ns, err = c.Flags().GetString("namespace")
		//fmt.Println("cacheCmd  get namespace is:", ns)
		//if err != nil {
		//	log.Fatalln(err)
		//}


		//arg := fmt.Sprintf("%s", args[1])
		//if ns == ""{ns=args[1]}
		if len(args) != 0{
			fmt.Println("args[1] fu zhi")
			ns=args[1]
			fmt.Printf("++++++++++++++++%s\n", args[1])
		}else {
			if ns == ""{ns="default"}
			//ns="default"
		}

		pods, err := fact.Core().V1().Pods().Lister().Pods(ns).
			List(labels.Everything())
		if err != nil {
			return err
		}
		fmt.Println("cong huan cun zhong qu")
		table := tablewriter.NewWriter(os.Stdout)

		commonHeaders := []string{"Name", "Namespace", "Ip","Phase"}
		//
		//if ShowLables{
		//	commonHeaders = append(commonHeaders,"tag")
		//}

		table.SetHeader(commonHeaders)

		for _,pod :=range pods{
			//fmt.Println(pod.Name)
			p, err := json.Marshal(pod)
			if err != nil {
				log.Fatalln(err)
			}
			ret := gjson.Get(string(p), "metadata.name")

			var podRow  []string
			if m,err := regexp.MatchString(PodName,ret.String());err == nil && m {
				podRow = []string{pod.Name,pod.Namespace,pod.Status.PodIP,string(pod.Status.Phase)}

			}

			table.Append(podRow)
		}
		table.SetAutoWrapText(false)
		table.SetAutoFormatHeaders(true)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetRowSeparator("")
		table.SetHeaderLine(false)
		table.SetBorder(false)
		table.SetTablePadding("\t") // pad with tabs
		table.SetNoWhiteSpace(true)
		table.Render()
		return nil
		////////////////////////////////////////

	},
}