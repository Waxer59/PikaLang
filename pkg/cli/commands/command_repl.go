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

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func SetUpRepl(app *cli.App) *cli.Command {
	replCommand := &cli.Command{
		Name:   "repl",
		Usage:  "Start the repl",
		Action: startRepl,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "ast",
				Aliases: []string{"a"},
				Usage:   "Print the AST",
			},
		},
	}

	return replCommand
}

func startRepl(cCtx *cli.Context) error {
	parser := parser.New()
	env := interpreter_env.New(nil)

	isAstActivated := cCtx.Bool("ast")

	for {
		c := color.New(color.FgBlue).Add(color.Bold)
		c.Print("Pika > ")
		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()
		code := scanner.Text()

		switch code {
		case "exit":
			return cli.Exit("\nGoodbye! :)", int(exitCodes.Success))
		case "clear", "cls":
			utils.CallClearConsoleSc()
			continue
		}

		program, err := parser.ProduceAST(code)

		if isAstActivated {
			fmt.Println(program)
		}

		if err != nil {
			return fmt.Errorf(err.Error())
		}

		eval, err := interpreter_eval.Evaluate(*program, env)

		if err != nil {
			color.Red(err.Error())
		} else {
			fmt.Println(eval.GetValue())
		}

	}
}
