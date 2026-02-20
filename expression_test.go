package gorules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBooleanExpressionTrue(t *testing.T) {
	expr := createBooleanExpression(true)
	result, err := expr.Evaluate()
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestBooleanExpressionFalse(t *testing.T) {
	expr := createBooleanExpression(false)
	result, err := expr.Evaluate()
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestRuleExpressionEquals(t *testing.T) {
	expr, err := createRuleExpressionWithTarget("IsEqualTo", "USA", "USA")
	assert.NoError(t, err)
	result, err := expr.Evaluate()
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestRuleExpressionNotEquals(t *testing.T) {
	expr, err := createRuleExpressionWithTarget("IsEqualTo", "USA", "CANADA")
	assert.NoError(t, err)
	result, err := expr.Evaluate()
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestRuleExpressionInvalidOperator(t *testing.T) {
	_, err := createRuleExpressionWithTarget("InvalidOp", "a", "b")
	assert.Error(t, err)
}

func TestRuleExpressionGreaterThan(t *testing.T) {
	expr, err := createRuleExpressionWithTarget("IsGreaterThan", "20", "10")
	assert.NoError(t, err)
	result, err := expr.Evaluate()
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestConjunctionExpressionAnd(t *testing.T) {
	conj := createAndConjunctionExpression(TrueExpression).(ConjunctionExpression)
	conj = conj.Add(createBooleanExpression(true))
	result, err := conj.Evaluate()
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestConjunctionExpressionAndFalse(t *testing.T) {
	conj := createAndConjunctionExpression(TrueExpression).(ConjunctionExpression)
	conj = conj.Add(createBooleanExpression(false))
	result, err := conj.Evaluate()
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestConjunctionExpressionOr(t *testing.T) {
	conj := createOrConjunctionExpression(FalseExpression).(ConjunctionExpression)
	conj = conj.Add(createBooleanExpression(true))
	result, err := conj.Evaluate()
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestConjunctionExpressionOrBothFalse(t *testing.T) {
	conj := createOrConjunctionExpression(FalseExpression).(ConjunctionExpression)
	conj = conj.Add(createBooleanExpression(false))
	result, err := conj.Evaluate()
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestIsConjunctionExpressionTrue(t *testing.T) {
	conj := createAndConjunctionExpression(TrueExpression)
	assert.True(t, isConjunctionExpression(conj))
}

func TestIsConjunctionExpressionFalseForRule(t *testing.T) {
	expr := createBooleanExpression(true)
	assert.False(t, isConjunctionExpression(expr))
}

func TestNilOperatorDispatch(t *testing.T) {
	expr := RuleExpression{Operator: maxOperatorFlag, Value: "a", Target: "b"}
	_, err := expr.Evaluate()
	assert.Error(t, err)
}
