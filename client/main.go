package main

import (
	"os"
	"sort"

	"github.com/codegangsta/cli"
)

var (
	gitversion string
)

func main() {
	app := cli.NewApp()
	app.Name = "snappm"
	app.Version = gitversion
	app.Usage = "A powerful snap plugin manager"
	//app.Flags = []cli.Flag{flURL, flSecure, flAPIVer, flPassword, flConfig}
	app.Commands = append(commands)
	sort.Sort(ByCommand(app.Commands))
	//app.Before = beforeAction
	app.Run(os.Args)
}
