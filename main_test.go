package main

import (
	"encoding/json"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/eugeniucarvalho/metric-query-parser/parser"
)

func TestParseSuccess(t *testing.T) {

	parser.RegisterHandler("input", func(context map[string]interface{}) (interface{}, error) {
		return map[string]interface{}{}, nil
	})

	parser.RegisterHandler("input2", func(context map[string]interface{}) (interface{}, error) {
		return map[string]interface{}{}, nil
	})

	input := `{
		"prop1": {"query": "input(input2(nome='eugenio asdasd'),numero=1,vivo=true)", time: 1613477429690},
	}`
	// "prop2": {"query": "input(input2(nome='vandressa'),numero=2,vivo=true)", time: 1613477429690}

	metrics := parser.MetricsMap{}

	json.Unmarshal([]byte(input), metrics)

	parser := parser.NewMetricQueryParser()

	result, _ := parser.Resolve(metrics)

	spew.Dump(result)

	// out := json.MarshalIndent(result, "", "  ")

}
