package parser

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type QueryListener struct {
	*BaseGrammarListener
	ArgumentsStack []map[string]interface{}
	Arguments      map[string]interface{}
	Context        *MetricQueryItem
}

var (
	typeBooleanValues = map[string]bool{
		"true":  true,
		"false": false,
		"TRUE":  true,
		"FALSE": false,
		"True":  true,
		"False": false,
	}
	typeString = reflect.TypeOf(&TypeStringContext{})
	typeBool   = reflect.TypeOf(&TypeBoolContext{})
	typeInt    = reflect.TypeOf(&TypeIntContext{})
	typeFloat  = reflect.TypeOf(&TypeFloatContext{})

	typesList = map[interface{}]func(interface{}) interface{}{
		typeString: func(v interface{}) interface{} {
			return strings.Trim(v.(*TypeStringContext).GetText(), "'")
		},
		typeBool: func(v interface{}) interface{} {
			return typeBooleanValues[v.(*TypeBoolContext).GetText()]
		},
		typeInt: func(v interface{}) interface{} {
			value, _ := strconv.ParseInt(v.(*TypeIntContext).GetText(), 10, 64)
			return value
		},
		typeFloat: func(v interface{}) interface{} {
			value, _ := strconv.ParseFloat(v.(*TypeFloatContext).GetText(), 64)
			return value
		},
	}
)

func (s *QueryListener) EnterArgument(ctx *ArgumentContext) {
	if ctx.ID() != nil {
		valueExpression := ctx.Value().GetChild(0)
		for t, value := range typesList {
			if reflect.TypeOf(valueExpression) == t {
				name := ctx.ID().GetText()
				s.Arguments[name] = value(valueExpression)
				s.Arguments[fmt.Sprintf("%s.op", name)] = ctx.Operator().GetText()
				break
			}
		}
	}
}

// EnterHandler is called when production handler is entered.
func (ql *QueryListener) EnterHandler(ctx *HandlerContext) {
	ql.PushArguments()
}

func (ql *QueryListener) ExitHandler(ctx *HandlerContext) {
	var (
		handler ParseHandler
		found   bool
		result  interface{}
	)

	if handler, found = handlers[ctx.ID().GetText()]; !found {

	}

	ql.Arguments["_context"] = ql.Context

	result, _ = handler(ql.Arguments)

	ql.PopArguments()
	ql.Arguments["previous.handler.result"] = result
}

func (ql *QueryListener) PushArguments() {
	arguments := map[string]interface{}{}
	ql.ArgumentsStack = append(ql.ArgumentsStack, arguments)
	ql.Arguments = arguments
}

func (ql *QueryListener) PopArguments() {
	size := len(ql.ArgumentsStack)
	if size > 1 {
		ql.ArgumentsStack = ql.ArgumentsStack[:size-1]
		index := size - 2
		if index > 0 {
			ql.Arguments = ql.ArgumentsStack[index]
		}
	}
}

func (ql *QueryListener) Result() *MetricQueryItemResult {
	return &MetricQueryItemResult{
		Metric: nil,
		Value:  ql.Arguments["previous.handler.result"],
	}
}

func NewQueryListener(context *MetricQueryItem) *QueryListener {
	return &QueryListener{
		Context: context,
		ArgumentsStack: []map[string]interface{}{
			{
				"previous.handler.result": nil,
			},
		},
	}
}
