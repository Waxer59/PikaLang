package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"pika/internal/utils"
	"pika/pkg/cli/exitCodes"
	"pika/pkg/interpreter/interpreter_env"
	"pika/pkg/interpreter/interpreter_eval"
	"pika/pkg/parser"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

const DEFAULT_FILE_NAME = "main.pk"

func SetUpRunCommand(app *cli.App) *cli.Command {
	runCommand := cli.Command{
		Name:   "run",
		Usage:  "Run a file",
		Action: runApp,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "ast",
				Aliases: []string{"a"},
				Usage:   "Print the AST",
			},
		},
	}

	return &runCommand
}

func runApp(cCtx *cli.Context) error {
	fileName := cCtx.Args().Get(0)
	ext := filepath.Ext(fileName)

	isAstActivated := cCtx.Bool("ast")

	if ext != ".pk" && !strings.HasSuffix(fileName, "/") && fileName != "." {
		return cli.Exit("File extension must be .pk", int(exitCodes.FileExtensionError))
	}

	env := interpreter_env.New(nil)
	wd, err := os.Getwd()

	if err != nil {
		fmt.Println("Error:", err)
		return cli.Exit("Error getting working directory", int(exitCodes.GetWDError))
	}

	if fileName == "" {
		return cli.Exit("File name is required", int(exitCodes.FileNameError))
	}

	if fileName == "." || strings.HasSuffix(fileName, "/") {
		fileName = filepath.Join(wd, DEFAULT_FILE_NAME)
	} else {
		fileName = filepath.Join(wd, fileName)
	}

	src, err := utils.ScanFile(fileName)

	if err != nil {
		return cli.Exit(err.Error(), int(exitCodes.FileReadError))
	}

	parser := parser.New()

	program, err := parser.ProduceAST(src)

	if isAstActivated {
		fmt.Println(program)
	}

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	_, err = interpreter_eval.Evaluate(*program, env)

	if err != nil {
		color.Red(err.Error())
	}

	return nil
}
