package expression

type Expression interface {
	Eval() int
}

type Program struct {
	Expressions []Expression
}

type Number struct {
	Num int
}

type VariableDeclaration struct {
	ID    string
	Type  string
	Value int
}

type Variable struct {
	ID string
}

type Multiplication struct {
	Left  Expression
	Right Expression
}

type Addition struct {
	Left  Expression
	Right Expression
}

func (p *Program) addExpression(expr Expression) {
	p.Expressions = append(p.Expressions, expr)
}
