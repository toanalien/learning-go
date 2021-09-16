package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/nsf/termbox-go"
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
	"time"
)

var optsId = 0
var arrOpts = []types.ContainerListOptions{
	{All: true},
	{Filters: filters.NewArgs(filters.Arg("status", "running"))},
	{Filters: filters.NewArgs(filters.Arg("status", "exited"),
		filters.Arg("status", "dead"),
		filters.Arg("status", "created"),
		filters.Arg("status", "restarting"),
		filters.Arg("status", "paused"))},
}

func doCmd(ctx context.Context, cli *client.Client, cmd string) {
	fmt.Print("\033[H\033[2J")
	cmds := strings.Split(cmd, " ")
	if len(cmds) < 2 {
		fmt.Println("Cannot stop container")
	} else {
		switch cmds[0] {
		case "stop":
			if err := cli.ContainerStop(ctx, cmds[1], nil); err != nil {
				fmt.Println("Cannot stop container")
			}
		case "start":
			if err := cli.ContainerStart(ctx, cmds[1], types.ContainerStartOptions{}); err != nil {
				fmt.Println("Cannot start container")
			}
		default:
			fmt.Println("Command not found")
		}
	}
	getContainer(ctx, cli, arrOpts[0])
}

func getContainer(ctx context.Context, cli *client.Client, opts types.ContainerListOptions) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"CONTAINER ID", "IMAGE", "COMMAND", "CREATED", "STATUS", "PORTS", "NAMES"})

	containers, err := cli.ContainerList(ctx, opts)
	if err != nil {
		panic(err)
	}
	for _, con := range containers {
		var ports string
		for _, port := range con.Ports {
			ports += fmt.Sprintf("%s:%d->%d/%s\n", port.IP, port.PublicPort, port.PrivatePort, port.Type)
		}
		var names string
		for _, name := range con.Names {
			names += fmt.Sprintf("%s\n", strings.TrimPrefix(name, "/"))
		}

		var status string

		if strings.Contains(con.Status, "Up") {
			status = "Running"
		} else {
			status = "Stopped"
		}

		created := time.Unix(con.Created, 0)
		table.Append([]string{con.ID[:12], con.Image, con.Command, created.Format("2006-01-02 15:04:05 -0700"), status, ports, names})
	}
	table.Render()
	fmt.Println()
}

func main() {

	// relate docker api
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	// relate termbox
	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	getContainer(ctx, cli, arrOpts[optsId])
	redraw_all()
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyArrowLeft, termbox.KeyCtrlB:
				edit_box.MoveCursorOneRuneBackward()
			case termbox.KeyArrowRight, termbox.KeyCtrlF:
				edit_box.MoveCursorOneRuneForward()
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				edit_box.DeleteRuneBackward()
			case termbox.KeyDelete, termbox.KeyCtrlD:
				edit_box.DeleteRuneForward()
			case termbox.KeyTab:
				edit_box.InsertRune('\t')
			case termbox.KeySpace:
				edit_box.InsertRune(' ')
			case termbox.KeyCtrlK:
				edit_box.DeleteTheRestOfTheLine()
			case termbox.KeyHome, termbox.KeyCtrlA:
				edit_box.MoveCursorToBeginningOfTheLine()
			case termbox.KeyEnd, termbox.KeyCtrlE:
				edit_box.MoveCursorToEndOfTheLine()
			case termbox.KeyF5:
				optsId = (optsId + 1) % 3
				fmt.Print("\033[H\033[2J")
				getContainer(ctx, cli, arrOpts[optsId])
			case termbox.KeyEnter:
				cmd := string(edit_box.text)
				edit_box.MoveCursorTo(0)
				edit_box.DeleteTheRestOfTheLine()
				edit_box.MoveCursorToEndOfTheLine()
				doCmd(ctx, cli, cmd)
			default:
				if ev.Ch != 0 {
					edit_box.InsertRune(ev.Ch)
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
		redraw_all()
	}

}
