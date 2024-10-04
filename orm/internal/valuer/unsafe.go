package valuer

import (
	"database/sql"
	"go-actions/orm/meta"
	"reflect"
	"unsafe"
)

type unsafeValue struct {
	model *meta.Model
	val   any
}

func NewUnsafeValue(model *meta.Model, val any) Valuer {
	return unsafeValue{
		model: model,
		val:   val,
	}
}

func (r unsafeValue) SetColumns(rows *sql.Rows) error {
	cs, err := rows.Columns()
	if err != nil {
		return err
	}

	colVals := make([]any, 0, len(cs))

	tpValue := reflect.ValueOf(r.val)
	address := tpValue.UnsafePointer()

	for _, c := range cs {
		for _, cm := range r.model.FieldMap {
			if cm.ColName == c {
				fdAddr := unsafe.Pointer(uintptr(address) + cm.Offset)
				val := reflect.NewAt(cm.Type, fdAddr)
				colVals = append(colVals, val.Interface())
			}
		}
	}

	err = rows.Scan(colVals...)
	if err != nil {
		return err
	}
	return nil
}
