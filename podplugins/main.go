package main

import (

	"github.com/kevinlisr/podplugins/lib"

)

//type K8sConfig struct {
//}
//
//func NewK8sConfig() *K8sConfig {
//	return &K8sConfig{}
//}
//
//func (*K8sConfig) K8sRestConfig() *rest.Config {
//	// 使用当前上下文环境
//	var cliKubeconfig *string
//	if home := homedir.HomeDir(); home != "" {
//		cliKubeconfig = flag.String("cliKubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
//	} else {
//		cliKubeconfig = flag.String("cliKubeconfig", "", "absolute path to the kubeconfig file")
//
//	}
//	flag.Parse()
//
//	config, err := clientcmd.BuildConfigFromFlags("", *cliKubeconfig)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	return config
//}
//
//func (this *K8sConfig) InitClient() *kubernetes.Clientset {
//	cfgFlags := genericclioptions.NewConfigFlags(true)
//	config, err := cfgFlags.ToRawKubeConfigLoader().ClientConfig()
//	if err != nil {
//		log.Fatalln(err)
//	}
//	c, err := kubernetes.NewForConfig(config)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	return c
//	//c, err := kubernetes.NewForConfig(this.K8sRestConfig())
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//return c
//}
//
//var ShowLables bool
//var Labels string
//var fields string
//var PodName string
//
//func MergeFlags(cmd *cobra.Command) {
//	cmd.Flags().StringP("namespace", "n", "", "kubectl pods --namespace=\"kube-system\"")
//	//cmd.Flags().Bool("show-labels",false,"kubectl pods --show-labels")
//	cmd.Flags().BoolVar(&ShowLables,"show-labels",false,"kubectl pods --show-labels")
//	cmd.Flags().StringVar(&Labels,"labels","",
//		"kubectl pods --labels app=ngx or kubectl pods --labels=\"app=ngx,version=v1\"")
//	cmd.Flags().StringVar(&fields,"fields","",
//		"kubectl pods --fields=\"status.phase=Running\"")
//	cmd.Flags().StringVar(&PodName,"name","",
//		"kubectl pods --name=\"^ng\"")
//}
//
//func RunCmd(run func(cmd *cobra.Command, args []string) error) {
//	cmd := &cobra.Command{
//		Use:          "kubectl pods [flags]",
//		Short:        "list pods",
//		Example:      "kubectl pods [flags]",
//		SilenceUsage: true,
//		RunE:         run,
//	}
//	MergeFlags(cmd)
//
//	//添加参数
//	//BoolVar用来支持 是否
//	//cmd.Flags().BoolVar(&ShowLabels, "show-labels", false, "kubectl pods --show-labels")
//	//cmd.Flags().StringVar(&Labels, "labels", "", "kubectl pods --labels=\"app=nginx\"")
//
//	err := cmd.Execute()
//	fmt.Println("stop exec  cmd")
//	if err != nil {
//		log.Fatalln(err, "exec bao cuo")
//	}
//}

//func run(cmd *cobra.Command, args []string) error {
//	ns, err := cmd.Flags().GetString("namespace")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(ns)
//	return nil
//}

//func run(c *cobra.Command,args []string) error {
//	client := NewK8sConfig().InitClient()
//	ns,err := c.Flags().GetString("namespace")
//	fmt.Println(ns)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	if ns==""{ns="default"}
//
//	list, err := client.CoreV1().Pods(ns).List(context.Background(),
//		v1.ListOptions{
//			LabelSelector: Labels,
//			FieldSelector: fields,
//		})
//	if err != nil {
//		log.Fatalln(err,"huo qu pod list")
//	}
//
//	//for _, p := range list.Items{
//	//	podsJson,_ := json.Marshal(p)
//	//}
//	podsJson,_ := json.Marshal(list)
//
//	err = WriteFile("pods.json", []byte(podsJson), 0666)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	table := tablewriter.NewWriter(os.Stdout)
//	commonHeaders := []string{"Name", "Namespace", "Ip","Phase"}
//
//	if ShowLables{
//		commonHeaders = append(commonHeaders,"tag")
//	}
//
//	table.SetHeader(commonHeaders)
//
//	for _,pod :=range list.Items{
//		//fmt.Println(pod.Name)
//		p, err := json.Marshal(pod)
//		if err != nil {
//			log.Fatalln(err)
//		}
//		ret := gjson.Get(string(p), "metadata.name")
//
//		var podRow  []string
//		if m,err := regexp.MatchString(PodName,ret.String());err == nil && m {
//			podRow = []string{pod.Name,pod.Namespace,pod.Status.PodIP,string(pod.Status.Phase)}
//
//		}
//
//		//podRow = []string{pod.Name,pod.Namespace,pod.Status.PodIP,string(pod.Status.Phase)}
//		if ShowLables{
//			podRow = append(podRow,lib.Map2String(pod.Labels))
//		}
//		table.Append(podRow)
//	}
//	table.SetAutoWrapText(false)
//	table.SetAutoFormatHeaders(true)
//	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
//	table.SetAlignment(tablewriter.ALIGN_LEFT)
//	table.SetCenterSeparator("")
//	table.SetColumnSeparator("")
//	table.SetRowSeparator("")
//	table.SetHeaderLine(false)
//	table.SetBorder(false)
//	table.SetTablePadding("\t") // pad with tabs
//	table.SetNoWhiteSpace(true)
//	table.Render()
//	return nil
//}
//
//
//// 通用的文件打开函数(综合和 Create 和 Open的作用)
//// OpenFile第二个参数 flag 有如下可选项
////    O_RDONLY  文件以只读模式打开
////    O_WRONLY  文件以只写模式打开
////    O_RDWR   文件以读写模式打开
////    O_APPEND 追加写入
////    O_CREATE 文件不存在时创建
////    O_EXCL   和 O_CREATE 配合使用,创建的文件必须不存在
////    O_SYNC   开启同步 I/O
////    O_TRUNC  打开时截断常规可写文件
//func WriteFile(filename string, data []byte, perm os.FileMode) error {
//	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
//	if err != nil {
//		return err
//	}
//	n, err := f.Write(data)
//	if err == nil && n < len(data) {
//		err = io.ErrShortWrite
//	}
//	if err1 := f.Close(); err == nil {
//		err = err1
//	}
//	return err
//}

func main() {

	lib.RunCmd()
}
