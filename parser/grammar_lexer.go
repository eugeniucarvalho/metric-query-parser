// Generated from Grammar.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 30, 182,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14,
	3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 6, 25, 128, 10, 25, 13, 25, 14, 25,
	129, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 7,
	29, 141, 10, 29, 12, 29, 14, 29, 144, 11, 29, 3, 30, 3, 30, 3, 30, 3, 30,
	7, 30, 150, 10, 30, 12, 30, 14, 30, 153, 11, 30, 3, 30, 5, 30, 156, 10,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 6, 31, 163, 10, 31, 13, 31, 14,
	31, 164, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 7, 32, 173, 10, 32,
	12, 32, 14, 32, 176, 11, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 4, 151,
	174, 2, 33, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20,
	39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 2, 53, 2, 55, 2, 57,
	27, 59, 28, 61, 29, 63, 30, 3, 2, 6, 3, 2, 50, 59, 5, 2, 67, 92, 97, 97,
	99, 124, 3, 2, 67, 92, 5, 2, 11, 12, 15, 15, 34, 34, 2, 185, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27,
	3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2,
	35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2,
	2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2,
	2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2,
	2, 2, 3, 65, 3, 2, 2, 2, 5, 70, 3, 2, 2, 2, 7, 76, 3, 2, 2, 2, 9, 78, 3,
	2, 2, 2, 11, 80, 3, 2, 2, 2, 13, 82, 3, 2, 2, 2, 15, 84, 3, 2, 2, 2, 17,
	86, 3, 2, 2, 2, 19, 88, 3, 2, 2, 2, 21, 90, 3, 2, 2, 2, 23, 92, 3, 2, 2,
	2, 25, 95, 3, 2, 2, 2, 27, 98, 3, 2, 2, 2, 29, 101, 3, 2, 2, 2, 31, 104,
	3, 2, 2, 2, 33, 106, 3, 2, 2, 2, 35, 109, 3, 2, 2, 2, 37, 111, 3, 2, 2,
	2, 39, 114, 3, 2, 2, 2, 41, 117, 3, 2, 2, 2, 43, 120, 3, 2, 2, 2, 45, 122,
	3, 2, 2, 2, 47, 124, 3, 2, 2, 2, 49, 127, 3, 2, 2, 2, 51, 131, 3, 2, 2,
	2, 53, 133, 3, 2, 2, 2, 55, 135, 3, 2, 2, 2, 57, 137, 3, 2, 2, 2, 59, 145,
	3, 2, 2, 2, 61, 162, 3, 2, 2, 2, 63, 168, 3, 2, 2, 2, 65, 66, 7, 118, 2,
	2, 66, 67, 7, 116, 2, 2, 67, 68, 7, 119, 2, 2, 68, 69, 7, 103, 2, 2, 69,
	4, 3, 2, 2, 2, 70, 71, 7, 104, 2, 2, 71, 72, 7, 99, 2, 2, 72, 73, 7, 110,
	2, 2, 73, 74, 7, 117, 2, 2, 74, 75, 7, 103, 2, 2, 75, 6, 3, 2, 2, 2, 76,
	77, 7, 35, 2, 2, 77, 8, 3, 2, 2, 2, 78, 79, 7, 41, 2, 2, 79, 10, 3, 2,
	2, 2, 80, 81, 7, 125, 2, 2, 81, 12, 3, 2, 2, 2, 82, 83, 7, 127, 2, 2, 83,
	14, 3, 2, 2, 2, 84, 85, 7, 42, 2, 2, 85, 16, 3, 2, 2, 2, 86, 87, 7, 43,
	2, 2, 87, 18, 3, 2, 2, 2, 88, 89, 7, 126, 2, 2, 89, 20, 3, 2, 2, 2, 90,
	91, 7, 63, 2, 2, 91, 22, 3, 2, 2, 2, 92, 93, 7, 128, 2, 2, 93, 94, 7, 63,
	2, 2, 94, 24, 3, 2, 2, 2, 95, 96, 7, 126, 2, 2, 96, 97, 7, 63, 2, 2, 97,
	26, 3, 2, 2, 2, 98, 99, 7, 63, 2, 2, 99, 100, 7, 126, 2, 2, 100, 28, 3,
	2, 2, 2, 101, 102, 7, 107, 2, 2, 102, 103, 7, 112, 2, 2, 103, 30, 3, 2,
	2, 2, 104, 105, 7, 64, 2, 2, 105, 32, 3, 2, 2, 2, 106, 107, 7, 64, 2, 2,
	107, 108, 7, 63, 2, 2, 108, 34, 3, 2, 2, 2, 109, 110, 7, 62, 2, 2, 110,
	36, 3, 2, 2, 2, 111, 112, 7, 62, 2, 2, 112, 113, 7, 63, 2, 2, 113, 38,
	3, 2, 2, 2, 114, 115, 7, 64, 2, 2, 115, 116, 7, 62, 2, 2, 116, 40, 3, 2,
	2, 2, 117, 118, 7, 62, 2, 2, 118, 119, 7, 64, 2, 2, 119, 42, 3, 2, 2, 2,
	120, 121, 7, 48, 2, 2, 121, 44, 3, 2, 2, 2, 122, 123, 7, 46, 2, 2, 123,
	46, 3, 2, 2, 2, 124, 125, 7, 60, 2, 2, 125, 48, 3, 2, 2, 2, 126, 128, 5,
	51, 26, 2, 127, 126, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 127, 3, 2,
	2, 2, 129, 130, 3, 2, 2, 2, 130, 50, 3, 2, 2, 2, 131, 132, 9, 2, 2, 2,
	132, 52, 3, 2, 2, 2, 133, 134, 9, 3, 2, 2, 134, 54, 3, 2, 2, 2, 135, 136,
	9, 4, 2, 2, 136, 56, 3, 2, 2, 2, 137, 142, 5, 53, 27, 2, 138, 141, 5, 53,
	27, 2, 139, 141, 5, 51, 26, 2, 140, 138, 3, 2, 2, 2, 140, 139, 3, 2, 2,
	2, 141, 144, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143,
	58, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 145, 146, 7, 49, 2, 2, 146, 147,
	7, 49, 2, 2, 147, 151, 3, 2, 2, 2, 148, 150, 11, 2, 2, 2, 149, 148, 3,
	2, 2, 2, 150, 153, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 151, 149, 3, 2, 2,
	2, 152, 155, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 154, 156, 7, 15, 2, 2, 155,
	154, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 158,
	7, 12, 2, 2, 158, 159, 3, 2, 2, 2, 159, 160, 8, 30, 2, 2, 160, 60, 3, 2,
	2, 2, 161, 163, 9, 5, 2, 2, 162, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2,
	164, 162, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166,
	167, 8, 31, 2, 2, 167, 62, 3, 2, 2, 2, 168, 169, 7, 49, 2, 2, 169, 170,
	7, 44, 2, 2, 170, 174, 3, 2, 2, 2, 171, 173, 11, 2, 2, 2, 172, 171, 3,
	2, 2, 2, 173, 176, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 174, 172, 3, 2, 2,
	2, 175, 177, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 177, 178, 7, 44, 2, 2, 178,
	179, 7, 49, 2, 2, 179, 180, 3, 2, 2, 2, 180, 181, 8, 32, 2, 2, 181, 64,
	3, 2, 2, 2, 11, 2, 129, 140, 142, 151, 155, 162, 164, 174, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'true'", "'false'", "'!'", "'''", "'{'", "'}'", "'('", "')'", "'|'",
	"'='", "'~='", "'|='", "'=|'", "'in'", "'>'", "'>='", "'<'", "'<='", "'><'",
	"'<>'", "'.'", "','", "':'",
}

var lexerSymbolicNames = []string{
	"", "", "", "R_NEGATE", "R_SIGLE_QUOTA", "R_BRACE_L", "R_BRACE_R", "R_BRACKET_L",
	"R_BRACKET_R", "R_PIPE", "R_EQ", "R_CONTAINS", "R_START_WITH", "R_END_WITH",
	"R_IN", "R_GT", "R_GTE", "R_LT", "R_LTE", "R_BETWEEN", "R_NOT_BETWEEN",
	"R_DOT", "R_COMA", "R_COLON", "T_INTEIRO", "ID", "R_LINE_COMMENT", "R_WS",
	"R_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "R_NEGATE", "R_SIGLE_QUOTA", "R_BRACE_L", "R_BRACE_R",
	"R_BRACKET_L", "R_BRACKET_R", "R_PIPE", "R_EQ", "R_CONTAINS", "R_START_WITH",
	"R_END_WITH", "R_IN", "R_GT", "R_GTE", "R_LT", "R_LTE", "R_BETWEEN", "R_NOT_BETWEEN",
	"R_DOT", "R_COMA", "R_COLON", "T_INTEIRO", "DIGITO", "CHARACTER", "CHARACTER_UP",
	"ID", "R_LINE_COMMENT", "R_WS", "R_COMMENT",
}

type GrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewGrammarLexer(input antlr.CharStream) *GrammarLexer {

	l := new(GrammarLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Grammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GrammarLexer tokens.
const (
	GrammarLexerT__0           = 1
	GrammarLexerT__1           = 2
	GrammarLexerR_NEGATE       = 3
	GrammarLexerR_SIGLE_QUOTA  = 4
	GrammarLexerR_BRACE_L      = 5
	GrammarLexerR_BRACE_R      = 6
	GrammarLexerR_BRACKET_L    = 7
	GrammarLexerR_BRACKET_R    = 8
	GrammarLexerR_PIPE         = 9
	GrammarLexerR_EQ           = 10
	GrammarLexerR_CONTAINS     = 11
	GrammarLexerR_START_WITH   = 12
	GrammarLexerR_END_WITH     = 13
	GrammarLexerR_IN           = 14
	GrammarLexerR_GT           = 15
	GrammarLexerR_GTE          = 16
	GrammarLexerR_LT           = 17
	GrammarLexerR_LTE          = 18
	GrammarLexerR_BETWEEN      = 19
	GrammarLexerR_NOT_BETWEEN  = 20
	GrammarLexerR_DOT          = 21
	GrammarLexerR_COMA         = 22
	GrammarLexerR_COLON        = 23
	GrammarLexerT_INTEIRO      = 24
	GrammarLexerID             = 25
	GrammarLexerR_LINE_COMMENT = 26
	GrammarLexerR_WS           = 27
	GrammarLexerR_COMMENT      = 28
)
