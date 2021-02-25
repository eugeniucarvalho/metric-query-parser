package mongo

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/eugeniucarvalho/metric-query-parser/parser"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	Handlers = parser.Handlers{
		"match.aggregate":  match,
		"sort.aggregate":   sort,
		"count.aggregate":  count,
		"group.aggregate":  group,
		"limit.aggregate":  limit,
		"skip.aggregate":   skip,
		"lookup.aggregate": lookup,
	}

	ErrPipelineUnidefined = func(method string) error {
		return fmt.Errorf("Can't find pipeline instance. Set one before call method'%s'", method)
	}
	ErrGroupIdIsRequired = fmt.Errorf("Group requires an id.")
)

func match(context parser.Context, props map[string]interface{}) (interface{}, error) {
	var (
		found    bool
		pipeline *bson.A
		query    = bson.M{}
	)

	// if pipeline, found = context.Get("pipeline").(*bson.A); !found {
	// 	return nil, ErrPipelineUnidefined("match")
	// }
	if pipeline, found = props["_$context"].(*bson.A); !found {
		pipeline = &bson.A{}
		// return nil, ErrPipelineUnidefined("match")
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

	// if pipeline, found = context.Get("pipeline").(*bson.A); !found {
	// 	return nil, ErrPipelineUnidefined("sort")
	// }
	if pipeline, found = props["_$context"].(*bson.A); !found {
		pipeline = &bson.A{}
		// return nil, ErrPipelineUnidefined("match")
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

	// if pipeline, found = context.Get("pipeline").(*bson.A); !found {
	// 	return nil, ErrPipelineUnidefined("count")
	// }
	if pipeline, found = props["_$context"].(*bson.A); !found {
		pipeline = &bson.A{}
		// return nil, ErrPipelineUnidefined("match")
	}

	context.Set("singleResult", true)

	if as, found = props["as"].(string); !found {
		as = "count"
	}

	*pipeline = append(*pipeline, bson.M{"$count": as})

	return pipeline, nil
}

func group(context parser.Context, props map[string]interface{}) (interface{}, error) {
	var (
		input    string
		found    bool
		pipeline *bson.A
		id       = bson.A{}
	)

	// if pipeline, found = context.Get("pipeline").(*bson.A); !found {
	// 	return nil, ErrPipelineUnidefined("count")
	// }
	if pipeline, found = props["_$context"].(*bson.A); !found {
		pipeline = &bson.A{}
		// return nil, ErrPipelineUnidefined("match")
	}

	if input, found = props["id"].(string); !found {
		return nil, ErrGroupIdIsRequired
	}

	idProps := strings.Split(input, ",")
	idPropsLimit := len(idProps) - 1
	for index, prop := range idProps {
		id = append(id, fmt.Sprintf("$%s", prop))
		if index < idPropsLimit {
			id = append(id, "-")
		}
	}

	*pipeline = append(*pipeline, bson.M{
		"$group": bson.M{
			"_id":   bson.M{"$concat": id},
			"count": bson.M{"$sum": 1},
		},
	})

	return pipeline, nil
}

func limit(context parser.Context, props map[string]interface{}) (interface{}, error) {
	var (
		count    int64
		found    bool
		pipeline *bson.A
	)

	// if pipeline, found = context.Get("pipeline").(*bson.A); !found {
	// 	return nil, ErrPipelineUnidefined("count")
	// }
	if pipeline, found = props["_$context"].(*bson.A); !found {
		// return nil, ErrPipelineUnidefined("match")
		pipeline = &bson.A{}
	}

	if count, found = props["count"].(int64); !found {
		return nil, ErrGroupIdIsRequired
	}

	*pipeline = append(*pipeline, bson.M{"$limit": count})
	return pipeline, nil
}

func skip(context parser.Context, props map[string]interface{}) (interface{}, error) {
	var (
		count    int64
		found    bool
		pipeline *bson.A
	)

	// if pipeline, found = context.Get("pipeline").(*bson.A); !found {
	// 	return nil, ErrPipelineUnidefined("count")
	// }
	if pipeline, found = props["_$context"].(*bson.A); !found {
		pipeline = &bson.A{}
		// return nil, ErrPipelineUnidefined("match")
	}

	if count, found = props["count"].(int64); !found {
		return nil, ErrGroupIdIsRequired
	}

	*pipeline = append(*pipeline, bson.M{"$skip": count})
	return pipeline, nil
}

func lookup(context parser.Context, props map[string]interface{}) (interface{}, error) {
	var (
		found          bool
		let, as, from  string
		pipe, pipeline *bson.A
		lookup         = bson.M{}
	)

	// if pipeline, found = context.Get("pipeline").(*bson.A); !found {
	// 	return nil, ErrPipelineUnidefined("count")
	// }

	if pipeline, found = props["_$context"].(*bson.A); !found {
		pipeline = &bson.A{}
		// return nil, ErrPipelineUnidefined("match")
	}

	if from, found = props["from"].(string); found {
		lookup["from"] = from
	}

	if as, found = props["as"].(string); found {
		lookup["as"] = as
	}

	if pipe, found = props["pipeline"].(*bson.A); found {
		lookup["pipeline"] = pipe
	}

	spew.Dump(props)

	if let, found = props["let"].(string); found {
		terms := bson.M{}
		for _, term := range strings.Split(let, ",") {
			vars := strings.Split(term, "as")
			terms[vars[1]] = fmt.Sprintf("$%s", vars[0])
		}
		lookup["let"] = terms
	}
	*pipeline = append(*pipeline, bson.M{"$lookup": lookup})
	return pipeline, nil
}
