package orm

import (
	"context"
	"database/sql"
)

type Deleter[T any] struct {
	db *sql.DB
}

func NewDeleter[T any]() *Deleter[T] {
	return &Deleter[T]{}
}

func (d *Deleter[T]) From(table string) *Deleter[T] {
	return d
}

func (d *Deleter[T]) Where(condition ...*Condition) *Deleter[T] {
	return d
}

func (d *Deleter[T]) Build() (*Query, error) {
	return nil, nil
}

func (d *Deleter[T]) Exec(ctx context.Context) (sql.Result, error) {
	q, err := d.Build()
	if err != nil {
		return nil, err
	}
	return d.db.ExecContext(ctx, q.SQL, q.Args...)
}
