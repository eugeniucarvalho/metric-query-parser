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
	typeString  = reflect.TypeOf(&TypeStringContext{})
	typeBool    = reflect.TypeOf(&TypeBoolContext{})
	typeInt     = reflect.TypeOf(&TypeIntContext{})
	typeFloat   = reflect.TypeOf(&TypeFloatContext{})
	typeHandler = reflect.TypeOf(&TypeHandlerContext{})

	typesList = map[interface{}]func(interface{}) interface{}{
		typeString: func(v interface{}) interface{} {
			return strings.Trim(v.(*TypeStringContext).GetText(), "\"")
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

func (ql *QueryListener) ExitArgument(ctx *ArgumentContext) {
	// if s.Error != nil {
	// 	return
	// }
	var name string

	if ctx.ID() != nil {
		name = ctx.ID().GetText()
	} else {
		name = "_$context"
	}

	switch {
	case ctx.Value() != nil:
		valueExpression := ctx.Value().GetChild(0)
		if reflect.TypeOf(valueExpression) != typeHandler {
			operator := ctx.Operator().GetText()
			for t, value := range typesList {
				if reflect.TypeOf(valueExpression) == t {
					ql.Arguments[name] = value(valueExpression)
					ql.Arguments[fmt.Sprintf("_$%s.op", name)] = operator
					return
				}
			}
			return
		}
		fallthrough
	default:
		ql.Arguments[name] = ql.Arguments["__$lastHandlerResult"]
	}
}

// EnterHandler is called when production handler is entered.
func (ql *QueryListener) EnterHandler(ctx *HandlerContext) {
	// if ql.Error != nil {
	// 	return
	// }
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

	fmt.Println("\n\n\nexecute ....", ctx.ID().GetText(), "with arguments: ")
	// spew.Dump(ql.Arguments)

	if result, err = handler(ql.Context, ql.Arguments); err != nil {
		ql.Error = err
	}

	ql.PopArguments()
	ql.Arguments["__$lastHandlerResult"] = result
}

func (ql *QueryListener) PushArguments() {
	// if ql.Error != nil {
	// 	return
	// }
	arguments := map[string]interface{}{}
	ql.ArgumentsStack = append(ql.ArgumentsStack, arguments)
	ql.Arguments = arguments
}

func (ql *QueryListener) PopArguments() {
	// if ql.Error != nil {
	// 	return
	// }
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
		Value:  ql.Arguments["__$lastHandlerResult"],
	}
}

func NewQueryListener(context Context, metric *MetricQueryItem) *QueryListener {
	return &QueryListener{
		Context: context,
		Metric:  metric,
		ArgumentsStack: []map[string]interface{}{
			{
				"__$lastHandlerResult": nil,
			},
		},
	}
}
