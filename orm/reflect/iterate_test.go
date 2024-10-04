package reflect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterateArray(t *testing.T) {
	type args struct {
		entity any
	}
	tests := []struct {
		name    string
		args    args
		want    []any
		wantErr error
	}{
		{
			name: "array",
			args: args{
				entity: [3]string{"a", "b", "c"},
			},
			want: []any{"a", "b", "c"},
		},
		{
			name: "slice",
			args: args{
				entity: []string{"a", "b", "c"},
			},
			want: []any{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IterateArrayOrSlice(tt.args.entity)
			assert.Equal(t, err, tt.wantErr)
			if err != nil {
				return
			}
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestIterateMap(t *testing.T) {
	type args struct {
		entity any
	}
	tests := []struct {
		name    string
		args    args
		want    []any
		want1   []any
		wantErr error
	}{
		{
			name: "map",
			args: args{
				entity: map[string]string{
					"a": "1",
					"b": "2",
					"c": "3"},
			},
			want:  []any{"a", "b", "c"},
			want1: []any{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := IterateMap(tt.args.entity)
			assert.Equal(t, err, tt.wantErr)
			if err != nil {
				return
			}

			assert.EqualValues(t, got, tt.want)
			assert.EqualValues(t, got1, tt.want1)
		})
	}
}
