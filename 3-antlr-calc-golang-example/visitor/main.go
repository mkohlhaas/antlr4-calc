// for vistor pattern
package main

import (
	"fmt"
	"strconv"

	"calc/parser"

	"github.com/antlr4-go/antlr/v4"
	prompt "github.com/c-bata/go-prompt"
)

type Visitor struct {
	parser.BaseCalcVisitor
	stack []int
}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func (l *Visitor) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *Visitor) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (v *Visitor) visit(node antlr.RuleNode) any {
	node.Accept(v)
	return nil
}

func (v *Visitor) VisitStart(ctx *parser.StartContext) any {
	return v.visit(ctx.Expression())
}

func (v *Visitor) VisitNumber(ctx *parser.NumberContext) any {
	i, err := strconv.Atoi(ctx.NUMBER().GetText())
	if err != nil {
		panic(err.Error())
	}

	v.push(i)
	return nil
}

func (v *Visitor) VisitMulDiv(ctx *parser.MulDivContext) any {
	//push expression result to stack
	v.visit(ctx.Expression(0))
	v.visit(ctx.Expression(1))
	//push result to stack
	var t antlr.Token = ctx.GetOp()
	right := v.pop()
	left := v.pop()
	switch t.GetTokenType() {
	case parser.CalcParserMUL:
		v.push(left * right)
	case parser.CalcParserDIV:
		v.push(left / right)
	default:
		panic("should not happen")
	}
	return nil
}

func (v *Visitor) VisitAddSub(ctx *parser.AddSubContext) any {
	//push expression result to stack
	v.visit(ctx.Expression(0))
	v.visit(ctx.Expression(1))

	//push result to stack
	var t antlr.Token = ctx.GetOp()
	right := v.pop()
	left := v.pop()
	switch t.GetTokenType() {
	case parser.CalcParserADD:
		v.push(left + right)
	case parser.CalcParserSUB:
		v.push(left - right)
	default:
		panic("should not happen")
	}

	return nil
}

func (v *Visitor) VisitParenthesis(ctx *parser.ParenthesisContext) any {
	v.visit(ctx.Expression())
	return nil
}

func calc(input string) int {

	inputStream := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewCalcLexer(inputStream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(tokens)

	v := NewVisitor()
	p.Start_().Accept(v)
	return v.pop()
}

func executor(in string) {
	fmt.Printf("Answer: %d\n", calc(in))
}

func completer(in prompt.Document) []prompt.Suggest {
	var ret []prompt.Suggest
	return ret
}

func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionTitle("calc"),
	)
	p.Run()
}
