package reflect

import (
	"fmt"
	"reflect"
)

func IterateFunc(entity any) (map[string]FuncInfo, error) {

	typ := reflect.TypeOf(entity)

	res := map[string]FuncInfo{}

	// Iterate over all methods
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		fn := method.Func

		numIn := fn.Type().NumIn() // 函数的类型信息中，有输入参数数量信息
		input := make([]reflect.Type, 0, numIn)
		inputVals := make([]reflect.Value, 0, numIn)

		// 对于GetAge函数，第一个参数是entity本身，所以需要特殊处理
		input = append(input, reflect.TypeOf(entity))
		inputVals = append(inputVals, reflect.ValueOf(entity))

		// 其他输入参数
		for j := 1; j < numIn; j++ {
			fnIntype := fn.Type().In(j)
			input = append(input, fnIntype)
			inputVals = append(inputVals, reflect.Zero(fnIntype))
		}

		// 处理输出参数
		numOut := fn.Type().NumOut()

		output := make([]reflect.Type, 0, numOut)

		for j := 0; j < numOut; j++ {
			output = append(output, fn.Type().Out(j))
		}

		fmt.Printf("Method: %s, Input: %v, Output: %v\n", method.Name, inputVals, output)

		// 函数调用执行
		resVals := fn.Call(inputVals)

		result := make([]any, numOut)
		for j := 0; j < numOut; j++ {
			result[j] = resVals[j].Interface()
		}

		res[method.Name] = FuncInfo{
			Name:        method.Name,
			InputTypes:  input,
			OutputTypes: output,
			Result:      result,
		}
	}

	return res, nil
}

type FuncInfo struct {
	Name        string
	InputTypes  []reflect.Type
	OutputTypes []reflect.Type

	Result []any
}
