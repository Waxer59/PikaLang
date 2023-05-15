package commands

import (
	"bufio"
	"fmt"
	"os"
	"pika/internal/utils"
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

		switch code {
		case "exit":
			return cli.Exit("Goodbye! :)", int(exitCodes.Success))
		case "clear", "cls":
			utils.CallClearConsoleSc()
		}

		program := parser.ProduceAST(code)
		result, _ := interpreter_eval.Evaluate(program, env)

		fmt.Println(result)
	}
}
