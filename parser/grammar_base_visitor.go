// Generated from Grammar.g4 by ANTLR 4.7.

package parser // Grammar

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGrammarVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitHandler(ctx *HandlerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitTypeString(ctx *TypeStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitTypeBool(ctx *TypeBoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitTypeInt(ctx *TypeIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitTypeFloat(ctx *TypeFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitOperator(ctx *OperatorContext) interface{} {
	return v.VisitChildren(ctx)
}
