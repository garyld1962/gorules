package gorules

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestToOperator(t *testing.T) {
	operator, _ := toOperator("IsEqualTo")
	assert.Equal(t, operator, IsEqualTo)
}

func TestToOperatorInValid(t *testing.T) {
	_, err := toOperator("IsEqualToS")
	assert.NotEqual(t, err, nil)
}

func TestOperatorAsString(t *testing.T) {
	operator, _ := toOperator("IsEqualTo")
	txt := fmt.Sprintf("%s", operator)
	assert.Equal(t, txt, "IsEqualTo")
}

func TestIsOperator(t *testing.T) {
	isO := isOperator("IsEqualTo")
	assert.True(t, isO)

}
