package main

import (
	"fmt"
	"log"
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
}
