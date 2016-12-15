package gorules_test

import (
	"fmt"
	"gorules"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToConjunction(t *testing.T) {
	conjunction, _ := gorules.ToConjunction("AND")
	assert.Equal(t, conjunction, gorules.And)
}

func TestToConjunctionFail(t *testing.T) {
	_, err := gorules.ToConjunction("AN")
	assert.NotEqual(t, err, nil)
}

func TestToConjunctionAsString(t *testing.T) {
	conjunction, _ := gorules.ToConjunction("AND")
	txt := fmt.Sprintf("%s", conjunction)
	assert.Equal(t, txt, "AND")
}

func TestToConjunctionAsStringFail(t *testing.T) {
	conjunction, _ := gorules.ToConjunction("AN")
	txt := fmt.Sprintf("%s", conjunction)
	assert.NotEqual(t, txt, "AND")
}

func TestIsConjunction(t *testing.T) {
	isConjunction := gorules.IsConjunction("AND")
	assert.True(t, isConjunction)

}

func TestIsOperatorFail(t *testing.T) {
	isConjunction := gorules.IsConjunction("AN")
	assert.False(t, isConjunction)
}
