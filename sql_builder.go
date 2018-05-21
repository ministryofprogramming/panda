package panda

// SQLBuilder interface responsible for building Query object to proper SQL query
type SQLBuilder interface {
	Build() (string, []interface{})
}
