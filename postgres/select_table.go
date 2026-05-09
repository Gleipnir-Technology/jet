package postgres

import "github.com/Gleipnir-Technology/jet/internal/jet"

// SelectTable is interface for postgres temporary tables like sub-queries, VALUES, CTEs etc...
type SelectTable interface {
	readableTable
	jet.SelectTable
}

type selectTableImpl struct {
	jet.SelectTable
	readableTableInterfaceImpl
}

func newSelectTable(serializerWithProjections jet.SerializerHasProjections, alias string, columnAliases []jet.ColumnExpression) SelectTable {
	subQuery := &selectTableImpl{
		SelectTable: jet.NewSelectTable(serializerWithProjections, alias, columnAliases),
	}

	subQuery.readableTableInterfaceImpl.root = subQuery

	return subQuery
}
