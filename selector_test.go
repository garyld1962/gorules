package gorules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSelectorThis(t *testing.T) {
	sel, err := toSelector("THIS")
	assert.NoError(t, err)
	assert.Equal(t, This, sel)
}

func TestToSelectorAny(t *testing.T) {
	sel, err := toSelector("ANY")
	assert.NoError(t, err)
	assert.Equal(t, Any, sel)
}

func TestToSelectorAll(t *testing.T) {
	sel, err := toSelector("ALL")
	assert.NoError(t, err)
	assert.Equal(t, All, sel)
}

func TestToSelectorInvalid(t *testing.T) {
	_, err := toSelector("INVALID")
	assert.Error(t, err)
}

func TestSelectorString(t *testing.T) {
	assert.Equal(t, "THIS", This.String())
	assert.Equal(t, "ANY", Any.String())
	assert.Equal(t, "ALL", All.String())
}

func TestSelectorStringInvalid(t *testing.T) {
	assert.Equal(t, "Invalid Selector", maxSelectorFlag.String())
}

func TestIsSelectorValid(t *testing.T) {
	assert.True(t, isSelector("THIS"))
	assert.True(t, isSelector("ANY"))
	assert.True(t, isSelector("ALL"))
}

func TestIsSelectorInvalid(t *testing.T) {
	assert.False(t, isSelector("INVALID"))
	assert.False(t, isSelector(""))
}

func TestSelectorFunctions(t *testing.T) {
	assert.NotNil(t, selectorFunctions(This))
	assert.NotNil(t, selectorFunctions(Any))
	assert.NotNil(t, selectorFunctions(All))
}

func TestSelectorConjExprMap(t *testing.T) {
	anyExpr := selectorConjExprMap(Any)
	allExpr := selectorConjExprMap(All)
	assert.Equal(t, Or, anyExpr.Conjunction)
	assert.Equal(t, And, allExpr.Conjunction)
}
