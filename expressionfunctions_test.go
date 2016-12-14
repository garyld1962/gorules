package gorules_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	 rules "gorules"
	"testing"
)

func TestIsEqual(t *testing.T) {

	tt := rules.CreateValueExpressionWithTarget("IsEqualTo", "", "one", "one")

	ret, _ := tt.Evaluate()
	fmt.Println(ret)

	assert.True(t, ret, "Target should have been null")

}
