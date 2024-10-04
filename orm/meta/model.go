package meta

import (
	"go-actions/orm/internal/errs"
	"reflect"
	"strings"
	"sync"
	"unicode"
)

var (
	ErrInvalidModelType = errs.ErrPointerOnly
)

const (
	tagKeyColumn = "column"
)

type TableName interface {
	TableName() string
}

type ModelOpt func(*Model) error

type Registry interface {
	Get(val any) (*Model, error)
	Register(val any, opts ...ModelOpt) (*Model, error)
}

type Model struct {
	TableName string
	FieldMap  map[string]*Column // 结构体字段的映射
}

type Column struct {
	ColName string
	GoName  string
	Type    reflect.Type
	Offset  uintptr
}

type registry struct {
	// lock sync.RWMutex
	// models map[reflect.Type]*model
	models sync.Map
}

func NewRegistry() *registry {
	return &registry{
		// models: make(map[reflect.Type]*model, 64),
	}
}

func (r *registry) Get(val any) (*Model, error) {
	typ := reflect.TypeOf(val)
	m, ok := r.models.Load(typ)
	if ok {
		return m.(*Model), nil
	}
	return r.Register(val)
}

// func (r *registry) get1(val any) (*model, error) {
// 	typ := reflect.TypeOf(val)
// 	r.lock.RLock()
// 	m, ok := r.models[typ]
// 	r.lock.RUnlock()
// 	if !ok {
// 		var err error

// 		r.lock.Lock()
// 		defer r.lock.Unlock()
// 		m, ok = r.models[typ]
// 		if ok {
// 			return m, nil
// 		}

// 		m, err = r.parseModel(val)
// 		if err != nil {
// 			return nil, err
// 		}
// 		r.models[reflect.TypeOf(val)] = m
// 	}

// 	return m, nil
// }

func (r *registry) Register(val any, opts ...ModelOpt) (*Model, error) {
	m, err := r.parseModel(val)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		err := opt(m)
		if err != nil {
			return nil, err
		}
	}

	typ := reflect.TypeOf(val)
	r.models.Store(typ, m)
	return m, nil
}

func ModelWithTableName(tableName string) ModelOpt {
	return func(m *Model) error {
		m.TableName = tableName
		return nil
	}
}

func ModelWithColumnName(field, colName string) ModelOpt {
	return func(m *Model) error {
		fd, ok := m.FieldMap[field]
		if !ok {
			return errs.NewErrUnknownField(field)
		}
		fd.ColName = colName
		return nil
	}
}

func (r *registry) parseModel(val any) (*Model, error) {
	typ := reflect.TypeOf(val)

	if typ.Kind() != reflect.Pointer || typ.Elem().Kind() != reflect.Struct {
		return nil, ErrInvalidModelType
	}
	typ = typ.Elem()

	numField := typ.NumField()
	fds := make(map[string]*Column, numField)

	for i := 0; i < numField; i++ {
		f := typ.Field(i)
		pair, err := r.parseTag(f.Tag)
		if err != nil {
			return nil, err
		}

		columnName, ok := pair[tagKeyColumn]
		if columnName == "" || !ok {
			columnName = underscoreName(f.Name)
		}

		fds[f.Name] = &Column{
			GoName:  f.Name,
			ColName: columnName,
			Type:    f.Type,
			Offset:  f.Offset,
		}
	}

	var tableName string
	if tbl, ok := val.(TableName); ok { // 类型断言为实现TableName()函数的接口，然后调用该接口获取表名
		tableName = tbl.TableName()
	}

	if tableName == "" {
		tableName = underscoreName(typ.Name())
	}

	return &Model{
		TableName: tableName,
		FieldMap:  fds,
	}, nil
}

func (r *registry) parseTag(tag reflect.StructTag) (map[string]string, error) {
	ormTag, ok := tag.Lookup("orm")
	if !ok {
		return map[string]string{}, nil
	}

	res := make(map[string]string, 1)
	pairs := strings.Split(ormTag, ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			return nil, errs.NewErrInvalidTagContent(pair)
		}
		res[kv[0]] = kv[1]
	}
	return res, nil
}

func underscoreName(name string) string {
	var buf []byte

	for i, v := range name {
		if unicode.IsUpper(v) {
			if i != 0 {
				buf = append(buf, '_')
			}
			buf = append(buf, byte(unicode.ToLower(v)))
		} else {
			buf = append(buf, byte(v))
		}
	}

	return string(buf)
}
