package main

import (
	"fmt"
	"github.com/hackerlist/hljson"
	"log"
	"strings"
)

// Interface of all commands
type Cmd interface {
	Name() string
	Help() string
	Run(argv []string) string
}

type RunFn func(argv []string) string

type Command struct {
	name, help string
	fn         RunFn
}

func (c *Command) Name() string {
	return c.name
}

func (c *Command) Help() string {
	return c.help
}

func (c *Command) Run(argv []string) string {
	return c.fn(argv)
}

var Commands map[string]Cmd

func init() {
	Commands = make(map[string]Cmd)

	Commands["ls"] = &Command{
		"ls",
		"list tasks",
		func(argv []string) string {
			log.Printf("args: %+v", argv)
			return fmt.Sprintf("%+v", TheApi.Ls())
		},
	}

	Commands["missions"] = &Command{
		"missions",
		"list missions",
		func(argv []string) string {
			missions, err := hljson.GetMissions()
			if err != nil {
				log.Printf("Error retrieving missions: %s", err)
			}

			twidth := TerminalWidth()

			fmt.Printf("%s\n", strings.Repeat("-", twidth))

			for _, m := range missions {
				fmt.Printf("Id: %-5d Openings: %-5d Budget: %-10.0f\n", m.Id, m.Openings, m.BudgetEst)
				fmt.Printf("Title: %s\n", m.Title)
				desclines := wrap(m.Description, twidth)
				fmt.Print("Description:\n")
				for _, l := range desclines {
					fmt.Println(l)
				}
				fmt.Printf("%s\n", strings.Repeat("-", twidth))
			}

			return ""
		},
	}
}
