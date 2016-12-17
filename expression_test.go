package gorules_test

import (
	"gorules"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestAndIsConjunctionExpression(t *testing.T) {
	// trueExpr = gorules.
	tests := gorules.CreateAndConjunctionExpression(gorules.TrueExpression)
	assert.True(t, gorules.IsConjunctionExpression(tests))
}

func TestOrIsConjunctionExpression(t *testing.T) {
	tests := gorules.CreateOrConjunctionExpression(gorules.FalseExpression)
	assert.True(t, gorules.IsConjunctionExpression(tests))
}

func TestIsConjunctionExpressionFail(t *testing.T) {
	tests := gorules.CreateValueExpressionWithTarget("IsEqualTo", "test", "USA", "USA")
	assert.False(t, gorules.IsConjunctionExpression(tests))
}
