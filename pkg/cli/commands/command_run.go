package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"pika/internal/utils"
	"pika/pkg/cli/exitCodes"
	"pika/pkg/interpreter"
	"pika/pkg/interpreter/interpreterEnvironment"
	"pika/pkg/parser"

	"github.com/urfave/cli/v2"
)

const DEFAULT_FILE_NAME = "main.pk"

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
	env := interpreterEnvironment.New(nil)
	wd, err := os.Getwd()

	if err != nil {
		fmt.Println("Error:", err)
		return cli.Exit("Error getting working directory", int(exitCodes.GetWDError))
	}

	if fileName == "" {
		return cli.Exit("File name is required", int(exitCodes.FileNameError))
	}

	if fileName == "." {
		fileName = filepath.Join(wd, DEFAULT_FILE_NAME)
	} else {
		fileName = filepath.Join(wd, fileName)
	}

	src, err := utils.ScanFile(fileName)

	if err != nil {
		return cli.Exit(err.Error(), int(exitCodes.FileReadError))
	}

	parser := parser.New()

	fmt.Println(parser.ProduceAST(src))

	program := parser.ProduceAST(src)

	result := interpreter.Evaluate(program, env)

	fmt.Println(result)

	return nil
}
