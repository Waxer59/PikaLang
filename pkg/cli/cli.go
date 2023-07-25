package cli

import (
	"os"
	"pika/pkg/cli/commands"

	"github.com/urfave/cli/v2"
)

func New() error {
	app := cli.NewApp()
	app.Name = "pika"
	app.Usage = "A simple pika compiler"
	app.Version = "0.3.1"

	setUp(app)

	err := app.Run(os.Args)

	return err
}

func setUp(app *cli.App) {
	app.Commands = []*cli.Command{
		commands.SetUpRunCommand(app),
		commands.SetUpHelpCommand(app),
		commands.SetUpRepl(app),
	}
}
