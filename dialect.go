package panda

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

// Dialect interface contains behaviours that differ acress SQL database
type Dialect interface {
	// GetName - gets dialec't name
	GetName() string

	// SetDB - sets DB for dialect
	SetDB(db *sql.DB)

	// BindVar return the placeholder for actual values in SQL statements, in many dbs it is "?", Postgres using $1
	BindVar(i int) string

	// Quote - quotes field name to avoid SQL parsing excwptions by using reserved words as a field name
	Quote(key string) string

	// HasIndex check has index or not
	HasIndex(tableName string, indexName string) bool

	// HasForeignKey - check has foreign key or not
	HasForeignKey(tableName string, foreignKeyName string) bool

	// Remove index
	RemoveIndex(tableName string, indexName string) error

	// HasTable - check has table or not
	HasTable(tableName string) bool

	// HasColumn - check has column or not
	HasColumn(tableName string, columnName string) bool

	// ModifyColumn - modify column's type
	ModifyColumn(tableName string, columnName string, typ string) error

	// LimitAndOffsetSQL - returns generated SQL with Limit and Offset
	LimitAndOffsetSQL(limit, offset interface{}) string

	// SelectFromDummyTable return select values, for most dbs, `SELECT values` just works, mysql needs `SELECT value FROM DUAL`
	SelectFromDummyTable() string

	// LastInsertIDReturningSuffix - most dbs support LastInsertId, butp postgress needs to use 'RETURNING'
	LastInsertIDReturningSuffix(tableName string, columnName string) string

	// DefaultValueStr
	DefaultValueStr() string

	// BuildKeyName returns a valid key name (foreign key, index key) for the given table, field and reference
	BuildKeyName(kind, tableName string, fields ...string) string

	// CurrentDatabase - returns current database name
	CurrentDatabase() string

	//BuildQuerySQL creates SQL query from query object
	GetSQLBuilder(q Query) SQLBuilder
}

var dialectsMap = map[string]Dialect{}

func newDialect(name string, db *sql.DB) (Dialect, error) {
	if val, ok := dialectsMap[name]; ok {
		dialect := reflect.New(reflect.TypeOf(val).Elem()).Interface().(Dialect)
		dialect.SetDB(db)
		return dialect, nil
	}

	err := fmt.Errorf("'%s' dialect is not supported", name)
	return nil, err
}

// RegisterDialect - registers new dialect
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func currentDatabaseAndTable(dialect Dialect, tableName string) (string, string) {
	if strings.Contains(tableName, ".") {
		splitStrings := strings.SplitN(tableName, ".", 2)
		return splitStrings[0], splitStrings[1]
	}
	return dialect.CurrentDatabase(), tableName
}
