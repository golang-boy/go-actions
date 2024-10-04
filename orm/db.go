package orm

import (
	"database/sql"
	"go-actions/orm/internal/valuer"
	"go-actions/orm/meta"
)

type DBOption func(db *DB)

type DB struct {
	// DB is the database connection
	r       meta.Registry
	db      *sql.DB
	Creator valuer.Creator
}

func Open(driver, sourceName string, opts ...DBOption) (*DB, error) {
	db, err := sql.Open(driver, sourceName)
	if err != nil {
		return nil, err
	}
	return OpenDB(db, opts...)
}

func UseReflectValuer() DBOption {
	return func(db *DB) {
		db.Creator = valuer.NewReflectValue
	}
}

func OpenDB(db *sql.DB, opts ...DBOption) (*DB, error) {
	res := &DB{
		db:      db,
		r:       meta.NewRegistry(),
		Creator: valuer.NewReflectValue,
	}

	for _, opt := range opts {
		opt(res)
	}

	return res, nil
}
