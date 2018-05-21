package panda

import (
	"database/sql"

	"github.com/pkg/errors"
)

// Select executes select statement
func (q *Query) Select() (*sql.Rows, error) {
	var rows *sql.Rows
	err := q.timeFunc("select", func() error {
		sql, args, err := q.ToSQL()
		if err != nil {
			return errors.WithStack(err)
		}
		Log(sql, args...)
		rows, err = q.Conn.Query(sql, args...)

		return nil
	})

	return rows, err
}

// SelectOne executes select statement
func (q *Query) SelectOne() (*sql.Row, error) {
	var row *sql.Row
	err := q.timeFunc("SelectOne", func() error {
		sql, args, err := q.ToSQL()
		if err != nil {
			return errors.WithStack(err)
		}
		Log(sql, args...)
		row = q.Conn.QueryRow(sql, args...)

		return nil
	})
	return row, err
}
