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
	Context        Context
	Metric         *MetricQueryItem
	Error          error
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
	if s.Error != nil {
		return
	}
	if ctx.ID() != nil {
		valueExpression := ctx.Value().GetChild(0)
		for t, value := range typesList {
			if reflect.TypeOf(valueExpression) == t {
				name := ctx.ID().GetText()
				s.Arguments[name] = value(valueExpression)
				s.Arguments[fmt.Sprintf("_$%s.op", name)] = ctx.Operator().GetText()
				break
			}
		}
	}
}

// EnterHandler is called when production handler is entered.
func (ql *QueryListener) EnterHandler(ctx *HandlerContext) {
	if ql.Error != nil {
		return
	}
	ql.PushArguments()
}

func (ql *QueryListener) ExitHandler(ctx *HandlerContext) {
	var (
		handler ParseHandler
		found   bool
		result  interface{}
		err     error
	)
	if ql.Error != nil {
		return
	}

	if handler, found = handlersMap[ctx.ID().GetText()]; !found {

	}

	ql.Arguments["_$metric"] = ql.Metric

	if result, err = handler(ql.Context, ql.Arguments); err != nil {
		ql.Error = err
	}

	ql.PopArguments()
	ql.Arguments["_$previous.handler.result"] = result
}

func (ql *QueryListener) PushArguments() {
	if ql.Error != nil {
		return
	}
	arguments := map[string]interface{}{}
	ql.ArgumentsStack = append(ql.ArgumentsStack, arguments)
	ql.Arguments = arguments
}

func (ql *QueryListener) PopArguments() {
	if ql.Error != nil {
		return
	}
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
		Error:  ql.Error,
		Value:  ql.Arguments["_$previous.handler.result"],
	}
}

func NewQueryListener(context Context, metric *MetricQueryItem) *QueryListener {
	return &QueryListener{
		Context: context,
		Metric:  metric,
		ArgumentsStack: []map[string]interface{}{
			{
				"_$previous.handler.result": nil,
			},
		},
	}
}
