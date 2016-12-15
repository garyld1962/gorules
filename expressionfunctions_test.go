package gorules_test

import (
	"fmt"
	rules "gorules"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEqual(t *testing.T) {

	tt := rules.CreateValueExpressionWithTarget("IsEqualTo", "", "one", "one")

	ret, _ := tt.Evaluate()
	fmt.Println(ret)

	assert.True(t, ret, "Target should have been null")

}
