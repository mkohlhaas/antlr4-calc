// Code generated from parser/Calc.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Calc

import "github.com/antlr4-go/antlr/v4"

type BaseCalcVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCalcVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcVisitor) VisitParenthesis(ctx *ParenthesisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}
