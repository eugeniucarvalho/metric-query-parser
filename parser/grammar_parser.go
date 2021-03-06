// Generated from Grammar.g4 by ANTLR 4.7.

package parser // Grammar

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 29, 84, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 5, 3, 31, 10, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 7, 4, 38, 10, 4, 12, 4, 14, 4, 41, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 5, 5, 48, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 55, 10, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 5, 12, 82, 10, 12, 3, 12, 2, 2, 13, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 2, 3, 3, 2, 3, 4, 2, 90, 2, 24, 3, 2, 2, 2, 4, 27,
	3, 2, 2, 2, 6, 34, 3, 2, 2, 2, 8, 47, 3, 2, 2, 2, 10, 54, 3, 2, 2, 2, 12,
	56, 3, 2, 2, 2, 14, 58, 3, 2, 2, 2, 16, 60, 3, 2, 2, 2, 18, 62, 3, 2, 2,
	2, 20, 64, 3, 2, 2, 2, 22, 81, 3, 2, 2, 2, 24, 25, 5, 4, 3, 2, 25, 26,
	7, 2, 2, 3, 26, 3, 3, 2, 2, 2, 27, 28, 7, 28, 2, 2, 28, 30, 7, 10, 2, 2,
	29, 31, 5, 6, 4, 2, 30, 29, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 32, 3,
	2, 2, 2, 32, 33, 7, 11, 2, 2, 33, 5, 3, 2, 2, 2, 34, 39, 5, 8, 5, 2, 35,
	36, 7, 25, 2, 2, 36, 38, 5, 8, 5, 2, 37, 35, 3, 2, 2, 2, 38, 41, 3, 2,
	2, 2, 39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 7, 3, 2, 2, 2, 41, 39,
	3, 2, 2, 2, 42, 48, 5, 4, 3, 2, 43, 44, 7, 28, 2, 2, 44, 45, 5, 22, 12,
	2, 45, 46, 5, 10, 6, 2, 46, 48, 3, 2, 2, 2, 47, 42, 3, 2, 2, 2, 47, 43,
	3, 2, 2, 2, 48, 9, 3, 2, 2, 2, 49, 55, 5, 14, 8, 2, 50, 55, 5, 16, 9, 2,
	51, 55, 5, 18, 10, 2, 52, 55, 5, 20, 11, 2, 53, 55, 5, 12, 7, 2, 54, 49,
	3, 2, 2, 2, 54, 50, 3, 2, 2, 2, 54, 51, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2,
	54, 53, 3, 2, 2, 2, 55, 11, 3, 2, 2, 2, 56, 57, 5, 4, 3, 2, 57, 13, 3,
	2, 2, 2, 58, 59, 7, 5, 2, 2, 59, 15, 3, 2, 2, 2, 60, 61, 9, 2, 2, 2, 61,
	17, 3, 2, 2, 2, 62, 63, 7, 27, 2, 2, 63, 19, 3, 2, 2, 2, 64, 65, 7, 27,
	2, 2, 65, 66, 7, 24, 2, 2, 66, 67, 7, 27, 2, 2, 67, 21, 3, 2, 2, 2, 68,
	69, 7, 6, 2, 2, 69, 82, 5, 22, 12, 2, 70, 82, 7, 13, 2, 2, 71, 82, 7, 14,
	2, 2, 72, 82, 7, 15, 2, 2, 73, 82, 7, 16, 2, 2, 74, 82, 7, 17, 2, 2, 75,
	82, 7, 18, 2, 2, 76, 82, 7, 19, 2, 2, 77, 82, 7, 20, 2, 2, 78, 82, 7, 21,
	2, 2, 79, 82, 7, 22, 2, 2, 80, 82, 7, 23, 2, 2, 81, 68, 3, 2, 2, 2, 81,
	70, 3, 2, 2, 2, 81, 71, 3, 2, 2, 2, 81, 72, 3, 2, 2, 2, 81, 73, 3, 2, 2,
	2, 81, 74, 3, 2, 2, 2, 81, 75, 3, 2, 2, 2, 81, 76, 3, 2, 2, 2, 81, 77,
	3, 2, 2, 2, 81, 78, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 80, 3, 2, 2, 2,
	82, 23, 3, 2, 2, 2, 7, 30, 39, 47, 54, 81,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'true'", "'false'", "", "'!'", "'''", "'{'", "'}'", "'('", "')'",
	"'|'", "'='", "'~='", "'|='", "'=|'", "'in'", "'>'", "'>='", "'<'", "'<='",
	"'><'", "'<>'", "'.'", "','", "':'",
}
var symbolicNames = []string{
	"", "", "", "TypeString", "R_NEGATE", "R_SIGLE_QUOTA", "R_BRACE_L", "R_BRACE_R",
	"R_BRACKET_L", "R_BRACKET_R", "R_PIPE", "R_EQ", "R_CONTAINS", "R_START_WITH",
	"R_END_WITH", "R_IN", "R_GT", "R_GTE", "R_LT", "R_LTE", "R_BETWEEN", "R_NOT_BETWEEN",
	"R_DOT", "R_COMA", "R_COLON", "T_INTEIRO", "ID", "WS",
}

var ruleNames = []string{
	"expression", "handler", "arguments", "argument", "value", "typeHandler",
	"typeString", "typeBool", "typeInt", "typeFloat", "operator",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GrammarParser struct {
	*antlr.BaseParser
}

func NewGrammarParser(input antlr.TokenStream) *GrammarParser {
	this := new(GrammarParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Grammar.g4"

	return this
}

// GrammarParser tokens.
const (
	GrammarParserEOF           = antlr.TokenEOF
	GrammarParserT__0          = 1
	GrammarParserT__1          = 2
	GrammarParserTypeString    = 3
	GrammarParserR_NEGATE      = 4
	GrammarParserR_SIGLE_QUOTA = 5
	GrammarParserR_BRACE_L     = 6
	GrammarParserR_BRACE_R     = 7
	GrammarParserR_BRACKET_L   = 8
	GrammarParserR_BRACKET_R   = 9
	GrammarParserR_PIPE        = 10
	GrammarParserR_EQ          = 11
	GrammarParserR_CONTAINS    = 12
	GrammarParserR_START_WITH  = 13
	GrammarParserR_END_WITH    = 14
	GrammarParserR_IN          = 15
	GrammarParserR_GT          = 16
	GrammarParserR_GTE         = 17
	GrammarParserR_LT          = 18
	GrammarParserR_LTE         = 19
	GrammarParserR_BETWEEN     = 20
	GrammarParserR_NOT_BETWEEN = 21
	GrammarParserR_DOT         = 22
	GrammarParserR_COMA        = 23
	GrammarParserR_COLON       = 24
	GrammarParserT_INTEIRO     = 25
	GrammarParserID            = 26
	GrammarParserWS            = 27
)

// GrammarParser rules.
const (
	GrammarParserRULE_expression  = 0
	GrammarParserRULE_handler     = 1
	GrammarParserRULE_arguments   = 2
	GrammarParserRULE_argument    = 3
	GrammarParserRULE_value       = 4
	GrammarParserRULE_typeHandler = 5
	GrammarParserRULE_typeString  = 6
	GrammarParserRULE_typeBool    = 7
	GrammarParserRULE_typeInt     = 8
	GrammarParserRULE_typeFloat   = 9
	GrammarParserRULE_operator    = 10
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Handler() IHandlerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHandlerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHandlerContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(GrammarParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GrammarParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GrammarParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Handler()
	}
	{
		p.SetState(23)
		p.Match(GrammarParserEOF)
	}

	return localctx
}

// IHandlerContext is an interface to support dynamic dispatch.
type IHandlerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHandlerContext differentiates from other interfaces.
	IsHandlerContext()
}

type HandlerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHandlerContext() *HandlerContext {
	var p = new(HandlerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_handler
	return p
}

func (*HandlerContext) IsHandlerContext() {}

func NewHandlerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HandlerContext {
	var p = new(HandlerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_handler

	return p
}

func (s *HandlerContext) GetParser() antlr.Parser { return s.parser }

func (s *HandlerContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *HandlerContext) R_BRACKET_L() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_BRACKET_L, 0)
}

func (s *HandlerContext) R_BRACKET_R() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_BRACKET_R, 0)
}

func (s *HandlerContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *HandlerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandlerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HandlerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterHandler(s)
	}
}

func (s *HandlerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitHandler(s)
	}
}

func (p *GrammarParser) Handler() (localctx IHandlerContext) {
	localctx = NewHandlerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GrammarParserRULE_handler)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(GrammarParserID)
	}
	{
		p.SetState(26)
		p.Match(GrammarParserR_BRACKET_L)
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserID {
		{
			p.SetState(27)
			p.Arguments()
		}

	}
	{
		p.SetState(30)
		p.Match(GrammarParserR_BRACKET_R)
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) AllR_COMA() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserR_COMA)
}

func (s *ArgumentsContext) R_COMA(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserR_COMA, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *GrammarParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GrammarParserRULE_arguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Argument()
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GrammarParserR_COMA {
		{
			p.SetState(33)
			p.Match(GrammarParserR_COMA)
		}
		{
			p.SetState(34)
			p.Argument()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Handler() IHandlerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHandlerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHandlerContext)
}

func (s *ArgumentContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *ArgumentContext) Operator() IOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *ArgumentContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *GrammarParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GrammarParserRULE_argument)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.Handler()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(41)
			p.Match(GrammarParserID)
		}
		{
			p.SetState(42)
			p.Operator()
		}
		{
			p.SetState(43)
			p.Value()
		}

	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) TypeString() ITypeStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeStringContext)
}

func (s *ValueContext) TypeBool() ITypeBoolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeBoolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeBoolContext)
}

func (s *ValueContext) TypeInt() ITypeIntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeIntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeIntContext)
}

func (s *ValueContext) TypeFloat() ITypeFloatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFloatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFloatContext)
}

func (s *ValueContext) TypeHandler() ITypeHandlerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeHandlerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeHandlerContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *GrammarParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GrammarParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.TypeString()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.TypeBool()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(49)
			p.TypeInt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(50)
			p.TypeFloat()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(51)
			p.TypeHandler()
		}

	}

	return localctx
}

// ITypeHandlerContext is an interface to support dynamic dispatch.
type ITypeHandlerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeHandlerContext differentiates from other interfaces.
	IsTypeHandlerContext()
}

type TypeHandlerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeHandlerContext() *TypeHandlerContext {
	var p = new(TypeHandlerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_typeHandler
	return p
}

func (*TypeHandlerContext) IsTypeHandlerContext() {}

func NewTypeHandlerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeHandlerContext {
	var p = new(TypeHandlerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_typeHandler

	return p
}

func (s *TypeHandlerContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeHandlerContext) Handler() IHandlerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHandlerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHandlerContext)
}

func (s *TypeHandlerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeHandlerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeHandlerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTypeHandler(s)
	}
}

func (s *TypeHandlerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTypeHandler(s)
	}
}

func (p *GrammarParser) TypeHandler() (localctx ITypeHandlerContext) {
	localctx = NewTypeHandlerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GrammarParserRULE_typeHandler)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Handler()
	}

	return localctx
}

// ITypeStringContext is an interface to support dynamic dispatch.
type ITypeStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeStringContext differentiates from other interfaces.
	IsTypeStringContext()
}

type TypeStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeStringContext() *TypeStringContext {
	var p = new(TypeStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_typeString
	return p
}

func (*TypeStringContext) IsTypeStringContext() {}

func NewTypeStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeStringContext {
	var p = new(TypeStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_typeString

	return p
}

func (s *TypeStringContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeStringContext) TypeString() antlr.TerminalNode {
	return s.GetToken(GrammarParserTypeString, 0)
}

func (s *TypeStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTypeString(s)
	}
}

func (s *TypeStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTypeString(s)
	}
}

func (p *GrammarParser) TypeString() (localctx ITypeStringContext) {
	localctx = NewTypeStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GrammarParserRULE_typeString)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(GrammarParserTypeString)
	}

	return localctx
}

// ITypeBoolContext is an interface to support dynamic dispatch.
type ITypeBoolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeBoolContext differentiates from other interfaces.
	IsTypeBoolContext()
}

type TypeBoolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeBoolContext() *TypeBoolContext {
	var p = new(TypeBoolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_typeBool
	return p
}

func (*TypeBoolContext) IsTypeBoolContext() {}

func NewTypeBoolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBoolContext {
	var p = new(TypeBoolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_typeBool

	return p
}

func (s *TypeBoolContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBoolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTypeBool(s)
	}
}

func (s *TypeBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTypeBool(s)
	}
}

func (p *GrammarParser) TypeBool() (localctx ITypeBoolContext) {
	localctx = NewTypeBoolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GrammarParserRULE_typeBool)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(58)
	_la = p.GetTokenStream().LA(1)

	if !(_la == GrammarParserT__0 || _la == GrammarParserT__1) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ITypeIntContext is an interface to support dynamic dispatch.
type ITypeIntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeIntContext differentiates from other interfaces.
	IsTypeIntContext()
}

type TypeIntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeIntContext() *TypeIntContext {
	var p = new(TypeIntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_typeInt
	return p
}

func (*TypeIntContext) IsTypeIntContext() {}

func NewTypeIntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeIntContext {
	var p = new(TypeIntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_typeInt

	return p
}

func (s *TypeIntContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeIntContext) T_INTEIRO() antlr.TerminalNode {
	return s.GetToken(GrammarParserT_INTEIRO, 0)
}

func (s *TypeIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeIntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTypeInt(s)
	}
}

func (s *TypeIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTypeInt(s)
	}
}

func (p *GrammarParser) TypeInt() (localctx ITypeIntContext) {
	localctx = NewTypeIntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GrammarParserRULE_typeInt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(GrammarParserT_INTEIRO)
	}

	return localctx
}

// ITypeFloatContext is an interface to support dynamic dispatch.
type ITypeFloatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeFloatContext differentiates from other interfaces.
	IsTypeFloatContext()
}

type TypeFloatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeFloatContext() *TypeFloatContext {
	var p = new(TypeFloatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_typeFloat
	return p
}

func (*TypeFloatContext) IsTypeFloatContext() {}

func NewTypeFloatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeFloatContext {
	var p = new(TypeFloatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_typeFloat

	return p
}

func (s *TypeFloatContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeFloatContext) AllT_INTEIRO() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserT_INTEIRO)
}

func (s *TypeFloatContext) T_INTEIRO(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserT_INTEIRO, i)
}

func (s *TypeFloatContext) R_DOT() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_DOT, 0)
}

func (s *TypeFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeFloatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTypeFloat(s)
	}
}

func (s *TypeFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTypeFloat(s)
	}
}

func (p *GrammarParser) TypeFloat() (localctx ITypeFloatContext) {
	localctx = NewTypeFloatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GrammarParserRULE_typeFloat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(GrammarParserT_INTEIRO)
	}
	{
		p.SetState(63)
		p.Match(GrammarParserR_DOT)
	}
	{
		p.SetState(64)
		p.Match(GrammarParserT_INTEIRO)
	}

	return localctx
}

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_operator
	return p
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) R_NEGATE() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_NEGATE, 0)
}

func (s *OperatorContext) Operator() IOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *OperatorContext) R_EQ() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_EQ, 0)
}

func (s *OperatorContext) R_CONTAINS() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_CONTAINS, 0)
}

func (s *OperatorContext) R_START_WITH() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_START_WITH, 0)
}

func (s *OperatorContext) R_END_WITH() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_END_WITH, 0)
}

func (s *OperatorContext) R_IN() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_IN, 0)
}

func (s *OperatorContext) R_GT() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_GT, 0)
}

func (s *OperatorContext) R_GTE() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_GTE, 0)
}

func (s *OperatorContext) R_LT() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_LT, 0)
}

func (s *OperatorContext) R_LTE() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_LTE, 0)
}

func (s *OperatorContext) R_BETWEEN() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_BETWEEN, 0)
}

func (s *OperatorContext) R_NOT_BETWEEN() antlr.TerminalNode {
	return s.GetToken(GrammarParserR_NOT_BETWEEN, 0)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (p *GrammarParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GrammarParserRULE_operator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(79)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserR_NEGATE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Match(GrammarParserR_NEGATE)
		}
		{
			p.SetState(67)
			p.Operator()
		}

	case GrammarParserR_EQ:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.Match(GrammarParserR_EQ)
		}

	case GrammarParserR_CONTAINS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)
			p.Match(GrammarParserR_CONTAINS)
		}

	case GrammarParserR_START_WITH:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(70)
			p.Match(GrammarParserR_START_WITH)
		}

	case GrammarParserR_END_WITH:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(71)
			p.Match(GrammarParserR_END_WITH)
		}

	case GrammarParserR_IN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(72)
			p.Match(GrammarParserR_IN)
		}

	case GrammarParserR_GT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(73)
			p.Match(GrammarParserR_GT)
		}

	case GrammarParserR_GTE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(74)
			p.Match(GrammarParserR_GTE)
		}

	case GrammarParserR_LT:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(75)
			p.Match(GrammarParserR_LT)
		}

	case GrammarParserR_LTE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(76)
			p.Match(GrammarParserR_LTE)
		}

	case GrammarParserR_BETWEEN:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(77)
			p.Match(GrammarParserR_BETWEEN)
		}

	case GrammarParserR_NOT_BETWEEN:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(78)
			p.Match(GrammarParserR_NOT_BETWEEN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
