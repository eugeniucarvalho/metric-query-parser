// Generated from Grammar.g4 by ANTLR 4.7.

package parser // Grammar

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GrammarListener is a complete listener for a parse tree produced by GrammarParser.
type GrammarListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHandler is called when entering the handler production.
	EnterHandler(c *HandlerContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterTypeString is called when entering the typeString production.
	EnterTypeString(c *TypeStringContext)

	// EnterTypeBool is called when entering the typeBool production.
	EnterTypeBool(c *TypeBoolContext)

	// EnterTypeInt is called when entering the typeInt production.
	EnterTypeInt(c *TypeIntContext)

	// EnterTypeFloat is called when entering the typeFloat production.
	EnterTypeFloat(c *TypeFloatContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHandler is called when exiting the handler production.
	ExitHandler(c *HandlerContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitTypeString is called when exiting the typeString production.
	ExitTypeString(c *TypeStringContext)

	// ExitTypeBool is called when exiting the typeBool production.
	ExitTypeBool(c *TypeBoolContext)

	// ExitTypeInt is called when exiting the typeInt production.
	ExitTypeInt(c *TypeIntContext)

	// ExitTypeFloat is called when exiting the typeFloat production.
	ExitTypeFloat(c *TypeFloatContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)
}
