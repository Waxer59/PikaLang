package commands

import (
	"bufio"
	"fmt"
	"os"
	"pika/pkg/cli/exitCodes"
	"pika/pkg/interpreter"
	"pika/pkg/parser"

	"github.com/urfave/cli/v2"
)

func SetUpRepl(app *cli.App) *cli.Command {
	replCommand := &cli.Command{
		Name:   "repl",
		Usage:  "Start the repl",
		Action: startRepl,
	}

	return replCommand
}

func startRepl(cCtx *cli.Context) error {
	parser := parser.New()
	env := interpreter.NewEnvironment(nil)

	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()
		code := scanner.Text()

		if code == "exit" {
			return cli.Exit("Goodbye! :)", int(exitCodes.Success))
		}

		program := parser.ProduceAST(code)

		result := interpreter.Evaluate(program, env)

		fmt.Println(result)
	}
}
