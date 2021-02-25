// Generated from Grammar.g4 by ANTLR 4.7.

package parser // Grammar

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGrammarListener is a complete listener for a parse tree produced by GrammarParser.
type BaseGrammarListener struct{}

var _ GrammarListener = &BaseGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGrammarListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGrammarListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHandler is called when production handler is entered.
func (s *BaseGrammarListener) EnterHandler(ctx *HandlerContext) {}

// ExitHandler is called when production handler is exited.
func (s *BaseGrammarListener) ExitHandler(ctx *HandlerContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseGrammarListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseGrammarListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseGrammarListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseGrammarListener) ExitArgument(ctx *ArgumentContext) {}

// EnterValue is called when production value is entered.
func (s *BaseGrammarListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseGrammarListener) ExitValue(ctx *ValueContext) {}

// EnterTypeHandler is called when production typeHandler is entered.
func (s *BaseGrammarListener) EnterTypeHandler(ctx *TypeHandlerContext) {}

// ExitTypeHandler is called when production typeHandler is exited.
func (s *BaseGrammarListener) ExitTypeHandler(ctx *TypeHandlerContext) {}

// EnterTypeString is called when production typeString is entered.
func (s *BaseGrammarListener) EnterTypeString(ctx *TypeStringContext) {}

// ExitTypeString is called when production typeString is exited.
func (s *BaseGrammarListener) ExitTypeString(ctx *TypeStringContext) {}

// EnterTypeBool is called when production typeBool is entered.
func (s *BaseGrammarListener) EnterTypeBool(ctx *TypeBoolContext) {}

// ExitTypeBool is called when production typeBool is exited.
func (s *BaseGrammarListener) ExitTypeBool(ctx *TypeBoolContext) {}

// EnterTypeInt is called when production typeInt is entered.
func (s *BaseGrammarListener) EnterTypeInt(ctx *TypeIntContext) {}

// ExitTypeInt is called when production typeInt is exited.
func (s *BaseGrammarListener) ExitTypeInt(ctx *TypeIntContext) {}

// EnterTypeFloat is called when production typeFloat is entered.
func (s *BaseGrammarListener) EnterTypeFloat(ctx *TypeFloatContext) {}

// ExitTypeFloat is called when production typeFloat is exited.
func (s *BaseGrammarListener) ExitTypeFloat(ctx *TypeFloatContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseGrammarListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseGrammarListener) ExitOperator(ctx *OperatorContext) {}
