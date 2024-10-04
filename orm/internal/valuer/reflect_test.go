package valuer

import (
	"go-actions/orm/meta"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_reflectValue_SetColumns(t *testing.T) {
	testSetColumns(t, NewReflectValue)
}

func Test_unsafeValue_SetColumns(t *testing.T) {
	testSetColumns(t, NewUnsafeValue)
}

func testSetColumns(t *testing.T, creator Creator) {
	type args struct {
		entity any
		rows   func() *sqlmock.Rows
	}
	tests := []struct {
		name       string
		r          reflectValue
		args       args
		wantErr    error
		wantEntity any
	}{
		{
			name: "set colums",
			args: args{
				entity: &TestModel{},
				rows: func() *sqlmock.Rows {
					rows := sqlmock.NewRows([]string{"name", "age"})
					rows.AddRow("tom", 10)
					return rows
				},
			},
			wantEntity: &TestModel{Name: "tom", Age: 10},
		},
	}

	r := meta.NewRegistry()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockRows := tt.args.rows()
			mock.ExpectQuery("SELECT").WillReturnRows(mockRows)
			rows, err := mockDB.Query("SELECT")
			require.NoError(t, err)

			rows.Next()

			m, err := r.Get(tt.args.entity)
			require.NoError(t, err)
			val := creator(m, tt.args.entity)
			err = val.SetColumns(rows)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantEntity, tt.args.entity)
		})
	}
}

type TestModel struct {
	Name string
	Age  int
}
