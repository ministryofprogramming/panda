package panda

import (
	"fmt"
	"regexp"
	"strings"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
)

// Query is used to build and execute SQL query on Connection
type Query struct {
	Conn          *Connection
	RawSQL        Clause
	columnClauses ColumnClauses
	whereClause   Clause
	fromClause    FromClause
	orderClause   Clause
	joinClauses   JoinClauses
	groupClause   GroupClause
	havingClause  HavingClause
	limitResults  int
	Elapsed       int64
}

// NewQuery will create a new "empty" query from the current connection.
func NewQuery(c *Connection) *Query {
	return &Query{
		Conn: c,
	}
}

// inRegex is regular expression for detecting sql IN statement
var inRegex = regexp.MustCompile(`(?i)in\s*\(\s*\?\s*\)`)

// RawQuery will override the query building feature and will use
// whatever query you want to execute against the `Connection`. You can continue
// to use the `?` argument syntax.
//
//	c.RawQuery("select * from foo where id = ?", 1)
func (c *Connection) RawQuery(stmt string, args ...interface{}) *Query {
	return NewQuery(c).RawQuery(stmt, args)
}

// RawQuery will override the query building feature and will use
// whatever query you want to execute against the `Connection`. You can continue
// to use the `?` argument syntax.
//
//	q.RawQuery("select * from foo where id = ?", 1)
func (q *Query) RawQuery(stmt string, args ...interface{}) *Query {
	q.RawSQL = Clause{stmt, args}
	return q
}

// Columns will add Columns for SELECT clause to the query.
//
// 	c.Select("id,name")
//  c.Select("u.id, u.name")
func (c *Connection) Columns(stmt ...string) *Query {
	return NewQuery(c).Columns(stmt...)
}

// Columns will add Columns for SELECT clause to the query.
//
// 	q.Select("id,name")
//  q.Select("u.id, u.name")
func (q *Query) Columns(columns ...string) *Query {
	for _, column := range columns {
		q.columnClauses = append(q.columnClauses, ColumnClause{column})
	}
	return q
}

// From will add from clause to the query.
//
// 	c.From("Users","")
func (c *Connection) From(from string) *Query {
	return NewQuery(c).From(from)
}

// From will add from clause to the query.
//
// 	q.From("Users","")
func (q *Query) From(from string) *Query {
	q.fromClause = FromClause{from}
	return q
}

// Where will add where clause to the query. You may use `?` in place of
// arguments.
//
// 	c.Where("id = ?", 1)
func (c *Connection) Where(stmt string, args ...interface{}) *Query {
	return NewQuery(c).Where(stmt, args)
}

// Where will add where clause to the query. You may use `?` in place of
// arguments.
//
// 	q.Where("id in (?)", 1, 2, 3)
func (q *Query) Where(stmt string, args ...interface{}) *Query {
	if q.RawSQL.Statement != "" {
		fmt.Println("Warning: Query is setup to use raw SQL")
		return q
	}

	if inRegex.MatchString(stmt) {
		var vals []string
		for i := 0; i < len(args); i++ {
			vals = append(vals, "?")
		}
		qry := fmt.Sprintf("(%s)", strings.Join(vals, ","))
		stmt = strings.Replace(stmt, "(?)", qry, 1)
	}

	q.whereClause = Clause{stmt, args}
	return q
}

// Order will append an order clause to the query.
//
// 	c.Order("name desc")
func (c *Connection) Order(stmt string) *Query {
	return NewQuery(c).Order(stmt)
}

// Order will append an order clause to the query.
//
// 	q.Order("name desc")
func (q *Query) Order(stmt string) *Query {
	if q.RawSQL.Statement != "" {
		fmt.Println("Warning: Query is setup to use raw SQL")
		return q
	}

	q.orderClause = Clause{stmt, []interface{}{}}
	return nil
}

// Limit will add a limit clause to the query.
func (c *Connection) Limit(limit int) *Query {
	return NewQuery(c).Limit(limit)
}

// Limit will add a limit clause to the query.
func (q *Query) Limit(limit int) *Query {
	q.limitResults = limit
	return q
}

// GroupBy will append a GROUP BY clause to the query
func (c *Connection) GroupBy(field string) *Query {
	return NewQuery(c).GroupBy(field)
}

// GroupBy will append a GROUP BY clause to the query
func (q *Query) GroupBy(field string) *Query {
	if q.RawSQL.Statement != "" {
		fmt.Println("Warning: Query is setup to use raw SQL")
		return q
	}

	q.groupClause = GroupClause{field}
	return q
}

// Having will append a HAVING clause to the query
func (c *Connection) Having(condition string, args ...interface{}) *Query {
	return NewQuery(c).Having(condition, args)
}

// Having will append a HAVING clause to the query
func (q *Query) Having(condition string, args ...interface{}) *Query {
	if q.RawSQL.Statement != "" {
		fmt.Println("Warning: Query is setup to use raw SQL")
		return q
	}
	q.havingClause = HavingClause{condition, args}
	return q
}

