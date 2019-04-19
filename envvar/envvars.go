package main

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
	"io"
	"log"
	"os"
)

func main()  {
	app := cli.NewApp()
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "run, r",
			Value: "alpine",
			Usage: "run alpine",
		},
	}

	app.Authors = []cli.Author {
		{
			Name:  "Hai Dam",
			Email: "haidv@tomochain.com",
		},
	}
	app.Action = func(c * cli.Context) error {
		ctx := context.Background()
		_cli, err := client.NewEnvClient()
		if err != nil {
			panic(err)
		}

		reader, err := _cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
		if err != nil {
			panic(err)
		}
		_, _ = io.Copy(os.Stdout, reader)

		resp, err := _cli.ContainerCreate(ctx, &container.Config{
			Image: "alpine",
			Cmd:   []string{"echo", "hello world"},
			Tty:   true,
		}, nil, nil, "")
		if err != nil {
			panic(err)
		}

		if err := _cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
			panic(err)
		}

		statusCh, errCh := _cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
		select {
		case err := <-errCh:
			if err != nil {
				panic(err)
			}
		case <-statusCh:
		}

		out, err := _cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
		if err != nil {
			panic(err)
		}

		io.Copy(os.Stdout, out)
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
