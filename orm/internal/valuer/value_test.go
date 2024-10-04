package valuer

import (
	"database/sql/driver"
	"go-actions/orm/meta"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func BenchmarkSetColumns(b *testing.B) {
	b.Run("reflect", func(b *testing.B) {
		benchmarkSetColumns(b, NewReflectValue)
	})

	b.Run("unsafe", func(b *testing.B) {
		benchmarkSetColumns(b, NewUnsafeValue)
	})
}

func benchmarkSetColumns(b *testing.B, creator Creator) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(b, err)
	defer mockDB.Close()

	mockRows := sqlmock.NewRows([]string{"name", "age"})

	row := []driver.Value{"tom", 10}
	for i := 0; i < b.N; i++ {
		mockRows.AddRow(row...)
	}
	mock.ExpectQuery("SELECT").WillReturnRows(mockRows)
	rows, err := mockDB.Query("SELECT")
	require.NoError(b, err)

	r := meta.NewRegistry()
	m, err := r.Get(&TestModel{})
	require.NoError(b, err)

	val := creator(m, &TestModel{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rows.Next()
		val.SetColumns(rows)
	}
}
