package gorules_test

import (
	"gorules"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

var valueTestData = `{
  "id": 25,
  "zip5": 33076,
  "zip3": 333,
  "state": "FL",
  "country": "SOUTH",
  "subtotal": "25.00",
  "promoamount": 1.00,
  "testobj":{
      "id": 3,
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
      "availableInventory": "",
      "promos": []
    },{
      "id": 3,
      "productId": 34354,
      "quantity": 3,
      "weight" : "11",
      "warehouse": "",
      "name": "test product 1",
      "availableInventory": "",
      "promos": []
    }
  ],
  "promos": []
}`

func TestConstants(t *testing.T) {
	result := gorules.NewValue("'25'")
	str, _ := result.Evaluate(make([]interface{}, 0))
	fmt.Println(result, str)
	//fmt.Println(valueTestData)
	assert.Equal(t, "25", str)
}

func TestPath(t *testing.T) {
	result := gorules.NewValue("country")
	str, _ := result.Evaluate(gorules.ParseStringToJSON(valueTestData))
	fmt.Println(result, str)
	//fmt.Println(valueTestData)
	assert.Equal(t, "SOUTH", str)
}

func TestMultiplicationSeparationByOperator(t *testing.T) {
	inputString := "'6' * '6'"
	newMathExpression := gorules.NewMathExpression(inputString)
	fmt.Println(newMathExpression)
	outputValue, _ := newMathExpression.Evaluate("Dummy Data")
	assert.Equal(t, "36", outputValue)
}

func TestDivisionSeparationByOperator(t *testing.T) {
	inputString := "'120' / '6'"
	newMathExpression := gorules.NewMathExpression(inputString)
	fmt.Println(newMathExpression)
	outputValue, _ := newMathExpression.Evaluate("Dummy Data")
	assert.Equal(t, "20", outputValue)
}

func TestAdditionSeparationByOperator(t *testing.T) {
	inputString := "'16' + '26'"
	newMathExpression := gorules.NewMathExpression(inputString)
	fmt.Println(newMathExpression)
	outputValue, _ := newMathExpression.Evaluate("Dummy Data")
	assert.Equal(t, "42", outputValue)
}
