package meta

import (
	"database/sql"
	"go-actions/orm/internal/errs"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_parseModel(t *testing.T) {
	type args struct {
		val any
	}
	tests := []struct {
		name    string
		args    args
		want    *Model
		wantErr error
	}{
		{
			name: "test model",
			args: args{
				val: &TestModel{},
			},
			want: &Model{
				TableName: "test_model",
				FieldMap: map[string]*Column{
					"Id": {
						ColName: "id",
						GoName:  "Id",
						Type:    reflect.TypeOf(int64(0)),
						Offset:  0,
					},
					"FirstName": {
						ColName: "first_name",
						GoName:  "FirstName",
						Type:    reflect.TypeOf(""),
						Offset:  uintptr(8),
					},
					"Age": {
						ColName: "age",
						GoName:  "Age",
						Type:    reflect.TypeOf(int8(0)),
						Offset:  uintptr(24),
					},
					"LastName": {
						ColName: "last_name",
						GoName:  "LastName",
						Type:    reflect.TypeOf(""),
						Offset:  uintptr(32),
					},
				},
			},
		},
		{
			name: "pointer",
			args: args{
				val: &TestModel{},
			},
			want: &Model{
				TableName: "test_model",
				FieldMap: map[string]*Column{
					"Id": {
						ColName: "id",
						GoName:  "Id",
						Type:    reflect.TypeOf(int64(0)),
						Offset:  0,
					},
					"FirstName": {
						ColName: "first_name",
						GoName:  "FirstName",
						Type:    reflect.TypeOf(""),
						Offset:  8,
					},
					"Age": {
						ColName: "age",
						GoName:  "Age",
						Type:    reflect.TypeOf(int8(0)),
						Offset:  24,
					},
					"LastName": {
						ColName: "last_name",
						GoName:  "LastName",
						Type:    reflect.TypeOf(""),
						Offset:  32,
					},
				},
			},
		},
	}
	r := NewRegistry()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.parseModel(tt.args.val)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			for _, v := range got.FieldMap {
				t.Logf("%s: %+v", v.ColName, v.Offset)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_underscoreName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "upper cases",
			args: args{
				name: "ID",
			},
			want: "i_d",
		},
		{
			name: "use number",
			args: args{
				name: "Table1Name",
			},
			want: "table1_name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := underscoreName(tt.args.name)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_registry_Get(t *testing.T) {
	type args struct {
		val any
	}
	tests := []struct {
		name    string
		r       *registry
		args    args
		want    *Model
		wantErr error
	}{
		{
			name: "test model",
			r:    NewRegistry(),

			args: args{
				val: &TestModel{},
			},
			want: &Model{
				TableName: "test_model",
				FieldMap: map[string]*Column{
					"Id": {
						ColName: "id",
					},
					"FirstName": {
						ColName: "first_name",
					},
					"Age": {
						ColName: "age",
					},
					"LastName": {
						ColName: "last_name",
					},
				},
			},
		},
		{
			name: "test tag",
			r:    NewRegistry(),

			args: args{
				val: func() any {
					type TagTable struct {
						FirstName string `orm:"column=first_name_t"`
					}
					return &TagTable{}
				}(),
			},
			want: &Model{
				TableName: "tag_table",
				FieldMap: map[string]*Column{
					"FirstName": {
						ColName: "first_name_t",
						GoName:  "FirstName",
						Type:    reflect.TypeOf(""),
						Offset:  0,
					},
				},
			},
		},

		{
			name: "empty column",
			r:    NewRegistry(),

			args: args{
				val: func() any {
					type TagTable struct {
						FirstName string `orm:"column="`
					}
					return &TagTable{}
				}(),
			},
			want: &Model{
				TableName: "tag_table",
				FieldMap: map[string]*Column{
					"FirstName": {
						ColName: "first_name",
						GoName:  "FirstName",
						Type:    reflect.TypeOf(""),
						Offset:  0,
					},
				},
			},
		},
		{
			name: "only column",
			r:    NewRegistry(),

			args: args{
				val: func() any {
					type TagTable struct {
						FirstName string `orm:"column"`
					}
					return &TagTable{}
				}(),
			},
			wantErr: errs.NewErrInvalidTagContent("column"),
		},
		{
			name: "ignore tag",
			r:    NewRegistry(),

			args: args{
				val: func() any {
					type TagTable struct {
						FirstName string `orm:"abc=acd"`
					}
					return &TagTable{}
				}(),
			},
			want: &Model{
				TableName: "tag_table",
				FieldMap: map[string]*Column{
					"FirstName": {
						ColName: "first_name",
						GoName:  "FirstName",
						Type:    reflect.TypeOf(""),
						Offset:  0,
					},
				},
			},
		},
		{
			name: "table name",
			r:    NewRegistry(),

			args: args{
				val: &CustomTableName{},
			},
			want: &Model{
				TableName: "custom_table_name_t",
				FieldMap: map[string]*Column{
					"Name": {
						ColName: "name",
						GoName:  "Name",
						Type:    reflect.TypeOf(""),
						Offset:  0,
					},
				},
			},
		},
		{
			name: "ptr table name",
			r:    NewRegistry(),

			args: args{
				val: &CustomTableNamePtr{},
			},
			want: &Model{
				TableName: "custom_table_name_ptr_t",
				FieldMap: map[string]*Column{
					"Name": {
						ColName: "name",
						GoName:  "Name",
						Type:    reflect.TypeOf(""),
						Offset:  0,
					},
				},
			},
		},
		{
			name: "empty table name",
			r:    NewRegistry(),

			args: args{
				val: &EmptyTableName{},
			},
			want: &Model{
				TableName: "empty_table_name",
				FieldMap: map[string]*Column{
					"Name": {
						ColName: "name",
						GoName:  "Name",
						Type:    reflect.TypeOf(""),
						Offset:  0,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Get(tt.args.val)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, got)

			typ := reflect.TypeOf(tt.args.val)
			m, ok := tt.r.models.Load(typ)
			assert.True(t, ok)
			assert.Equal(t, tt.want, m)
		})
	}
}

type CustomTableName struct {
	Name string
}

func (c CustomTableName) TableName() string {
	return "custom_table_name_t"
}

type CustomTableNamePtr struct {
	Name string
}

func (c *CustomTableNamePtr) TableName() string {
	return "custom_table_name_ptr_t"
}

type EmptyTableName struct {
	Name string
}

func (c *EmptyTableName) TableName() string {
	return ""
}

func TestModelWithTableName(t *testing.T) {

	r := NewRegistry()
	m, err := r.Register(&TestModel{}, ModelWithTableName("test_table_name_t"))
	require.NoError(t, err)
	assert.Equal(t, "test_table_name_t", m.TableName)
}

func TestModelWithColumnName(t *testing.T) {
	type args struct {
		obj     any
		field   string
		colName string
	}
	tests := []struct {
		name        string
		args        args
		wantColName string
		wantErr     error
	}{
		{
			name: "emtpy new name",
			args: args{
				obj:     &TestModel{},
				field:   "FirstName",
				colName: "",
			},
			wantColName: "",
		},
		{
			name: "invalid field name",
			args: args{
				obj:     &TestModel{},
				field:   "FirstNameXXXX",
				colName: "first_name",
			},
			wantErr: errs.NewErrUnknownField("FirstNameXXXX"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRegistry()
			got, err := r.Register(tt.args.obj, ModelWithColumnName(tt.args.field, tt.args.colName))
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}

			fd, ok := got.FieldMap[tt.args.field]
			require.True(t, ok)
			assert.Equal(t, tt.wantColName, fd.ColName)
		})
	}
}

type TestModel struct {
	Id        int64
	FirstName string
	Age       int8
	LastName  *sql.NullString
}
