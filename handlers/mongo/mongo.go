package mongo

import (
	"fmt"
	"strings"

	"github.com/eugeniucarvalho/metric-query-parser/parser"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	Handlers = parser.Handlers{
		"match.aggregate": match,
		"sort.aggregate":  sort,
		"count.aggregate": count,
	}

	ErrPipelineUnidefined = func(method string) error {
		return fmt.Errorf("Can't find pipeline instance. Set one before call method'%s'", method)
	}
)

func match(context parser.Context, props map[string]interface{}) (interface{}, error) {
	var (
		found    bool
		pipeline *bson.A
		query    = bson.M{}
	)

	if pipeline, found = context.Get("pipeline").(*bson.A); !found {
		return nil, ErrPipelineUnidefined("match")
	}

	for porp, value := range props {
		if strings.HasPrefix(porp, "_") {
			continue
		}
		query[porp] = value
	}

	*pipeline = append(*pipeline, bson.M{"$match": query})
	return pipeline, nil
}

func sort(context parser.Context, props map[string]interface{}) (interface{}, error) {
	var (
		found    bool
		pipeline *bson.A
		sort     = bson.M{}
	)

	if pipeline, found = context.Get("pipeline").(*bson.A); !found {
		return nil, ErrPipelineUnidefined("sort")
	}

	for porp, value := range props {
		if strings.HasPrefix(porp, "_") {
			continue
		}
		sort[porp] = value
	}

	*pipeline = append(*pipeline, bson.M{"$sort": sort})
	return pipeline, nil
}

func count(context parser.Context, props map[string]interface{}) (interface{}, error) {
	var (
		as       string
		found    bool
		pipeline *bson.A
	)

	if pipeline, found = context.Get("pipeline").(*bson.A); !found {
		return nil, ErrPipelineUnidefined("count")
	}

	if as, found = props["as"].(string); !found {
		as = "count"
	}

	*pipeline = append(*pipeline, bson.M{"$count": as})

	return pipeline, nil
}
