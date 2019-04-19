package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"sort"
)

func main()  {
	app := cli.NewApp()
	cli.AppHelpTemplate = fmt.Sprintf(`%s
SUPPORT: haidv@tomochain.com
`,cli.AppHelpTemplate)
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "lang, l",
			Value: "english",
			Usage: "Language for the greeting",
		},
		cli.StringFlag{
			Name: "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	app.Commands = []cli.Command {
		{
			Name: "complete",
			Aliases: []string{"c",},
			Usage: "complete a task on the list",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name: "add",
			Aliases: []string{"a"},
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}