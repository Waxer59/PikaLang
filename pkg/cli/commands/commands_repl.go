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
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()
		code := scanner.Text()

		if code == "exit" {
			return cli.Exit("exit", int(exitCodes.Success))
		}

		parser := parser.New()
		program := parser.ProduceAST(code)

		result := interpreter.Evaluate(program)

		fmt.Println(result)
	}
}
