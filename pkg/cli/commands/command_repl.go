package commands

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Waxer59/PikaLang/internal/utils"
	"github.com/Waxer59/PikaLang/pkg/cli/exitCodes"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_env"
	"github.com/Waxer59/PikaLang/pkg/interpreter/interpreter_eval"
	"github.com/Waxer59/PikaLang/pkg/parser"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func SetUpRepl() *cli.Command {
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
	p := parser.New()
	env := interpreter_env.New(nil)

	isAstActivated := cCtx.Bool("ast")

	for {
		c := color.New(color.FgBlue).Add(color.Bold)
		_, err := c.Print("Pika > ")

		if err != nil {
			return err
		}

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

		program, err := p.ProduceAST(code)

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
