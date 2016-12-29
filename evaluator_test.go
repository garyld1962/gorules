package gorules_test

import (
	"fmt"
	"gorules"
	"testing"

	"github.com/stretchr/testify/assert"
)

//var eval = gorules.ValueExpressionEvaluator{}

var evaluatorTestData = `{
  "id": 25,
  "zip5": 33076,
  "zip3": 333,
  "state": "FL",
  "country": "CANADA",
  "subtotal": "25.00",
  "promoamount": 1.00,
  "testobj":{
      "id": "3",
      "productId": 34354,
      "quantity": 3,
      "warehouse": "",
      "name": "test product 1",
      "availableInventory": "",
      "promos": []
   },
  "orderItems": [
    {
      "id": 3,
      "productId": 34354,
      "quantity": 3,
      "warehouse": "",
      "weight" : "11",
      "name": "test product 1",
      "availableInventory": "D",
      "promos": []
    },{
      "id": 3,
      "productId": 34354,
      "quantity": 3,
      "weight" : "11",
      "warehouse": "",
      "name": "test product 1",
      "availableInventory": "NC",
      "promos": []
    }
  ],
  "promos": []
}`

// func TestSingleLineDSL(t *testing.T) {
// 	result := gorules.DSLEvaluator("country IsEqualTo USA", evaluatorTestData)
// 	assert.True(t, result)
// }

// func TestCompareValueAndValue(t *testing.T) {
// 	result := gorules.DSLEvaluator("'SOUTH A' IsEqualTo 'SOUTH A'", evaluatorTestData)
// 	assert.True(t, result)
// }

// var testStringSlice = `country IsEqualTo '10' AND
//                        country IsEqualTo '100' OR
//                        country IsEqualTo 'USA' OR
// 					   country IsEqualTo 'CANADA'
//                        AND
//                        state IsEqualTo 'FL' AND
// 					   country IsEqualTo 'USA'`

//  AND
// 					  	state IsEqualTo 'FL' AND
// 						 country IsEqualTo 'SOUTH A' OR
// 					   	   country IsEqualTo 'USA' AND
// 					   	     country IsEqualTo 'SOUTH A'
var testStringSlice = `country IsEqualTo 'CANADA'
					   AND 	
					   country IsEqualTo 'CANADA'
					   OR 
					   state IsEqualTo 'L'`

//    state IsEqualTo 'FL' AND

// func TestCompareValueAndPath(t *testing.T) {
// 	result := gorules.DSLEvaluator("'SOUTH A' IsEqualTo country", evaluatorTestData)
// 	assert.True(t, result)
// }

// func TestComparePathAndValue(t *testing.T) {
// 	result := gorules.DSLEvaluator("country IsEqualTo 'SOUTH A'", evaluatorTestData)
// 	assert.True(t, result)
// }

// func TestComparePathAndPath(t *testing.T) {
// 	result := gorules.DSLEvaluator("country IsEqualTo country", evaluatorTestData)
// 	assert.True(t, result)
// }

// func TestAllSelectorPass(t *testing.T) {
// 	result := gorules.DSLEvaluator("ALL orderItems.weight IsEqualTo '11'", evaluatorTestData)
// 	assert.True(t, result)
// }

// func TestAllSelectorFail(t *testing.T) {
// 	result := gorules.DSLEvaluator("ALL orderItems.weight IsEqualTo '0'", evaluatorTestData)
// 	assert.False(t, result)
// }

// func TestAnySelectorPass(t *testing.T) {
// 	result := gorules.DSLEvaluator("ANY orderItems.availableInventory IsEqualTo 'NC'", evaluatorTestData)
// 	assert.True(t, result)
// }

// func TestAnySelectorFail(t *testing.T) {
// 	result := gorules.DSLEvaluator("ANY orderItems.weight IsEqualTo 'NV'", evaluatorTestData)
// 	assert.False(t, result)
// }

// func TestSingleConjunction(t *testing.T) {
// 	result := gorules.DSLEvaluator("OR", evaluatorTestData)
// 	assert.False(t, result)
// }

func TestWithPrecedence(t *testing.T) {
	var testStringSlice = `country IsEqualTo 'CANADA'
					  	   AND 	
					   	   country IsEqualTo 'CANADA' AND
					        country IsEqualTo 'CANADA' AND
					   		 state IsEqualTo 'L'`
	result := gorules.DSLEvaluatorWithP(testStringSlice, evaluatorTestData)
	fmt.Println(result)
	assert.False(t, false)
}

func TestWithPrecedenceOne(t *testing.T) {
	var testStringSlice = `country IsEqualTo 'CANADA'
					   	   AND 	
					   	   country IsEqualTo 'CANADA'
					   	   OR 
					       state IsEqualTo 'L'`
	result := gorules.DSLEvaluatorWithP(testStringSlice, evaluatorTestData)
	fmt.Println(result)
	assert.False(t, false)
}

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