// Join will append a JOIN clause to the query
func (c *Connection) Join(table string, on string, args ...interface{}) *Query {
	return NewQuery(c).Join(table, on, args)
}

// Join will append a JOIN clause to the query
func (q *Query) Join(table string, on string, args ...interface{}) *Query {
	if q.RawSQL.Statement != "" {
		fmt.Println("Warning: Query is setup to use raw SQL")
		return q
	}
	q.joinClauses = append(q.joinClauses, JoinClause{"JOIN", table, on, args})
	return q
}

// LeftJoin will append a LEFT JOIN clause to the query
func (c *Connection) LeftJoin(table string, on string, args ...interface{}) *Query {
	return NewQuery(c).LeftJoin(table, on, args)
}

// LeftJoin will append a LEFT JOIN clause to the query
func (q *Query) LeftJoin(table string, on string, args ...interface{}) *Query {
	if q.RawSQL.Statement != "" {
		fmt.Println("Warning: Query is setup to use raw SQL")
		return q
	}
	q.joinClauses = append(q.joinClauses, JoinClause{"LEFT JOIN", table, on, args})
	return q
}

// RightJoin will append a RIGHT JOIN clause to the query
func (c *Connection) RightJoin(table string, on string, args ...interface{}) *Query {
	return NewQuery(c).RightJoin(table, on, args)
}

// RightJoin will append a RIGHT JOIN clause to the query
func (q *Query) RightJoin(table string, on string, args ...interface{}) *Query {
	if q.RawSQL.Statement != "" {
		fmt.Println("Warning: Query is setup to use raw SQL")
		return q
	}
	q.joinClauses = append(q.joinClauses, JoinClause{"RIGHT JOIN", table, on, args})
	return q
}

// LeftOuterJoin will append a LEFT OUTER JOIN clause to the query
func (c *Connection) LeftOuterJoin(table string, on string, args ...interface{}) *Query {
	return NewQuery(c).LeftOuterJoin(table, on, args)
}

// LeftOuterJoin will append a LEFT OUTER JOIN clause to the query
func (q *Query) LeftOuterJoin(table string, on string, args ...interface{}) *Query {
	if q.RawSQL.Statement != "" {
		fmt.Println("Warning: Query is setup to use raw SQL")
		return q
	}
	q.joinClauses = append(q.joinClauses, JoinClause{"LEFT OUTER JOIN", table, on, args})
	return q
}

// RightOuterJoin will append a RIGHT OUTER JOIN clause to the query
func (c *Connection) RightOuterJoin(table string, on string, args ...interface{}) *Query {
	return NewQuery(c).RightOuterJoin(table, on, args)
}

// RightOuterJoin will append a RIGHT OUTER JOIN clause to the query
func (q *Query) RightOuterJoin(table string, on string, args ...interface{}) *Query {
	if q.RawSQL.Statement != "" {
		fmt.Println("Warning: Query is setup to use raw SQL")
		return q
	}
	q.joinClauses = append(q.joinClauses, JoinClause{"RIGHT OUTER JOIN", table, on, args})
	return q
}

// LeftInnerJoin will append a LEFT INNER JOIN clause to the query
func (c *Connection) LeftInnerJoin(table string, on string, args ...interface{}) *Query {
	return NewQuery(c).LeftInnerJoin(table, on, args)
}

// LeftInnerJoin will append a LEFT INNER JOIN clause to the query
func (q *Query) LeftInnerJoin(table string, on string, args ...interface{}) *Query {
	if q.RawSQL.Statement != "" {
		fmt.Println("Warning: Query is setup to use raw SQL")
		return q
	}
	q.joinClauses = append(q.joinClauses, JoinClause{"LEFT INNER JOIN", table, on, args})
	return q
}

// RightInnerJoin will append a RIGHT INNER JOIN clause to the query
func (c *Connection) RightInnerJoin(table string, on string, args ...interface{}) *Query {
	return NewQuery(c).RightInnerJoin(table, on, args)
}

// RightInnerJoin will append a RIGHT INNER JOIN clause to the query
func (q *Query) RightInnerJoin(table string, on string, args ...interface{}) *Query {
	if q.RawSQL.Statement != "" {
		fmt.Println("Warning: Query is setup to use raw SQL")
		return q
	}
	q.joinClauses = append(q.joinClauses, JoinClause{"RIGHT INNER JOIN", table, on, args})
	return q
}

//ToSQL generate SQL and the appropriate arguments for DB execution
func (q *Query) ToSQL() (string, []interface{}, error) {
	sb := q.Conn.Dialect().GetSQLBuilder(*q)

	sql, args := sb.Build()

	return sql, args, nil
}

func (q *Query) timeFunc(name string, fn func() error) error {
	now := time.Now()
	err := fn()
	atomic.AddInt64(&q.Elapsed, int64(time.Now().Sub(now)))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
