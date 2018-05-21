package panda

import "fmt"

// CommonSQLBuilder implements SQLBuilder interface for common Dialect
type CommonSQLBuilder struct {
	query Query
	sql   Clause
}

// NewCommonSQLBuilder creates new instance
func NewCommonSQLBuilder(q Query) *CommonSQLBuilder {
	return &CommonSQLBuilder{
		query: q,
		sql:   Clause{},
	}
}

// Build builds Query object to SQL string and appropriate values
func (b *CommonSQLBuilder) Build() (string, []interface{}) {
	if b.sql.Statement == "" {
		b.compile()
	}
	return b.sql.Statement, b.sql.Arguments
}

func (b *CommonSQLBuilder) compile() {
	// check if we are compiling raw SQL
	if b.query.RawSQL.Statement != "" {
		b.sql = b.query.RawSQL
	} else {
		b.sql.Statement = b.buildSQLQuery()
	}
}

func (b *CommonSQLBuilder) buildSQLQuery() string {

	sql := fmt.Sprintf("SELECT %s FROM %s", b.query.columnClauses, b.query.fromClause)
	sql = b.buildJoinClauses(sql)
	sql = b.buildWhereClauses(sql)
	sql = b.buildGroupClauses(sql)
	sql = b.buildOrderClauses(sql)
	return sql
}

func (b *CommonSQLBuilder) buildWhereClauses(sql string) string {
	wc := b.query.whereClause
	if wc.Statement != "" {
		sql = fmt.Sprintf("%s WHERE %s", sql, wc)
		for _, arg := range wc.Arguments {
			b.sql.Arguments = append(b.sql.Arguments, arg)
		}
	}
	return sql
}

func (b *CommonSQLBuilder) buildJoinClauses(sql string) string {
	jc := b.query.joinClauses
	if len(jc) > 0 {
		sql += jc.String()

		// assign all Arguments from each join clause to sql Arguments
		for i := range jc {
			for _, arg := range jc[i].Arguments {
				b.sql.Arguments = append(b.sql.Arguments, arg)
			}
		}
	}

	return sql
}

func (b *CommonSQLBuilder) buildGroupClauses(sql string) string {
	gc := b.query.groupClause
	if gc.Field != "" {
		sql = fmt.Sprintf("%s GROUP BY %s", sql, gc.String())

		hc := b.query.havingClause
		if hc.Condition != "" {
			sql = fmt.Sprintf("%s HAVING %s", sql, hc.String())
			for _, arg := range hc.Arguments {
				b.sql.Arguments = append(b.sql.Arguments, arg)
			}
		}

	}
	return sql
}

func (b *CommonSQLBuilder) buildOrderClauses(sql string) string {
	oc := b.query.orderClause
	if oc.Statement != "" {
		sql = fmt.Sprintf("%s ORDER BY %s", sql, oc.String())
		for _, arg := range oc.Arguments {
			b.sql.Arguments = append(b.sql.Arguments, arg)
		}
	}
	return sql
}
