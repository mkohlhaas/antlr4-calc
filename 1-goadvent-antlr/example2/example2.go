package main

import (
	"calc/parser"

	"github.com/antlr4-go/antlr/v4"
)

type calcListener struct {
	*parser.BaseCalcListener
}

func main() {
	// Setup the input
	input := antlr.NewInputStream("1 + 2 * 3")

	// Create the Lexer
	lexer := parser.NewCalcLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	antlr.ParseTreeWalkerDefault.Walk(&calcListener{}, p.Start_())
}
