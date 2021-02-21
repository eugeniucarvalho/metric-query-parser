package main

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
	mongo "github.com/eugeniucarvalho/metric-query-parser/handlers/mongo"
	"github.com/eugeniucarvalho/metric-query-parser/parser"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	parser.RegisterRandlers(mongo.Handlers)

	parser.RegisterRandlers(parser.Handlers{
		// "input2": func(context parser.Context, props map[string]interface{}) (interface{}, error) {
		// 	return props["nome"], nil
		// },
		// "input": func(context parser.Context, props map[string]interface{}) (interface{}, error) {
		// 	spew.Dump(props)
		// 	return fmt.Sprintf("%v.%v", props["previous.handler.result"], props["vivo"]), nil
		// },

		"query": func(context parser.Context, props map[string]interface{}) (interface{}, error) {
			var (
				found    bool
				pipeline *bson.A
			)

			if pipeline, found = context.Get("pipeline").(*bson.A); !found {

			}

			out, _ := json.Marshal(pipeline)
			// fmt.Println(string(out))
			return string(out), nil
		},
		"from.accounts": func(context parser.Context, props map[string]interface{}) (interface{}, error) {
			context.Set("type", "aggregate")
			context.Set("pipeline", &bson.A{})
			context.Set("db", "accounts")
			context.Set("collection", "accounts")
			// context.Set("entity", "")
			return nil, nil
		},
	})

	// input := `{
	// 	"prop1": {
	// 		"query": "input(input2(nome='eugenio asdasd'),numero=1,vivo=true)",
	// 		"time": 1613477429690
	// 	},
	// 	"prop2": {"query": "input(input2(nome='vandressa'),numero=2,vivo=true)", "time": 1613477429690}
	// }`
	input := `{
		"accounts": { 
			"query": "query(count(sort(match(from(name='accounts'), deleted=false), name=1)))"
		},
		"accounts2": { 
			"query": "query(count(sort(match(deleted=false), name=1)))"
		}
	}`

	context := parser.NewContext()

	if err := json.Unmarshal([]byte(input), &context.Metrics); err != nil {
		panic(err)
	}

	parser := parser.NewMetricQueryParser()

	result, _ := parser.Resolve(context)

	spew.Dump(result)

}
