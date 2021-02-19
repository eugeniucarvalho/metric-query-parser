package parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type MetricQueryParser struct {
}

func (parser *MetricQueryParser) Resolve(context Context) (result map[string]interface{}, err error) {
	result = map[string]interface{}{}

	for prop, metric := range context.Metrics {
		if result[prop], err = parser.ResolveMetric(context, metric); err != nil {
			return
		}
	}
	return
}

func (parser *MetricQueryParser) ResolveMetric(context Context, metric *MetricQueryItem) (result interface{}, err error) {

	listener := NewQueryListener(context, metric)
	input := antlr.NewInputStream(metric.Query)
	// input := antlr.NewInputStream(`aaa`)
	lexer := NewGrammarLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewGrammarParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Expression())

	result = listener.Result()

	return
}

func NewMetricQueryParser() *MetricQueryParser {
	return &MetricQueryParser{}
}
