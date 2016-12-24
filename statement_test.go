package gorules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConjunctionStatement(t *testing.T) {
	conjunction := createConjunctionStatement("AND")
	assert.Equal(t, (*conjunction).Conjunction, And)
}
