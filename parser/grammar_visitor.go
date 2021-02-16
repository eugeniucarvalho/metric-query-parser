// Generated from Grammar.g4 by ANTLR 4.7.

package parser // Grammar

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by GrammarParser.
type GrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GrammarParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by GrammarParser#handler.
	VisitHandler(ctx *HandlerContext) interface{}

	// Visit a parse tree produced by GrammarParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by GrammarParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by GrammarParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by GrammarParser#typeString.
	VisitTypeString(ctx *TypeStringContext) interface{}

	// Visit a parse tree produced by GrammarParser#typeBool.
	VisitTypeBool(ctx *TypeBoolContext) interface{}

	// Visit a parse tree produced by GrammarParser#typeInt.
	VisitTypeInt(ctx *TypeIntContext) interface{}

	// Visit a parse tree produced by GrammarParser#typeFloat.
	VisitTypeFloat(ctx *TypeFloatContext) interface{}

	// Visit a parse tree produced by GrammarParser#operator.
	VisitOperator(ctx *OperatorContext) interface{}
}
