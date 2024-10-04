package unsafe

import (
	"errors"
	"reflect"
	"unsafe"
)

type unsafeAccessor struct {
	fields  map[string]*FieldMeta
	address unsafe.Pointer
}

type FieldMeta struct {
	Offset uintptr
	typ    reflect.Type
}

func NewUnsafeAccessor(entity any) *unsafeAccessor {

	typ := reflect.TypeOf(entity)
	typ = typ.Elem() // 如果是结构体指针，则获取指针指向的结构体类型
	numField := typ.NumField()

	fields := make(map[string]*FieldMeta, numField)

	for i := 0; i < numField; i++ {
		field := typ.Field(i)
		fields[field.Name] = &FieldMeta{
			Offset: field.Offset,
			typ:    field.Type,
		}
	}

	val := reflect.ValueOf(entity)

	return &unsafeAccessor{
		fields:  fields,
		address: val.UnsafePointer(),
	}
}

func (a *unsafeAccessor) GetField(name string) (any, error) {

	fd, ok := a.fields[name]
	if !ok {
		return nil, errors.New("field not found")
	}

	fdAddr := unsafe.Pointer(uintptr(a.address) + fd.Offset)

	return reflect.NewAt(fd.typ, fdAddr).Elem().Interface(), nil

}

func (a *unsafeAccessor) SetField(name string, value any) error {

	fd, ok := a.fields[name]
	if !ok {
		return errors.New("field not found")
	}

	fdAddr := unsafe.Pointer(uintptr(a.address) + fd.Offset)

	reflect.NewAt(fd.typ, fdAddr).Elem().Set(reflect.ValueOf(value))
	return nil

}
