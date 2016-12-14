package gorules_test

import (
  "gorules"
	"testing"
  "github.com/stretchr/testify/assert"
  "fmt"
)
	// objects "github.com/stretchr/stew/objects"

var testData = `{
  "id": 25,
  "zip5": 33076,
  "zip3": 333,
  "state": "FL",
  "country": "USA",
  "subtotal": 25.00,
  "promoamount": 1.00,
  "orderItems": [
    {
      "id": 3,
      "productId": 34354,
      "quantity": 3,
      "warehouse": "",
      "name": "test product 1",
      "availableInventory": "",
      "promos": []
    }
  ],
  "promos": []
}`


func TestIsNull(t *testing.T) {

	tt := gorules.CreateValueExpressionWithTarget("IsEqualTo", "", "one", "one")

	ret, _ := tt.Evaluate()
  fmt.Println(ret)

	assert.True(t, ret, "Target should have been null")

}

 