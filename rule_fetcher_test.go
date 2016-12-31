package gorules_test

import (
	"gorules"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchFromFile(t *testing.T) {
	n := gorules.NewRuleFetcher("country").Fetch()
	assert.Equal(t, "country IsEqualTo 'CANADA'", n)
}
