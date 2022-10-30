//package main
//
//import (
//	"fmt"
//	"github.com/c-bata/go-prompt"
//	"os"
//	"strings"
//)
//
//func executor(in string) {
//	in = strings.TrimSpace(in)
//
//	blocks := strings.Split(in, " ")
//	switch blocks[0] {
//	case "exit":
//		fmt.Println("Bye!")
//		os.Exit(0)
//
//	}
//}
//var suggestions = []prompt.Suggest{
//	{"test","this is test"},
//	{"exit","Exit http-prompt"},
//}
//
//func completer(in prompt.Document) []prompt.Suggest {
//	w := in.GetWordBeforeCursor()
//	if w == ""{
//		return []prompt.Suggest{}
//	}
//	return prompt.FilterHasPrefix(suggestions,w,true)
//}
//
//func main() {
//	p := prompt.New(
//		executor,
//		completer,
//		prompt.OptionPrefix(">>>"),
//
//		)
//	p.Run()
//}


package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"regexp"
	"strings"
)

var suggestions = []prompt.Suggest{
	{"test","this is test"},
	{"exit","Exit http-prompt"},
	{"list","Get Pods"},
}

var podSuggestions = []prompt.Suggest{
	{"test","this is test"},
	{"exit","Exit http-prompt"},
	{"list","Get Pods"},
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	cmd, opt := parseCmd(d.TextBeforeCursor())
	if  cmd == "get"{
		return prompt.FilterHasPrefix(podSuggestions,opt,true)

	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func parseCmd(w string) (string, string) {
	w = regexp.MustCompile("\\s+").ReplaceAllString(w, "")
	l := strings.Split(w,"")
	if len(l)>= 2{
		return l[0],strings.Join(l[1:],"")
	}
	return w,""
}

func main() {
	fmt.Println("Please select table.")
	t := prompt.Input("> ", completer)
	fmt.Println("You selected " + t)
}