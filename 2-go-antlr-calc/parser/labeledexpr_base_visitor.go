// Code generated from parser/LabeledExpr.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // LabeledExpr

import "github.com/antlr4-go/antlr/v4"

type BaseLabeledExprVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLabeledExprVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabeledExprVisitor) VisitPrintExpr(ctx *PrintExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabeledExprVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabeledExprVisitor) VisitBlank(ctx *BlankContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabeledExprVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabeledExprVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabeledExprVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabeledExprVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabeledExprVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}
