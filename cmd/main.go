package main

import (
	"os"
	"pika/pkg/cli"

	"github.com/fatih/color"
)

func main() {
	err := cli.New()

	if err != nil {
		color.Red(err.Error())
		os.Exit(0)
	}
}
