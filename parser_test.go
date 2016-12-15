package gorules_test

import (
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
  "country": "SA",
  "subtotal": 25.00,
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
	m := gorules.ParseDSL("IF THIS country IsEqualTo USA |AND| IsEqualTo CANADA", parserTestData)
	//x := gorules.GetKeyFromJSON(m,"testobj")
	//fmt.Println("Type: ",m)
	assert.NotNil(t, m)
}
