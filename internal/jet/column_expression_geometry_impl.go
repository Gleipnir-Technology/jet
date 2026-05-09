package jet

import (
	"github.com/Gleipnir-Technology/jet/internal/3rdparty/snaker"
)


// ColumnExpressionGeometryImpl is base type for sql columns.
type ColumnExpressionGeometryImpl struct {
	ExpressionInterfaceImpl

	name      string
	tableName string

	subQuery SelectTable
}

// Name returns name of the column
func (c *ColumnExpressionGeometryImpl) Name() string {
	return c.name
}

// TableName returns column table name
func (c *ColumnExpressionGeometryImpl) TableName() string {
	return c.tableName
}

func (c *ColumnExpressionGeometryImpl) setTableName(table string) {
	c.tableName = table
}

func (c *ColumnExpressionGeometryImpl) setSubQuery(subQuery SelectTable) {
	c.subQuery = subQuery
}

func (c *ColumnExpressionGeometryImpl) defaultAlias() string {
	if c.tableName != "" {
		return c.tableName + "." + c.name
	}

	return c.name
}

func (c *ColumnExpressionGeometryImpl) serializeForOrderBy(statement StatementType, out *SQLBuilder) {
	if statement == SetStatementType {
		// set Statement (UNION, EXCEPT ...) can reference only select projections in order by clause
		out.WriteAlias(c.defaultAlias()) //always quote
		return
	}

	c.serialize(statement, out)
}

func (c *ColumnExpressionGeometryImpl) serializeForProjection(statement StatementType, out *SQLBuilder) {
	c.serialize(statement, out)

	out.WriteString("AS")

	out.WriteAlias(c.defaultAlias())
}

func (c *ColumnExpressionGeometryImpl) serializeForJsonObjEntry(statement StatementType, out *SQLBuilder) {
	out.WriteJsonObjKey(snaker.SnakeToCamel(c.name, false))
	c.Root.serializeForJsonValue(statement, out)
}

func (c *ColumnExpressionGeometryImpl) serializeForRowToJsonProjection(statement StatementType, out *SQLBuilder) {
	c.Root.serializeForJsonValue(statement, out)

	out.WriteString("AS")

	out.WriteAlias(snaker.SnakeToCamel(c.name, false))
}

func (c *ColumnExpressionGeometryImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {

	if c.subQuery != nil {
		out.WriteString("ST_AsGeoJSON(")
		out.WriteIdentifier(c.subQuery.Alias())
		out.WriteByte('.')
		out.WriteIdentifier(c.defaultAlias())
		out.WriteString(")")
	} else {
		//out.WriteString("ST_AsGeoJSON(")
		if c.tableName != "" && !contains(options, ShortName) {
			out.WriteIdentifier(c.tableName)
			out.WriteByte('.')
		}

		out.WriteIdentifier(c.name)
		//out.WriteString(")")
	}
}

func NewColumnGeometryImpl(name string, tableName string, root ColumnExpression) *ColumnExpressionGeometryImpl {
	newColumn := &ColumnExpressionGeometryImpl{
		name:      name,
		tableName: tableName,
	}

	if root != nil {
		newColumn.ExpressionInterfaceImpl.Root = root
	} else {
		newColumn.ExpressionInterfaceImpl.Root = newColumn
	}

	return newColumn
}

