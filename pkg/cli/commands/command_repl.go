package commands

import (
	"bufio"
	"fmt"
	"os"
	"pika/pkg/cli/exitCodes"
	"pika/pkg/interpreter/interpreter_env"
	"pika/pkg/interpreter/interpreter_eval"
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
	env := interpreter_env.New(nil)

	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()
		code := scanner.Text()

		if code == "exit" {
			return cli.Exit("Goodbye! :)", int(exitCodes.Success))
		}

		program := parser.ProduceAST(code)

		fmt.Println("AST: ", program)

		result := interpreter_eval.Evaluate(program, env)

		fmt.Println(result)
	}
}
