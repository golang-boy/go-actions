package orm

import (
	"context"
	"database/sql"
)

type Updater[T any] struct {
	db *sql.DB
}

func NewUpdater[T any]() *Updater[T] {
	return &Updater[T]{}
}

func (u *Updater[T]) Build() (*Query, error) {
	return nil, nil
}

func (u *Updater[T]) Exec(ctx context.Context) (sql.Result, error) {
	q, err := u.Build()
	if err != nil {
		return nil, err
	}
	return u.db.ExecContext(ctx, q.SQL, q.Args...)
}
