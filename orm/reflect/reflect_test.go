package reflect

import (
	"go-actions/orm/reflect/user"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIterateFields(t *testing.T) {
	type args struct {
		entity any
	}

	type User struct {
		Name string `json:"name"`
		age  int    `json:"age"`
	}

	tests := []struct {
		name    string
		args    args
		want    map[string]any
		wantErr error
	}{
		{
			name: "test",
			args: args{
				entity: User{
					Name: "tom",
					age:  18,
				},
			},
			want: map[string]any{
				"Name": "tom",
				"age":  0,
			},
		},

		{
			name: "多重指针",
			args: args{
				entity: func() **User {
					res := &User{Name: "tom",
						age: 18,
					}
					return &res
				}(),
			},
			want: map[string]any{
				"Name": "tom",
				"age":  0,
			},
		},

		{
			name: "nil",
			args: args{
				entity: nil,
			},
			wantErr: ErrNotSupportNilType,
		},
		{
			name: "(*User)(nil)",
			args: args{
				entity: (*User)(nil),
			},
			wantErr: ErrNotSupportZeroVal,
		},
		{
			name: "基础类型",
			args: args{
				entity: 18,
			},
			wantErr: ErrMustBeStruct,
		},
		{
			name: "other pkg user",
			args: args{
				entity: user.User{
					Name: "tom",
				},
			},
			want: map[string]any{
				"Name": "tom",
				"age":  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IterateFields(tt.args.entity)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSetField(t *testing.T) {
	type args struct {
		entity   any
		field    string
		newValue any
	}

	type User struct {
		Name string `json:"name"`
		age  int    `json:"age"`
	}

	tests := []struct {
		name       string
		args       args
		wantErr    error
		wantEntity any
	}{
		{
			name: "struct",
			args: args{
				entity: User{
					Name: "tom",
					age:  18,
				},
				field:    "Name",
				newValue: "jack",
			},
			wantErr: ErrCantSetField,
		},
		{
			name: "pointer",
			args: args{
				entity: &User{
					Name: "tom",
				},
				field:    "Name",
				newValue: "jack",
			},
			wantEntity: &User{
				Name: "jack",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetField(tt.args.entity, tt.args.field, tt.args.newValue)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantEntity, tt.args.entity)
		})
	}
}

func TestDefineStruct(t *testing.T) {
	// 定义测试用例
	timeNow := time.Now()
	tests := []struct {
		metas  []MetaJson
		expect interface{}
	}{
		{
			metas: []MetaJson{
				{Name: "Name", Type: "string", Tag: `json:"name"`, Value: "Alice"},
				{Name: "Age", Type: "int", Tag: `json:"age"`, Value: 30},
				{Name: "Height", Type: "float64", Tag: `json:"height"`, Value: 5.9},
				{Name: "IsEmployed", Type: "bool", Tag: `json:"is_employed"`, Value: true},
				{Name: "Birthdate", Type: "time.Time", Tag: `json:"birthdate"`, Value: timeNow},
			},
			expect: &struct {
				Name       string    `json:"name"`
				Age        int       `json:"age"`
				Height     float64   `json:"height"`
				IsEmployed bool      `json:"is_employed"`
				Birthdate  time.Time `json:"birthdate"`
			}{"Alice", 30, 5.9, true, timeNow}, // Birthdate will be set later
		},
	}

	for _, test := range tests {
		result := DefineStruct(test.metas)
		assert.Equal(t, test.expect, result)

	}
}
