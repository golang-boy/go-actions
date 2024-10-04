package valuer

import (
	"database/sql"
	"go-actions/orm/meta"
)

type Valuer interface {
	SetColumns(rows *sql.Rows) error
}

type Creator func(model *meta.Model, entity any) Valuer
