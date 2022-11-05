package lib

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"k8s.io/apimachinery/pkg/labels"
	"log"
	"regexp"

	"github.com/spf13/cobra"

	"os"

	"strings"
)

//var ns string

func setNamespace(c *cobra.Command,a []string) string {
	ns, err := c.Flags().GetString("namespace")
	if err != nil {
		log.Fatalln(err)
	}

	if ns == ""{

		fmt.Printf("ns is null set ns == %s", a[1])
		ns = a[1]
	}
	return ns
}

func clearConsole()  {
	MyConsoleWriter.EraseScreen()
	MyConsoleWriter.CursorGoTo(0,0)
	MyConsoleWriter.Flush()
}



func executorCmd(cmd *cobra.Command,ns *string) func(in string) {

	return func(in string) {
		in = strings.TrimSpace(in)
		blocks := strings.Split(in, " ")
		args := []string{}
		if len(blocks) > 1{
			args=blocks[1:]
		}

		switch blocks[0] {
		case "exit":
			fmt.Print("Bye!")
			os.Exit(0)

		case "list":
			//InitCache()  // chu shi hua huan cun
			//if err := listCmd.RunE(cmd, []string{});err !=nil{
			//	log.Fatalln(err)
			//}
			fmt.Println("+++++++++++++++",blocks)
			fmt.Println("+++++++++++++++",args)
			if err := cacheCmd.RunE(cmd,args);err !=nil{
				log.Fatalln(err)
			}
		case "get":
			//getPodDetail(args)
			runtea(args,cmd,*ns)
		case "set":
				fmt.Printf("ns is null set ns == %s", args[1])
				*ns = args[1]
		case "clear":
			clearConsole()

		//case "del":
		//	delPod(args,cmd)
		//case "ns":
		//	showNameSpace(cmd)
		case "exec":
			runteaExec(args,cmd)


	  }

	}
}
var suggestions = []prompt.Suggest{
	{"get","GET "},
	{"list","LIST"},
	{"exit","EXIT the interactive window"},
	{"exec","exec the container"},
}

//var suggestions = []prompt.Suggest{
//	// Command
//	{"exec", "pod shell cao zuo "},
//	{"get","get pod xiang xi xin xi"},
//	{"use", "she zhi dang qian ming ming kong jian"},
//	{"del", "shan chu mou ge pod"},
//	{"list","xian shi pod lie biao"},
//	{"clear", "qing chu ping mu"},
//	{"exit","tui chu jiao hu shi"},
//}

var podSuggestions = []prompt.Suggest{
	{"get","get pods details"},
	{"list","show pods list"},
	{"exit","exit the interactive window"},
	{"exec","exec the container"},
}

func getPodsList(ns *string)(ret []prompt.Suggest){
	pods , err :=fact.Core().V1().Pods().Lister().
		Pods(*ns).List(labels.Everything())
	if err != nil{return }
	for _, pod := range pods{
		ret=append(ret,prompt.Suggest{
			Text: pod.Name,
			Description: "node:"+pod.Spec.NodeName+" status:"+
				string(pod.Status.Phase)+" IP:"+pod.Status.PodIP,
		})
	}
	return
}

func parseCmd(w string) (string, string) {
	w = regexp.MustCompile("\\s+").ReplaceAllString(w, " ")
	l := strings.Split(w," ")
	if len(l)>= 2{
		return l[0],strings.Join(l[1:]," ")
	}
	return w,""
}

//func completer(in prompt.Document) []prompt.Suggest {
//	w := in.GetWordBeforeCursor()
//	if w == ""{
//		return []prompt.Suggest{}
//	}
//
//	cmd, opt := parseCmd(in.TextBeforeCursor())
//	//fmt.Println(cmd)
//	if cmd == "get"{
//		return prompt.FilterHasPrefix(getPodsList(ns),opt,true)
//	}
//
//	return prompt.FilterHasPrefix(suggestions,w,true)
//}

func completerCmd(ns *string) func (in prompt.Document) []prompt.Suggest {
	return func (in prompt.Document) []prompt.Suggest {
		w := in.GetWordBeforeCursor()
		if w == ""{
			return []prompt.Suggest{}
		}

		cmd, opt := parseCmd(in.TextBeforeCursor())
		//fmt.Println(cmd)

		//if inArray([]string{"get","del","exec"},cmd){
		//	return prompt.FilterHasPrefix(getPodsList(ns),opt,true)
		//}

		if cmd == "get"{
			return prompt.FilterHasPrefix(getPodsList(ns),opt,true)
		}
		if cmd == "exec"{
			return prompt.FilterHasPrefix(getPodsList(ns),opt,true)
		}

		return prompt.FilterHasPrefix(suggestions,w,true)
	}
}
var ns string
var err error
var MyConsoleWriter = prompt.NewStderrWriter()
var promptCmd = &cobra.Command{
	Use: "prompt",
	Short: "prompt pods",
	Example: "kubectl pods prompt [flags]",
	SilenceUsage: true,
	RunE: func(c *cobra.Command, args []string) error {

		ns, err = c.Flags().GetString("namespace")
		fmt.Println("promptCmd  get namespace is:", ns)
		if err != nil {
			log.Fatalln(err)
		}

		if len(args) > 1{
			ns = args[1]
		}

		if ns == ""{
			ns = "default"
		}

		InitCache()
		p := prompt.New(
			executorCmd(c,&ns),
			completerCmd(&ns),
			prompt.OptionPrefix(">>>"),
			// she zhi  "clear" ming ling lai qing ping
			prompt.OptionWriter(MyConsoleWriter),

		)
		p.Run()
		return nil
	},
}