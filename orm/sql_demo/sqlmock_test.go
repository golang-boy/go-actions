package sqldemo

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestSQLMock(t *testing.T) {

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mockrows := sqlmock.NewRows([]string{"id", "first_name"})
	mockrows.AddRow(1, "test")
	mock.ExpectQuery("SELECT id, first_name FROM users WHERE id = $1").WithArgs(1).WillReturnRows(mockrows)

	rows, err := db.Query("SELECT id, first_name FROM users WHERE id = $1", 1)
	require.NoError(t, err)

	for rows.Next() {
		tm := &TestModel{}
		err := rows.Scan(&tm.Id, &tm.FirstName)
		require.NoError(t, err)
		//fmt.Println(id, firstName)
	}
}

type TestModel struct {
	Id        int64
	FirstName string
	Age       int8
	LastName  *sql.NullString
}
