package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type MetricQueryParser struct {
}

func (parser *MetricQueryParser) Resolve(context Context) (result map[string]interface{}, err error) {
	result = map[string]interface{}{}

	for prop, metric := range context.Metrics {
		if result[prop], err = parser.ResolveMetric(context, metric); err != nil {
			// result[prop] = map[string]interface{}{"error": err}
		}
	}
	return
}

func (parser *MetricQueryParser) ResolveMetric(context Context, metric *MetricQueryItem) (result interface{}, err error) {

	listener := NewQueryListener(context.Copy(), metric)
	input := antlr.NewInputStream(metric.Query)
	lexer := NewGrammarLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewGrammarParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	expression := p.Expression()
	// fmt.Println(expression.ToStringTree(p.GetRuleNames(), p))
	antlr.ParseTreeWalkerDefault.Walk(listener, expression)

	result = listener.Result()

	return
}

func NewMetricQueryParser() *MetricQueryParser {
	return &MetricQueryParser{}
}
