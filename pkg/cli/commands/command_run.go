package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Waxer59/PikaLang/internal/utils"
	"github.com/Waxer59/PikaLang/pkg/cli/exitCodes"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_env"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_eval"
	"github.com/Waxer59/PikaLang/pkg/parser"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

const DEFAULT_FILE_NAME = "main.pk"

func SetUpRunCommand() *cli.Command {
	runCommand := cli.Command{
		Name:   "run",
		Usage:  "Run a file",
		Action: runApp,
	}

	return &runCommand
}

func runApp(cCtx *cli.Context) error {
	fileName := cCtx.Args().Get(0)
	ext := filepath.Ext(fileName)

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

	p := parser.New()

	program, err := p.ProduceAST(src)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	_, err = interpreter_eval.Evaluate(*program, env)

	if err != nil {
		color.Red(err.Error())
	}

	return nil
}
