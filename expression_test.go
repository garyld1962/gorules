package gorules

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

// func TestAndIsConjunctionExpression(t *testing.T) {
// 	tests := CreateAndConjunctionExpression(&gorules.TrueExpression)
// 	assert.True(t, isConjunctionExpression(tests))
// }

// func TestOrIsConjunctionExpression(t *testing.T) {
// 	tests := CreateOrConjunctionExpression(&gorules.FalseExpression)
// 	assert.True(t, isConjunctionExpression(tests))
// }
