package gorules_test

import (
	"gorules"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestCreateConjunctionStatement(t *testing.T) {
	conjunction := gorules.CreateConjunctionStatement("AND")
	fmt.Println(conjunction)
	assert.Equal(t, conjunction, gorules.And)
}
