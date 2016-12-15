package gorules_test

import (
	"gorules"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestToOperator(t *testing.T) {
	operator, _ := gorules.ToOperator("IsEqualTo")
	assert.Equal(t, operator, gorules.IsEqualTo)
}

func TestToOperatorInValid(t *testing.T) {
	_, err := gorules.ToOperator("IsEqualToS")
	assert.NotEqual(t, err, nil)
}

func TestOperatorAsString(t *testing.T) {
	operator, _ := gorules.ToOperator("IsEqualTo")
	txt := fmt.Sprintf("%s", operator)
	assert.Equal(t, txt, "IsEqualTo")
}

func TestIsOperator(t *testing.T) {
	isOperator := gorules.IsOperator("IsEqualTo")
	assert.True(t, isOperator)

}

func TestIsOperatorFail(t *testing.T) {
	isOperator := gorules.IsOperator("IsEqualTos")
	assert.False(t, isOperator)

}
