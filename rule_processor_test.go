package gorules_test

import (
	"gorules"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

var fetcherTestDataString = `{
  "id": 25,
  "zip5": 33076,
  "zip3": 333,
  "state": "FL",
  "country": "CANADA",
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

var fetcherTestData = gorules.ParseStringToJSON(fetcherTestDataString)

func TestFetchFromFile(t *testing.T) {
	n := gorules.NewRuleFetcher("is-country-canada").(gorules.RuleFromFile).Fetch()
	assert.Equal(t, "country IsEqualTo 'CANADA'", n)
}

func TestFetchFromRuleCollection(t *testing.T) {
	rules := []string{"country IsEqualTo 'USA'", "is-country-canada"}
	n := gorules.NewRuleFetcher(rules).(gorules.RuleCollection).Process(fetcherTestData)
	fmt.Println(n)
	assert.True(t, true)
}

func TestFetchFromObjectCollection(t *testing.T) {
	rules := map[string]string{"is-country-usa": "country IsEqualTo 'USA'", "is-country-canada": "is-country-canada"}
	n := gorules.NewRuleFetcher(rules).(gorules.RuleObjectCollection).Process(fetcherTestData)
	fmt.Println(n)
	assert.True(t, true)
}
