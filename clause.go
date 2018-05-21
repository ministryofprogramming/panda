package panda

import (
	"fmt"
	"strings"
)

// Clause  hold SQL clause data
type Clause struct {
	Statement string
	Arguments []interface{}
}

func (c Clause) String() string {
	return c.Statement
}

// ColumnClause holds column data
type ColumnClause struct {
	Column string
}

// ColumnClauses holds  the array of column data
type ColumnClauses []ColumnClause

func (c ColumnClause) String() string {
	return c.Column
}

func (c ColumnClauses) String() string {
	cs := []string{}
	for _, col := range c {
		cs = append(cs, col.String())
	}

	return strings.Join(cs, ",")
}

// FromClause  Holds FROM clause data
type FromClause struct {
	From string
}

func (c FromClause) String() string {
	return c.From
}

// JoinClause holds the field to apply the JOIN clause
type JoinClause struct {
	JoinType  string
	Table     string
	On        string
	Arguments []interface{}
}

// JoinClauses holds the array of JOIN clauses
type JoinClauses []JoinClause

func (c JoinClause) String() string {
	sql := fmt.Sprintf("%s %s", c.JoinType, c.Table)

	if len(c.On) > 0 {
		sql += " ON " + c.On
	}

	return sql
}

func (c JoinClauses) String() string {
	cs := []string{}
	for _, cl := range c {
		cs = append(cs, cl.String())
	}
	return strings.Join(cs, " ")
}

// GroupClause holds the field to apply the GROUP clause on
type GroupClause struct {
	Field string
}

func (c GroupClause) String() string {
	return c.Field
}

// HavingClause defines a condition and its arguments for a HAVING clause
type HavingClause struct {
	Condition string
	Arguments []interface{}
}

func (c HavingClause) String() string {
	sql := fmt.Sprintf("%s", c.Condition)

	return sql
}
