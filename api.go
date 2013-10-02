package main

type Job struct {
	shortname   string
	description string

	// repo paths
	local, remote string

	deadline string
}

type Api struct {
}

func (a *Api) Ls() []Job {
	return []Job{
		Job{"clitool", "make a cli tool for hackerlist", "/home/mischief/code/go/src/github.com/hackerlist/hl", "github.com/hackerlist/hl", "Oct 4 2013"},
	}
}

var TheApi Api

func init() {
	TheApi = Api{}
}
