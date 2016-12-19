package gorules_test

import (
	"fmt"
	"gorules"
	"testing"

	"github.com/stretchr/testify/assert"
	//"fmt"
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
      "name": "test product 1",
      "availableInventory": "",
      "promos": []
    }
  ],
  "promos": []
}`

//IF THIS country NOTEQUALS USA |AND| NOTEQUALS CANADA |AND| zip3 EQUALS "333"
func TestValue(t *testing.T) {
	m := gorules.ParseDSL("IF THIS country IsEqualTo USA |AND| IF THIS country IsEqualTo USA |AND| IF THIS state IsEqualTo FL |AND| IF THIS subtotal IsEqualTo 25.00", parserTestData)
	assert.NotNil(t, m)
}

func TestReduceRuleToBool(t *testing.T) {
	var rle = &gorules.Rule{}
	exp := gorules.CreateConjunctionStatement("AND")
	parsed, _ := exp.Parse(exp)
	fmt.Println("pared", exp)
	rle.Add(&parsed)
	assert.NotNil(t, nil)
}
