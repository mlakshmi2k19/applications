package main

import (
	"fmt"
	"os"

	cli "github.com/urfave/cli/v2"
)

var App = cli.App{
	Name:      "load-genie",
	Usage:     "Tool used to load data on demand",
	ArgsUsage: "",
	Commands: []*cli.Command{
		{
			Name:    "load",
			Aliases: []string{"l"},
			Usage:   "Loads data in database",
			Action:  doAction,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "file-path",
					Usage:    "Input csv file path should be configured",
					Required: true,
				},
				&cli.BoolFlag{
					Name:  "create-tables",
					Usage: "If it is enabled table creation will be executed else that part will be skipped",
				},
			},
		},
	},
}

func main() {
	err := App.Run(os.Args)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(2)
	}
}
