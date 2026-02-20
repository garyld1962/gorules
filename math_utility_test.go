package gorules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	result, err := multiply(6, 7)
	assert.NoError(t, err)
	assert.Equal(t, 42, result)
}

func TestMultiplyByZero(t *testing.T) {
	result, err := multiply(100, 0)
	assert.NoError(t, err)
	assert.Equal(t, 0, result)
}

func TestDivide(t *testing.T) {
	result, err := divide(120, 6)
	assert.NoError(t, err)
	assert.Equal(t, 20, result)
}

func TestDivideByZero(t *testing.T) {
	result, err := divide(10, 0)
	assert.Error(t, err)
	assert.Equal(t, 0, result)
	assert.Contains(t, err.Error(), "division by zero")
}

func TestAdd(t *testing.T) {
	result, err := add(16, 26)
	assert.NoError(t, err)
	assert.Equal(t, 42, result)
}

func TestSubtract(t *testing.T) {
	result, err := subtract(50, 8)
	assert.NoError(t, err)
	assert.Equal(t, 42, result)
}

func TestSquareValue(t *testing.T) {
	result, err := square(16, 0)
	assert.NoError(t, err)
	assert.Equal(t, 256, result)
}

func TestTakeTill(t *testing.T) {
	result, err := takeTill("-", "SOUTH-EAST")
	assert.NoError(t, err)
	assert.Equal(t, "SOUTH", result)
}

func TestTakeTillNoDelimiter(t *testing.T) {
	result, err := takeTill("-", "SOUTH")
	assert.NoError(t, err)
	assert.Equal(t, "SOUTH", result)
}

func TestToUpper(t *testing.T) {
	result, err := toUpper("hello", "")
	assert.NoError(t, err)
	assert.Equal(t, "HELLO", result)
}

func TestToLower(t *testing.T) {
	result, err := toLower("HELLO", "")
	assert.NoError(t, err)
	assert.Equal(t, "hello", result)
}
