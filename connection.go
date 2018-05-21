package panda

import (
	"database/sql"
	"errors"
)

// Connection struct describes current db connection
type Connection struct {
	*sql.DB
	dialect Dialect
}

// Connect creates new Database connection
func Connect(dialect, dataSource string) (*Connection, error) {
	if len(dataSource) == 0 {
		return nil, errors.New("No data source")
	}

	dbConn, err := sql.Open(dialect, dataSource)
	if err != nil {
		return nil, err
	}

	if err := dbConn.Ping(); err != nil {
		dbConn.Close()
		return nil, err
	}

	dialectObj, err := newDialect(dialect, dbConn)
	if err != nil {
		return nil, err
	}

	conn := &Connection{
		DB:      dbConn,
		dialect: dialectObj,
	}

	return conn, nil

}

// Dialect returns Database Dialect
func (d *Connection) Dialect() Dialect {
	return d.dialect
}
