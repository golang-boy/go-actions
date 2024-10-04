package valuer

import (
	"database/sql"
	"go-actions/orm/meta"
	"reflect"
)

type reflectValue struct {
	model *meta.Model
	val   any
}

func NewReflectValue(model *meta.Model, val any) Valuer {
	return reflectValue{
		model: model,
		val:   val,
	}
}

func (r reflectValue) SetColumns(rows *sql.Rows) error {

	cs, err := rows.Columns()
	if err != nil {
		return err
	}

	colVals := make([]any, 0, len(cs))

	// 根据类型构造指针对象数组
	for _, c := range cs {
		for _, cm := range r.model.FieldMap {
			if cm.ColName == c {
				val := reflect.New(cm.Type) // 创建的是指针类型
				colVals = append(colVals, val.Interface())
			}
		}
	}

	// 填充值
	err = rows.Scan(colVals...)
	if err != nil {
		return err
	}

	// 将值赋给结构体
	tpValue := reflect.ValueOf(r.val)

	for i, c := range cs {
		for _, cm := range r.model.FieldMap {
			if cm.ColName == c {
				tpValue.Elem().FieldByName(cm.GoName).
					Set(reflect.ValueOf(colVals[i]).Elem())
			}
		}
	}

	// 类型要匹配
	// 顺序要匹配
	return nil
}
