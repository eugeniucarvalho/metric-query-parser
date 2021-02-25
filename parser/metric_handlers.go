package parser

import "fmt"

type ParseHandler = func(Context, map[string]interface{}) (interface{}, error)
type Handlers = map[string]ParseHandler

var (
	handlersMap = map[string]ParseHandler{
		"from":          from,
		"sort":          proxyHandlerFactory("sort"),
		"count":         proxyHandlerFactory("count"),
		"match":         proxyHandlerFactory("match"),
		"query":         proxyHandlerFactory("query"),
		"group":         proxyHandlerFactory("group"),
		"limit":         proxyHandlerFactory("limit"),
		"skip":          proxyHandlerFactory("skip"),
		"lookup":        proxyHandlerFactory("lookup"),
		"sort.default":  proxyHandlerFactory("sort"),
		"count.default": proxyHandlerFactory("count"),
		"match.default": proxyHandlerFactory("match"),
		"query.default": proxyHandlerFactory("query"),
		"group.default": proxyHandlerFactory("group"),
		"limit.default": proxyHandlerFactory("limit"),
		"skip.default":  proxyHandlerFactory("skip"),
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

func proxyHandlerFactory(prefix string) func(context Context, props map[string]interface{}) (interface{}, error) {
	return func(context Context, props map[string]interface{}) (interface{}, error) {
		return handlerByContextType(prefix, context, props)
	}
}
