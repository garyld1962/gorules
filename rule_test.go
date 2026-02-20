package gorules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleAddAndEvaluate(t *testing.T) {
	rule := Rule{}
	rule = rule.Add(createBooleanExpression(true))
	result, err := rule.Evaluate()
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestRuleWithAndConjunction(t *testing.T) {
	rule := Rule{}
	rule = rule.Add(createBooleanExpression(true))
	rule = rule.Add(createAndConjunctionExpression(TrueExpression))
	rule = rule.Add(createBooleanExpression(true))
	result, err := rule.Evaluate()
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestRuleWithAndConjunctionFalse(t *testing.T) {
	rule := Rule{}
	rule = rule.Add(createBooleanExpression(true))
	rule = rule.Add(createAndConjunctionExpression(TrueExpression))
	rule = rule.Add(createBooleanExpression(false))
	result, err := rule.Evaluate()
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestRuleWithOrConjunction(t *testing.T) {
	rule := Rule{}
	rule = rule.Add(createBooleanExpression(false))
	rule = rule.Add(createOrConjunctionExpression(FalseExpression))
	rule = rule.Add(createBooleanExpression(true))
	result, err := rule.Evaluate()
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestIsRule(t *testing.T) {
	rule := Rule{}
	assert.True(t, isRule(rule))
}

func TestIsRuleFalse(t *testing.T) {
	expr := createBooleanExpression(true)
	assert.False(t, isRule(expr))
}

func TestEmptyRuleEvaluatesToFalse(t *testing.T) {
	rule := Rule{}
	result, err := rule.Evaluate()
	assert.NoError(t, err)
	assert.False(t, result)
}
