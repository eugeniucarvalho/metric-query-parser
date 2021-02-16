package main

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/eugeniucarvalho/metric-query-parser/parser"
)

func main() {
	parser.RegisterHandler("input", func(props map[string]interface{}) (interface{}, error) {
		spew.Dump(props)
		return fmt.Sprintf("%v.%v", props["previous.handler.result"], props["vivo"]), nil
	})

	parser.RegisterHandler("input2", func(props map[string]interface{}) (interface{}, error) {
		return props["nome"], nil
	})
	input := `{
		"prop1": { 
			"query": "input(input2(nome='eugenio asdasd'),numero=1,vivo=true)",
			"time": 1613477429690
		},
		"prop2": {"query": "input(input2(nome='vandressa'),numero=2,vivo=true)", "time": 1613477429690}
	}`

	metrics := parser.MetricsMap{}

	if err := json.Unmarshal([]byte(input), &metrics); err != nil {
		panic(err)
	}

	parser := parser.NewMetricQueryParser()

	result, _ := parser.Resolve(metrics)

	spew.Dump(result)

}
