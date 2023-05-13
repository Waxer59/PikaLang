package commands

import (
	"fmt"
	"pika/internal/utils"
	"pika/pkg/cli/exitCodes"
	"pika/pkg/parser"

	"github.com/urfave/cli/v2"
)

const DEFAULT_FILE_NAME = "main.pika"

func SetUpRunCommand(app *cli.App) *cli.Command {
	runCommand := cli.Command{
		Name:   "run",
		Usage:  "Run a file",
		Action: runApp,
	}

	return &runCommand
}

func runApp(cCtx *cli.Context) error {
	fileName := cCtx.Args().Get(0)

	if fileName == "" {
		return cli.Exit("File name is required", int(exitCodes.FileNameError))
	}

	if fileName == "." {
		fileName = DEFAULT_FILE_NAME
	}

	fileLines, err := utils.ScanFile(fileName)

	if err != nil {
		return cli.Exit(err.Error(), int(exitCodes.FileReadError))
	}

	for _, line := range fileLines {
		parser := parser.New()
		fmt.Println(parser.ProduceAST(line), line)
	}

	return nil
}
