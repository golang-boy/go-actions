package reflect

import (
	"go-actions/orm/reflect/types"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterateFunc(t *testing.T) {
	type args struct {
		entity any
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]FuncInfo
		wantErr error
	}{
		{
			name: "struct", // 结构体输入只能访问结构体对应的方法，如GetAge函数，不能访问到ChangeName函数
			args: args{
				entity: types.NewUser("tom", 12),
			},
			want: map[string]FuncInfo{
				"GetAge": {
					Name:       "GetAge",
					InputTypes: []reflect.Type{reflect.TypeOf(types.User{})}, // 输入参数的第一个永远是对象自己
					OutputTypes: []reflect.Type{
						reflect.TypeOf(0),
					},
					Result: []any{12},
				},
			},
		},
		{
			name: "pointer", // 指针方式输入能访问所有方法
			args: args{
				entity: types.NewUserPtr("tom", 12),
			},
			want: map[string]FuncInfo{
				"GetAge": {
					Name:       "GetAge",
					InputTypes: []reflect.Type{reflect.TypeOf(&types.User{})},
					OutputTypes: []reflect.Type{
						reflect.TypeOf(0),
					},
					Result: []any{12},
				},
				"ChangeName": {
					Name:        "ChangeName",
					InputTypes:  []reflect.Type{reflect.TypeOf(&types.User{}), reflect.TypeOf("")},
					OutputTypes: []reflect.Type{},
					Result:      []any{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IterateFunc(tt.args.entity)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
