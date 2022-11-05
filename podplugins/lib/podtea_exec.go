package lib

import (
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"io"
	v1 "k8s.io/api/core/v1"
	//"k8s.io/apimachinery/pkg/apis/meta/internalversion/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"os"
)

func execPod(ns, pod, container string) remotecommand.Executor {
	option := &v1.PodExecOptions{
		Container: container,
		Command: []string{"sh"},
		Stdin: true,
		Stdout: true,
		Stderr: true,
		TTY: true,
	}
	req := client.CoreV1().RESTClient().Post().Resource("pods").
		Namespace(ns).Name(pod).SubResource("exec").
		Param("color","false").
		VersionedParams(
			option,
			scheme.ParameterCodec,
			)
	exec, err := remotecommand.NewSPDYExecutor(restConfig, "POST",
		req.URL())
	if err != nil {
		panic(err)
	}
	return exec
}

type execModel struct {
	items []v1.Container
	index int
	cmd *cobra.Command
	podName string
	ns string
}



func (m *execModel) Init() tea.Cmd {
	// gen ju PodName qu chu container lie biao
	//m.ns = getNameSpace(m.cmd)
	m.ns = ns
	fmt.Println("get namespace is : ",m.ns)
	pod,err := client.CoreV1().Pods(m.ns).Get(context.Background(),m.podName,metav1.GetOptions{})
	if err != nil {
		log.Println(err)
		return tea.Quit

	}
	m.items=pod.Spec.Containers
	return nil
}

func (m *execModel)Update(msg tea.Msg) (tea.Model,tea.Cmd)  {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m,tea.Quit
		case "up":
			if m.index > 0{
				m.index--
			}
		case "down":
			if m.index < len(m.items)-1 {
				m.index++
			}
		case "enter":
			err := execPod(m.ns,m.podName,m.items[m.index].Name).
				Stream(remotecommand.StreamOptions{
					Stdin: os.Stdin,
					Stdout: os.Stdout,
					Stderr: os.Stderr,
					//Tty: false,
					Tty: true,
			})

			if err != nil {
				if err != io.EOF {
					log.Println("aaaaaaaa",err)
				}
			}
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *execModel) View() string {
	s := "an shang xia jian xuan ze rong qi\n"

	for i ,item := range m.items{
		selected := ""
		if m.index == i{
			selected = ">>"
		}
		//s += fmt.Sprintf("#{selected} #{item.Name}(jingxiang:#{item.Image})\n")
		s += fmt.Sprintf("%s %s (IMAGE: %s)\n",selected,item.Name,item.Image)
	}
	s += "\n enter Q logout"
	return s
}

func runteaExec(args []string, cmd *cobra.Command) {
	if len(args) == 0{
		log.Println("podname is required")
		return
	}
	var execmodel = &execModel{
		cmd: cmd,
		podName: args[0],
	}

	teaCmd := tea.NewProgram(execmodel)
	if err := teaCmd.Start();err !=nil{
		fmt.Println("Start failed:", err)
		os.Exit(1)
	}
}
