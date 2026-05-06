package jet

// GeometryExpression interface
type GeometryExpression interface {
	Expression

	EQ(rhs GeometryExpression) BoolExpression
	NOT_EQ(rhs GeometryExpression) BoolExpression
	IS_DISTINCT_FROM(rhs GeometryExpression) BoolExpression
	IS_NOT_DISTINCT_FROM(rhs GeometryExpression) BoolExpression

	LT(rhs GeometryExpression) BoolExpression
	LT_EQ(rhs GeometryExpression) BoolExpression
	GT(rhs GeometryExpression) BoolExpression
	GT_EQ(rhs GeometryExpression) BoolExpression
	BETWEEN(min, max GeometryExpression) BoolExpression
	NOT_BETWEEN(min, max GeometryExpression) BoolExpression

	ADD(rhs Interval) GeometryExpression
	SUB(rhs Interval) GeometryExpression
}

type geometryInterfaceImpl struct {
	root GeometryExpression
}

func (t *geometryInterfaceImpl) EQ(rhs GeometryExpression) BoolExpression {
	return Eq(t.root, rhs)
}

func (t *geometryInterfaceImpl) NOT_EQ(rhs GeometryExpression) BoolExpression {
	return NotEq(t.root, rhs)
}

func (t *geometryInterfaceImpl) IS_DISTINCT_FROM(rhs GeometryExpression) BoolExpression {
	return IsDistinctFrom(t.root, rhs)
}

func (t *geometryInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs GeometryExpression) BoolExpression {
	return IsNotDistinctFrom(t.root, rhs)
}

func (t *geometryInterfaceImpl) LT(rhs GeometryExpression) BoolExpression {
	return Lt(t.root, rhs)
}

func (t *geometryInterfaceImpl) LT_EQ(rhs GeometryExpression) BoolExpression {
	return LtEq(t.root, rhs)
}

func (t *geometryInterfaceImpl) GT(rhs GeometryExpression) BoolExpression {
	return Gt(t.root, rhs)
}

func (t *geometryInterfaceImpl) GT_EQ(rhs GeometryExpression) BoolExpression {
	return GtEq(t.root, rhs)
}

func (t *geometryInterfaceImpl) BETWEEN(min, max GeometryExpression) BoolExpression {
	return NewBetweenOperatorExpression(t.root, min, max, false)
}

func (t *geometryInterfaceImpl) NOT_BETWEEN(min, max GeometryExpression) BoolExpression {
	return NewBetweenOperatorExpression(t.root, min, max, true)
}

func (t *geometryInterfaceImpl) ADD(rhs Interval) GeometryExpression {
	return GeometryExp(Add(t.root, rhs))
}

func (t *geometryInterfaceImpl) SUB(rhs Interval) GeometryExpression {
	return GeometryExp(Sub(t.root, rhs))
}

//-------------------------------------------------

type geometryExpressionWrapper struct {
	geometryInterfaceImpl
	Expression
}

func newGeometryExpressionWrap(expression Expression) GeometryExpression {
	geometryExpressionWrap := &geometryExpressionWrapper{Expression: expression}
	geometryExpressionWrap.geometryInterfaceImpl.root = geometryExpressionWrap
	expression.setRoot(geometryExpressionWrap)
	return geometryExpressionWrap
}

// GeometryExp is timestamp with time zone expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as timestamp with time zone expression.
// Does not add sql cast to generated sql builder output.
func GeometryExp(expression Expression) GeometryExpression {
	return newGeometryExpressionWrap(expression)
}
