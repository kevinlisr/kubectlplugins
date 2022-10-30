package lib

import (
	"context"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"log"
	"os"
	"regexp"
)



var listCmd = &cobra.Command{
	Use: "list",
	Short: "list pods",
	Example: "kubectl pods list [flags]",
	SilenceUsage: true,
	RunE: func(c *cobra.Command, args []string) error {
		ns, err := c.Flags().GetString("namespace")
		if err != nil {
			log.Fatalln(err)
		}
		if ns == ""{
			ns = "default"
		}
		list, err := client.CoreV1().Pods(ns).List(context.Background(), v1.ListOptions{
			LabelSelector: Labels,
			FieldSelector: fields,
		})
		if err != nil {
			log.Fatalln(err)
		}
		//podsJson,_ := json.Marshal(list)

		table := tablewriter.NewWriter(os.Stdout)
		commonHeaders := []string{"Name", "Namespace", "Ip","Phase"}

		if ShowLables{
			commonHeaders = append(commonHeaders,"tag")
		}

		table.SetHeader(commonHeaders)

		for _,pod :=range list.Items{
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

			//podRow = []string{pod.Name,pod.Namespace,pod.Status.PodIP,string(pod.Status.Phase)}
			if ShowLables{
				podRow = append(podRow,Map2String(pod.Labels))
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


	},
}
