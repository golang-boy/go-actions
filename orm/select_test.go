package orm

import (
	"context"
	"database/sql"
	"errors"
	"go-actions/orm/internal/errs"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/glebarez/go-sqlite"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelector_Build(t *testing.T) {

	db := memDB(t)

	testCases := []struct {
		name    string
		s       QueryBuilder
		want    *Query
		wantErr error
	}{
		{
			name: "from",
			s:    NewSelector[TestModel](db).From("test_model_tab"),
			want: &Query{
				SQL: "SELECT * FROM test_model_tab;",
			},
		},
		{
			name: "no from",
			s:    NewSelector[TestModel](db),
			want: &Query{
				SQL: "SELECT * FROM `test_model`;",
			},
		},
		{
			name: "with db",
			s:    NewSelector[TestModel](db).From("`test_db`.`test_model`"),
			want: &Query{
				SQL: "SELECT * FROM `test_db`.`test_model`;",
			},
		},
		{
			name: "where",
			s:    NewSelector[TestModel](db).Where(F("Age").Eq(18)),
			want: &Query{
				SQL:  "SELECT * FROM `test_model` WHERE `age` = ?;",
				Args: []any{18},
			},
		},
		// {
		// 	// 单一简单条件
		// 	name: "single and simple predicate",
		// 	s: NewSelector[TestModel]().From("`test_model_t`").
		// 		Where(F("Id").Eq(1)),
		// 	want: &Query{
		// 		SQL:  "SELECT * FROM `test_model_t` WHERE `id` = ?;",
		// 		Args: []any{1},
		// 	},
		// },
		// {
		// 	// 多个 predicate
		// 	name: "multiple predicates",
		// 	s: NewSelector[TestModel]().
		// 		Where(F("Age").GT(18), F("Age").LT(35)),
		// 	want: &Query{
		// 		// TestModel -> test_model
		// 		SQL:  "SELECT * FROM `test_model` WHERE (`age` > ?) AND (`age` < ?);",
		// 		Args: []any{18, 35},
		// 	},
		// },
		{
			// 使用 AND
			name: "and",
			s: NewSelector[TestModel](db).
				Where(F("Age").GT(18).And(F("Age").LT(35))),
			want: &Query{
				SQL:  "SELECT * FROM `test_model` WHERE (`age` > ?) AND (`age` < ?);",
				Args: []any{18, 35},
			},
		},
		{
			// 使用 OR
			name: "or",
			s: NewSelector[TestModel](db).
				Where(F("Age").GT(18).Or(F("Age").LT(35))),
			want: &Query{
				SQL:  "SELECT * FROM `test_model` WHERE (`age` > ?) OR (`age` < ?);",
				Args: []any{18, 35},
			},
		},
		{
			// 使用 NOT
			name: "not",
			s:    NewSelector[TestModel](db).Where(Not(F("Age").GT(18))),
			want: &Query{
				// NOT 前面有两个空格，因为我们没有对 NOT 进行特殊处理
				SQL:  "SELECT * FROM `test_model` WHERE  NOT (`age` > ?);",
				Args: []any{18},
			},
		},
	}

	for _, tc := range testCases {

		q, err := tc.s.Build()
		assert.Equal(t, tc.wantErr, err)

		if err != nil {
			return
		}
		assert.Equal(t, tc.want, q)

	}
}

type TestModel struct {
	Id        int64
	FirstName string
	Age       int8
	LastName  *sql.NullString
}

func memDB(t *testing.T) *DB {
	db, err := Open("sqlite", ":memory:")
	require.NoError(t, err)
	return db
}

func TestSelector_Get(t *testing.T) {

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	db, err := OpenDB(mockDB)
	require.NoError(t, err)

	mock.ExpectQuery("SELECT .* ").WillReturnError(errors.New("query error"))

	rows := sqlmock.NewRows([]string{"id", "first_name", "age", "last_name"})
	mock.ExpectQuery("SELECT .* WHERE `age` > .*").WillReturnRows(rows)

	rows = sqlmock.NewRows([]string{"id", "first_name", "age", "last_name"})
	rows.AddRow(1, "Tom", 19, "Jerry")
	mock.ExpectQuery("SELECT .* WHERE `age` < .*").WillReturnRows(rows)

	testCase := []struct {
		name    string
		s       *Selector[TestModel]
		want    *TestModel
		wantErr error
	}{
		{
			name:    "invalid query",
			s:       NewSelector[TestModel](db).From("test_model").Where(F("AgeXXX").Eq(18)),
			wantErr: errs.NewErrUnknownField("AgeXXX"),
		},
		{
			name:    "query error",
			s:       NewSelector[TestModel](db).From("test_model").Where(F("Age").GT(18).And(F("Age").LT(35))),
			wantErr: errors.New("query error"),
		},
		{
			name:    "no rows",
			s:       NewSelector[TestModel](db).Where(F("Age").GT(18)),
			wantErr: ErrNoRows,
		},
		{
			name: "data",
			s:    NewSelector[TestModel](db).Where(F("Age").LT(18)),
			want: &TestModel{
				Id:        1,
				FirstName: "Tom",
				Age:       19,
				LastName: &sql.NullString{
					Valid:  true,
					String: "Jerry",
				},
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.s.Get(context.Background())
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
