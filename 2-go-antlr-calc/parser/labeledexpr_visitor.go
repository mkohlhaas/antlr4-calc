// Code generated from parser/LabeledExpr.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // LabeledExpr

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by LabeledExprParser.
type LabeledExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LabeledExprParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by LabeledExprParser#printExpr.
	VisitPrintExpr(ctx *PrintExprContext) interface{}

	// Visit a parse tree produced by LabeledExprParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by LabeledExprParser#blank.
	VisitBlank(ctx *BlankContext) interface{}

	// Visit a parse tree produced by LabeledExprParser#parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by LabeledExprParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by LabeledExprParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by LabeledExprParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by LabeledExprParser#int.
	VisitInt(ctx *IntContext) interface{}
}
