package gorules_test

import (
	"gorules"
)

var eval = gorules.ValueExpressionEvaluator{}

/*
func TestIsNull(t *testing.T) {

	tt := ValueExpression{"IsNull", "", "one", ""}

	ret, _ := eval.Evaluate(tt)

	assert.True(t, ret, "Target should have been null")

}

func TestIsNotNull(t *testing.T) {

	tt := ValueExpression{"IsNotNull", "", "one", "two"}

	ret, _ := eval.Evaluate(tt)

	assert.True(t, ret, "Target should be null")

}

func TestIsEqualString(t *testing.T) {

	tt := ValueExpression{"IsEqualTo", "", "one", "one"}

	ret, _ := eval.Evaluate(tt)

	assert.True(t, ret, "Source and Target strings should be equal")

}

func TestIsEqualStringFalse(t *testing.T) {

	tt := ValueExpression{"IsEqualTo", "", "one", "two"}

	ret, _ := eval.Evaluate(tt)

	assert.False(t, ret, "Source and Target strings should not be equal")

}

func TestIsNotEqualString(t *testing.T) {

	tt := ValueExpression{"IsNotEqualTo", "", "one", "two"}

	ret, _ := eval.Evaluate(tt)

	assert.True(t, ret, "Source and Target strings should be equal")

}

func TestIsEqualNumeric(t *testing.T) {

	tt := ValueExpression{"IsEqualTo", "", "1", "1"}

	ret, _ := eval.Evaluate(tt)

	assert.True(t, ret, "Source and Target numbers should be equal")

}

func TestIsEqualNumericFalse(t *testing.T) {

	tt := ValueExpression{"IsEqualTo", "", "1", "2"}

	ret, _ := eval.Evaluate(tt)

	assert.False(t, ret, "Source and Target numbers should be equal")

}

func TestIsGreaterThan(t *testing.T) {

	tt := ValueExpression{"IsGreater", "", "3", "1"}

	ret, _ := eval.Evaluate(tt)

	assert.True(t, ret, "Source should be greater than Target")
}



/*
func TestIsGreaterThanOrEqualTo(t *testing.T) {

	tt := ValueExpression{"IsGreaterThanOrEqualTo", "", "3", "1"}

	ret, _ := eval.Evaluate(tt)

	assert.True(t, ret, "Source should be greater than or equal to Target")
}

func TestIsGreaterThanOrEqualToEquals(t *testing.T) {

	tt := ValueExpression{"IsGreaterThanOrEqualTo", "", "3", "3"}
  st := ValueExpression{"IsGreaterThanOrEqualTo", "", "2", "3"}

	ret, _ := eval.Evaluate(tt)

	assert.True(t, ret, "Source should be greater than or equal to Target")

  ret1, _ := eval.Evaluate(st)
  assert.False(t, ret1, "Source should be greater than or equal to Target")
}
*/
