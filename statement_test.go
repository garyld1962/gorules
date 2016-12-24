package gorules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConjunctionStmt(t *testing.T) {
	conjunction := createConjunctionStmt("AND")
	assert.Equal(t, (*conjunction).Conjunction, And)
}
