package lib

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type podjson struct {
	title string
	path string
}

type podmodel struct {
	items []*podjson
	index int
	cmd *cobra.Command
	podName string
	ns string
}

func (m podmodel)Init() tea.Cmd {
	return nil
}
func (m podmodel)Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.index > 0{
				m.index--
			}
		case "down":
			if m.index < len(m.items)-1{
				m.index++
			}
		case "enter":
			getPodDetailByJSON(m.podName, m.items[m.index].path,m.ns,m.cmd)
			return m, tea.Quit
		}


	}
	return m, nil
}
func (m podmodel)View() string {
	s := "welcome to K8S Visualization system!\n"
	for i, item := range m.items{
		selected := " "
		if m.index == i{
			selected = ">>"
		}
		s += fmt.Sprintf("%s %s\n", selected, item.title)
	}
	s += "\nEnter Q logout\n"
	return s
}

const (
	PodEventType = "__event__"
	PodLogType = "__log__"
)

func runtea(args []string, cmd *cobra.Command,ns string) {
	if len(args) == 0{
		log.Println("pod name is required!")
		return
	}
	var podModel = podmodel{
		items: []*podjson{},
		cmd: cmd,
		podName: args[0],
		ns: ns,
	}
	//v1.Pod{}
	podModel.items = append(podModel.items,
		&podjson{title: "Meta Info",path: "metadata"},
		&podjson{title: "Labels",path: "metadata.labels"},
		&podjson{title: "Annotations",path: "metadata.annotations"},
		&podjson{title: "Containers",path: "spec.containers"},
		&podjson{title: "All Info", path: "@this"},
		&podjson{title: "*Events*", path: PodEventType},
		&podjson{title: "*Logs*", path: PodLogType},
		)
	teaCmd := tea.NewProgram(podModel)
	if err := teaCmd.Start();err != nil{
		fmt.Println("Start failed:", err)
		os.Exit(1)
	}
}

