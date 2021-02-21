package parser

import "fmt"

type ParseHandler = func(Context, map[string]interface{}) (interface{}, error)
type Handlers = map[string]ParseHandler

var (
	handlersMap = map[string]ParseHandler{
		"from":          from,
		"sort":          sort,
		"count":         count,
		"match":         match,
		"query":         query,
		"sort.default":  sort,
		"count.default": count,
		"match.default": match,
		"query.default": query,
	}

	UndefinedHandlerError = func(name string) error {
		return fmt.Errorf("Can't find handler '%s'", name)
	}
	ErrDataSourceEmpty      = fmt.Errorf("DataSource name can't be empty")
	ErrDataSourceUnidefined = fmt.Errorf("Can't find any dataSource. Use from(name='name_of_dataSource') method to set one")
)

func RegisterHandler(id string, fn ParseHandler) {
	handlersMap[id] = fn
}

func RegisterRandlers(handlers Handlers) {
	for id, handler := range handlers {
		RegisterHandler(id, handler)
	}
}

func handlerByContextType(prefix string, context Context, props map[string]interface{}) (interface{}, error) {
	var (
		typ     string
		found   bool
		handler ParseHandler
		err     error
	)

	if !context.HasDataSource() {
		return nil, ErrDataSourceUnidefined
	}

	if typ, found = context.Get("type").(string); !found {
		typ = "default"
	}

	if handler, err = context.GetHandler(fmt.Sprintf("%s.%s", prefix, typ)); err != nil {
		return nil, err
	}

	return handler(context, props)
}

func from(context Context, props map[string]interface{}) (interface{}, error) {
	var (
		name    string
		found   bool
		handler ParseHandler
		err     error
	)

	if name, found = props["name"].(string); !found {
		return nil, ErrDataSourceEmpty
	}

	if handler, err = context.GetHandler(fmt.Sprintf("from.%s", name)); err != nil {
		return nil, err
	}

	context.Set("hasDataSource", true)

	return handler(context, props)
}

func sort(context Context, props map[string]interface{}) (interface{}, error) {
	return handlerByContextType("sort", context, props)
}
func count(context Context, props map[string]interface{}) (interface{}, error) {
	return handlerByContextType("count", context, props)
}
func match(context Context, props map[string]interface{}) (interface{}, error) {
	return handlerByContextType("match", context, props)
}
func query(context Context, props map[string]interface{}) (interface{}, error) {
	return handlerByContextType("query", context, props)
}
