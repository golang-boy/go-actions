package sqldemo

import (
	"context"
	"database/sql"

	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestDB(t *testing.T) {

	db, err := sql.Open("sqlite3", "file:test.db?cache=shared&mode=memory")
	require.NoError(t, err)
	defer db.Close()
	db.Ping()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	table_ddl := `
CREATE TABLE IF NOT EXISTS test_model(
    id INTEGER PRIMARY KEY,
    first_name TEXT NOT NULL,
    age INTEGER,
    last_name TEXT NOT NULL
)
`
	_, err = db.ExecContext(ctx, table_ddl)
	cancel()
	require.NoError(t, err)

	row := db.QueryRowContext(context.Background(), "SELECT `id`, `first_name`,`age`, `last_name` FROM `test_model` LIMIT 1")
	if row.Err() != nil {
		t.Fatal(row.Err())
	}
}
