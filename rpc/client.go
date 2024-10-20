package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
)

func InitClientProxy(service Service) any {

	if service == nil {
		return errors.New("rpc: 不支持nil")
	}

	val := reflect.ValueOf(service)
	typ := val.Type()

	if typ.Kind() != reflect.Pointer || typ.Elem().Kind() != reflect.Struct {
		return errors.New("rpc: 只支持指向结构体的一级指针")
	}

	val = val.Elem()
	typ = typ.Elem()

	numField := typ.NumField()

	for i := 0; i < numField; i++ {
		fieldTyp := typ.Field(i)
		fieldVal := val.Field(i)

		fn := func(args []reflect.Value) (results []reflect.Value) {

			ctx := args[0].Interface().(context.Context)

			req := &Request{
				ServiceName: service.Name(),
				MethodName:  fieldTyp.Name,
				Args:        []any{args[1].Interface()},
			}

			var p Proxy = NewClient()

			//  构造结构体
			retVal := reflect.New(fieldTyp.Type.Out(0))

			resp, err := p.Invoke(ctx, req)
			if err != nil {
				return []reflect.Value{retVal, reflect.ValueOf(err)}
			}

			//  反射赋值
			err = json.Unmarshal(resp.Data, retVal.Interface())
			if err != nil {
				return []reflect.Value{retVal, reflect.ValueOf(err)}
			}

			return []reflect.Value{retVal, reflect.Zero(reflect.TypeOf(new(error)).Elem())}
		}

		if fieldVal.CanSet() {
			fnVal := reflect.MakeFunc(fieldTyp.Type, fn)
			fieldVal.Set(fnVal)
		}

	}

	return nil
}

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Invoke(ctx context.Context, req *Request) (*Response, error) {
	return nil, nil
}
