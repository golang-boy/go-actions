package reflect

import (
	"fmt"
	"reflect"
	"time"
)

var ErrMustBeStruct = fmt.Errorf("entity must be a struct")
var ErrNotSupportNilType = fmt.Errorf("entity type must not be nil")
var ErrNotSupportZeroVal = fmt.Errorf("entity type must not be zero value")
var ErrCantSetField = fmt.Errorf("can`t set field")

func IterateFields(entity any) (map[string]any, error) {
	if entity == nil {
		return nil, ErrNotSupportNilType
	}

	typ := reflect.TypeOf(entity)

	val := reflect.ValueOf(entity)
	if val.IsZero() {
		return nil, ErrNotSupportZeroVal
	}

	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
		val = val.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return nil, ErrMustBeStruct
	}

	res := map[string]any{}
	for i := 0; i < typ.NumField(); i++ {
		fieldType := typ.Field(i)
		fieldValue := val.Field(i)

		if fieldType.IsExported() {
			res[fieldType.Name] = fieldValue.Interface()
		} else {
			res[fieldType.Name] = reflect.Zero(fieldType.Type).Interface() // zero value
		}
	}

	return res, nil
}

func SetField(entity any, field string, newValue any) error {
	val := reflect.ValueOf(entity)
	for val.Type().Kind() == reflect.Pointer {
		val = val.Elem()
	}

	fieldVal := val.FieldByName(field)
	if !fieldVal.CanSet() {
		return ErrCantSetField
	}
	fieldVal.Set(reflect.ValueOf(newValue))
	return nil
}

type MetaJson struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Tag   string `json:"tag"`
	Value any    `json:"value"`
}

type FuncJson struct {
	Name   string `json:"name"`
	Args   []MetaJson
	Output []MetaJson
}

func DefineStruct(metas []MetaJson) any {

	typeFunc := func(typ string) reflect.Type {
		switch typ {
		case "int":
			return reflect.TypeOf(0)
		case "string":
			return reflect.TypeOf("")
		case "float64":
			return reflect.TypeOf(0.0)
		case "bool":
			return reflect.TypeOf(false)
		case "time.Time":
			return reflect.TypeOf(time.Time{})
		case "map": // 需要对值进行更进一步的定义与设计
			return reflect.TypeOf(map[string]string{})
		default:
			return nil
		}
	}

	fields := make([]reflect.StructField, 0, len(metas))
	for i := 0; i < len(metas); i++ {
		fields = append(fields, reflect.StructField{
			Name: metas[i].Name,
			Type: typeFunc(metas[i].Type),
			Tag:  reflect.StructTag(metas[i].Tag),
		})
	}

	structType := reflect.StructOf(fields)

	val := reflect.New(structType)
	vale := val.Elem()

	for i := 0; i < len(metas); i++ {

		fieldVal := reflect.ValueOf(metas[i].Value)
		targetField := vale.FieldByName(metas[i].Name)

		if fieldVal.Type().AssignableTo(targetField.Type()) {
			targetField.Set(fieldVal)
		} else {
			fmt.Printf("Type mismatch for field %s: expected %s, got %s\n", metas[i].Name, targetField.Type(), fieldVal.Type())
		}
	}
	return val.Interface()
}
