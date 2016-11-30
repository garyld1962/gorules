package gorules_test

import (
	"testing"
	objects "github.com/stretchr/stew/objects"
)

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

func TestValueFromJson(t *testing.T) {

	m, _ := objects.NewMapFromJSON(testData)
	value := m.Get("country").(string)
	if string(value) != "USA" {
		t.Error("country should return USA")
	}
}

 