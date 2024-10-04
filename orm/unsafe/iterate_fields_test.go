package unsafe

import "testing"

func TestPrintFieldOffset(t *testing.T) {
	type args struct {
		entity any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "user",
			args: args{
				entity: User{
					Name: "test",
					Age:  10,
				},
			},
		},
		{
			name: "userV1",
			args: args{
				entity: UserV1{
					Name: "test",
					Age:  10,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintFieldOffset(tt.args.entity)
		})
	}
}

type User struct {
	Name string
	Age  int32

	Alias   []string
	Address string
}

type UserV1 struct {
	Name string
	Age  int32
	Age1 int32

	Alias   []string
	Address string
}
