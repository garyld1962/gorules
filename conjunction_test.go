package gorules

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToConjunction(t *testing.T) {
	conjunction, _ := toConjunction("AND")
	assert.Equal(t, conjunction, And)
}

func TestToConjunctionFail(t *testing.T) {
	_, err := toConjunction("AN")
	assert.NotEqual(t, err, nil)
}

func TestToConjunctionAsString(t *testing.T) {
	conjunction, _ := toConjunction("AND")
	txt := fmt.Sprintf("%s", conjunction)
	assert.Equal(t, txt, "AND")
}

func TestToConjunctionAsStringFail(t *testing.T) {
	conjunction, _ := toConjunction("AN")
	txt := fmt.Sprintf("%s", conjunction)
	assert.NotEqual(t, txt, "AND")
}

func TestIsConjunction(t *testing.T) {
	isConjunction := isConjunction("AND")
	assert.True(t, isConjunction)

}

func TestIsOperatorFail(t *testing.T) {
	isConjunction := isConjunction("AN")
	assert.False(t, isConjunction)
}
