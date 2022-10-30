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


func executorCmd(cmd *cobra.Command) func(in string) {
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
			if err := cacheCmd.RunE(cmd, args);err !=nil{
				log.Fatalln(err)
			}
		case "get":
			getPodDetail(args)
	  }

	}
}
var suggestions = []prompt.Suggest{
	{"get","GET "},
	{"list","LIST"},
	{"exit","EXIT the interactive window"},
}

var podSuggestions = []prompt.Suggest{
	{"get","get pods details"},
	{"list","show pods list"},
	{"exit","exit the interactive window"},
}

func getPodsList()(ret []prompt.Suggest){
	pods , err :=fact.Core().V1().Pods().Lister().
		Pods("default").List(labels.Everything())
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

func completer(in prompt.Document) []prompt.Suggest {
	w := in.GetWordBeforeCursor()
	if w == ""{
		return []prompt.Suggest{}
	}

	cmd, opt := parseCmd(in.TextBeforeCursor())
	//fmt.Println(cmd)
	if cmd == "get"{
		return prompt.FilterHasPrefix(getPodsList(),opt,true)
	}
	return prompt.FilterHasPrefix(suggestions,w,true)
}


var promptCmd = &cobra.Command{
	Use: "prompt",
	Short: "prompt pods",
	Example: "kubectl pods prompt [flags]",
	SilenceUsage: true,
	RunE: func(c *cobra.Command, args []string) error {
		InitCache()
		p := prompt.New(
			executorCmd(c),
			completer,
			prompt.OptionPrefix(">>>"),

		)
		p.Run()
		return nil
	},
}