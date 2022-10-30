package lib

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)



// rootCmd represents the base command when called without any subcommands


//func MergeFlags(cmd *cobra.Command)  {
//	var Lables = "app=nginx"
//	var podCmd = cmd
//	err := podCmd.Execute()
//	if err != nil {
//		os.Exit(1)
//	}
//	podCmd.Flags().BoolP("pods", "p", false, "Help message for pods")
//	podCmd.Flags().BoolP("namespace", "ns", false, "Help message for namespace")
//	cmd.Flags().StringVar(&Lables, "labels", "", "kubectl pods --labels=\"app=nginx\"")
//}

func MergeFlags(cmd *cobra.Command)  {
	//var Namespace   = ""
	//var args   = ""

	//cfgFlags := genericclioptions.NewConfigFlags(true)
	//cfgFlags.AddFlags(cmd.PersistentFlags())

	//namespace, _ := cmd.PersistentFlags().GetString("-n")
	//cmd.Flags().StringVar(&Namespace, "namespace", "", "kubectl pods --namespace=\"kube-system\"")
	//cmd.Flags().StringVar(&args, "-n", "", "kubectl pods --namespace=\"kube-system\"")
	fmt.Println("ding yi namespace")
	cmd.Flags().StringP("namespace","n","","kubectl pods --namespace=\"kube-system\"")
	//cmd.PersistentFlags().StringVarP(&Namespace,"namespace","n","namespace","kubectl -n ")
}

func RunCmd(run func(cmd *cobra.Command, args []string) error) {
	cmd := &cobra.Command{
		Use:          "kubectl pods [flags]",
		Short:        "list pods",
		Example:      "kubectl pods [flags]",
		SilenceUsage: true,
		RunE:         run,
	}
	MergeFlags(cmd)

	//添加参数
	//BoolVar用来支持 是否
	//cmd.Flags().BoolVar(&ShowLabels, "show-labels", false, "kubectl pods --show-labels")
	//cmd.Flags().StringVar(&Labels, "labels", "", "kubectl pods --labels=\"app=nginx\"")

	err := cmd.Execute()
	fmt.Println("stop exec  cmd")
	if err != nil {
		log.Fatalln(err,"exec bao cuo")
	}
}


//func Execute() {
//	err := rootCmd.Execute()
//	if err != nil {
//		os.Exit(1)
//	}
//}

//func init() {
//
//
//	rootCmd.Flags().BoolP("pods", "p", false, "Help message for pods")
//
//}
