package main

import (
	"github.com/naviat/practice_cli/docker-sdk/container"
)

func main()  {
	container.CreateNewContainer("nginx:latest")
}