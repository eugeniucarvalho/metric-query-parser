package parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type MetricQueryParser struct {
}

func (parser *MetricQueryParser) Resolve(metrics MetricsMap) (result map[string]interface{}, err error) {
	result = map[string]interface{}{}

	for prop, metric := range metrics {
		if result[prop], err = parser.ResolveMetric(metric); err != nil {
			return
		}
	}
	return
}

func (parser *MetricQueryParser) ResolveMetric(context *MetricQueryItem) (result interface{}, err error) {

	listener := NewQueryListener(context)
	input := antlr.NewInputStream(context.Query)
	// input := antlr.NewInputStream(`aaa`)
	lexer := NewGrammarLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewGrammarParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Expression())

	// result, err = listener.Result()
	result = listener.Result()

	return
}

func NewMetricQueryParser() *MetricQueryParser {
	return &MetricQueryParser{}
}
