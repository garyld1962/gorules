package gorules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConjunctionStmt(t *testing.T) {
	conjunction, err := createConjunctionStmt("AND")
	assert.NoError(t, err)
	assert.Equal(t, conjunction.Conjunction, And)
}
