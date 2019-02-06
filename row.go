package ctxdb

import (
	"database/sql"
	"unsafe"
)

type row struct {
	err  error
	rows *sql.Rows
}

func (r *row) intoDBRow() *sql.Row {
	return (*sql.Row)(unsafe.Pointer(r))
}
