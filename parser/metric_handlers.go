package parser

type ParseHandler = func(map[string]interface{}) (interface{}, error)

var (
	handlers = map[string]ParseHandler{}
)

func RegisterHandler(id string, fn ParseHandler) {
	handlers[id] = fn
}
