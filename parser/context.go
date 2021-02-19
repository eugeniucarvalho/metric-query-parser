package parser

type context struct {
	values  map[string]interface{}
	Metrics MetricsMap
}

type Context = *context

func NewContext() Context {
	return &context{
		values:  make(map[string]interface{}),
		Metrics: make(map[string]*MetricQueryItem),
	}
}

func (ctx *context) Set(prop string, value interface{}) {
	ctx.values[prop] = value
}

func (ctx context) Get(prop string) (interface{}, bool) {
	value, found := ctx.values[prop]
	return value, found
}
