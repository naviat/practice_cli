package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main()  {
	app := cli.NewApp()

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "config, c",
			Usage: "Load configuration from `File`",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
