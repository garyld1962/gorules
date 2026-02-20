package gorules_test

import (
	"gorules"
	"testing"

	"github.com/stretchr/testify/assert"
)

var parserTestDataString = `{
  "id": 25,
  "zip5": 33076,
  "zip3": 333,
  "state": "FL",
  "country": "SOUTH A",
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

var parserTestData, _ = gorules.ParseStringToJSON(parserTestDataString)

func TestCollectionalue(t *testing.T) {
	result := gorules.DSLEvaluator("'SOUTH A' IsEqualTo country", parserTestData)
	assert.True(t, result)
}
