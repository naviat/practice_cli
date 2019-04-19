package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main()  {
	app := cli.NewApp()
	app.Flags =[]cli.Flag {
		cli.StringFlag {
			Name: "password, p",
			Usage: "password for the mysql database",
			FilePath: "/etc/mysql/password",
		},
	}

	err := cli.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
