package gorules_test

import (
	"gorules"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

var parserTestData = `{
  "id": 25,
  "zip5": 33076,
  "zip3": 333,
  "state": "FL",
  "country": "USA",
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

//IF THIS country NOTEQUALS USA |AND| NOTEQUALS CANADA |AND| zip3 EQUALS "333"
// func TestValue(t *testing.T) {
// 	//	m := gorules.ParseDSL("country IsEqualTo USA |AND| country IsEqualTo USA |AND| state IsEqualTo FL |AND| IF THIS subtotal IsEqualTo 25.0", parserTestData)
// 	result := gorules.EvaluateDSL("country IsEqualTo CANADA AND country IsEqualTo INDIA AND country IsEqualTo USA", parserTestData)
// 	fmt.Println(result)
// 	assert.NotNil(t, result)
// }

func TestCollectionalue(t *testing.T) {
	result := gorules.EvaluateDSL("ANY orderItems.weight IsEqualTo 10", parserTestData)
	fmt.Println(result)
	assert.NotNil(t, result)
}

// func TestReduceRuleToBool(t *testing.T) {
// 	var rle = &gorules.Rule{}
// 	exp := gorules.CreateConjunctionStatement("AND")
// 	parsed, _ := exp.Parse(exp)
// 	fmt.Println("pared", exp)
// 	rle.Add(&parsed)
// 	assert.NotNil(t, nil)
// }
