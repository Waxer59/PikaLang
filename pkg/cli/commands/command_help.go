package commands

import (
	"pika/pkg/cli/exitCodes"

	"github.com/urfave/cli/v2"
)

func SetUpHelpCommand(app *cli.App) *cli.Command {
	helpCommand := &cli.Command{
		Name:  "help",
		Usage: "Show help",
		Action: func(cCtx *cli.Context) error {
			cli.ShowAppHelpAndExit(cCtx, int(exitCodes.Success))
			return nil
		},
	}

	return helpCommand
}
