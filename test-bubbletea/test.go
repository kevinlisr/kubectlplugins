package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	items []string
	index int
}

func (m model)Init() tea.Cmd {
	return nil
}
func (m model)Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			fmt.Println(m.items[m.index])
			return m, tea.Quit
		}


	}
	return m, nil
}
func (m model)View() string {
	s := "welcome to K8S Visualization system!\n"
	for i, item := range m.items{
		selected := " "
		if m.index == i{
			selected = ">>"
		}
		s += fmt.Sprintf("%s %s\n", selected, item)
	}
	s += "\nEnter Q logout\n"
	return s
}

func main() {
	var initModel = model{
		items: []string{"I can see Pods","list Deployments","I can see configmaps"},
	}
	cmd:=tea.NewProgram(initModel)
	if err := cmd.Start();err != nil{
		fmt.Println("start failed:", err)
		os.Exit(1)
	}


}