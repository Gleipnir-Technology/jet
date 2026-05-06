// Modeling of columns

package jet

// Column is common column interface for all types of columns.
type Column interface {
	Name() string
	TableName() string

	setTableName(table string)
	setSubQuery(subQuery SelectTable)
	defaultAlias() string
}

// ColumnSerializer is interface for all serializable columns
type ColumnSerializer interface {
	Serializer
	Column
}

// ColumnExpression interface
type ColumnExpression interface {
	Column
	Expression
}

