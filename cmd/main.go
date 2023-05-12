package main

import (
	"bufio"
	"fmt"
	"os"
	"pikalang/pkg/parser"
)

func main() {
	file, err := os.Open("../test.txt")
	if err != nil {
		panic("There was an error opening the file: " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// tokens := lexer.Tokenize(line)
		parser := parser.Parser{}
		fmt.Println(parser.ProduceAST(line), line)
	}

	if err := scanner.Err(); err != nil {
		panic("There was an error reading the file: " + err.Error())
	}
}
