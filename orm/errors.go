package orm

import "go-actions/orm/internal/errs"

var (
	// ErrNoRows is returned when no rows in the database matches the query.
	ErrNoRows = errs.ErrNoRows
)
