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
			Name: "lang, l",
			Value: "english",
			Usage: "language for the greeting",
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
