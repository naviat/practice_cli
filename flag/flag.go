package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main()  {
	app := cli.NewApp()
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}
	app.Authors = []cli.Author {
		{
			Name:  "Hai Dam",
			Email: "haidv@tomochain.com",
		},
	}
	app.Action = func(c * cli.Context) error {
		name := "Hai"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		if c.String("lang") == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
