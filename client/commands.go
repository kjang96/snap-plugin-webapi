package main

import "github.com/codegangsta/cli"

var (
	commands = []cli.Command{
		{
			Name:   "list",
			Usage:  "Display the info of a plugin",
			Action: infoByName,
		},
	}
)

// ByCommand contains array of CLI commands.
type ByCommand []cli.Command

func (s ByCommand) Len() int {
	return len(s)
}
func (s ByCommand) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByCommand) Less(i, j int) bool {
	if s[i].Name == "help" {
		return false
	}
	if s[j].Name == "help" {
		return true
	}
	return s[i].Name < s[j].Name
}
