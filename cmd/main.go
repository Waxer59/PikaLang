package main

import "pika/pkg/cli"

func main() {
	err := cli.New()

	if err != nil {
		panic("Something went wrong")
	}
}
