package parser

type context struct {
	values   map[string]interface{}
	Metrics  MetricsMap
	Handlers Handlers
}

type Context = *context

func NewContext() Context {
	return &context{
		values:   make(map[string]interface{}),
		Metrics:  make(map[string]*MetricQueryItem),
		Handlers: handlersMap,
	}
}

func (ctx *context) Copy() Context {
	newContext := NewContext()
	newContext.Metrics = ctx.Metrics
	newContext.Handlers = ctx.Handlers
	for prop, value := range ctx.values {
		newContext.values[prop] = value
	}
	return newContext
}

func (ctx *context) Set(prop string, value interface{}) {
	ctx.values[prop] = value
}

func (ctx context) Get(prop string) interface{} {
	return ctx.values[prop]
}

func (ctx context) HasDataSource() bool {
	_, found := ctx.Get("hasDataSource").(bool)
	return found
}

func (ctx context) GetHandler(name string) (ParseHandler, error) {
	var (
		handler ParseHandler
		found   bool
	)
	if handler, found = ctx.Handlers[name]; !found {
		return nil, UndefinedHandlerError(name)
	}
	return handler, nil
}
