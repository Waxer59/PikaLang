package cli

import (
	"log"
	"os"

	"github.com/Waxer59/PikaLang/pkg/cli/commands"

	"github.com/urfave/cli/v2"
)

func New() error {
	app := &cli.App{
		Name:    "pika",
		Usage:   "A simple pika compiler",
		Version: "0.5.2",
	}

	setUp(app)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func setUp(app *cli.App) {
	app.Commands = []*cli.Command{
		commands.SetUpRunCommand(app),
		commands.SetUpHelpCommand(app),
		commands.SetUpRepl(app),
	}
}
