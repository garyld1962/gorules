package gorules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNumber(t *testing.T) {
	n, _ := getNumber("32")
	assert.Equal(t, n, float64(32))
	f, err := getNumber("abc")
	assert.Error(t, err)
	assert.Equal(t, f, float64(0))

}

func TestEquals(t *testing.T) {
	v := "test"
	s := "test"
	assert.Equal(t, v, s, "values are not equal")
	s = "me"
	assert.NotEqual(t, v, s, "values are should not be equal")
}

func TestStartsWithQuotesPass(t *testing.T) {
	s := "'test'"
	assert.True(t, startsWithSingleQuotes(s))
}

func TestStartsWithQuotesFail(t *testing.T) {
	s := "`test`"
	assert.False(t, startsWithSingleQuotes(s))
}

func TestStringBetweenSingleQuotesPass(t *testing.T) {
	s := "country isEQUALTO 'USA'"
	assert.Equal(t, "USA", stringBetweenSingleQuotes(s), "stringBetweenSingleQuotes works")
}

func TestStringBetweenSingleQuotesFail(t *testing.T) {
	s := "test"
	assert.Equal(t, "", stringBetweenSingleQuotes(s), "stringBetweenSingleQuotes works")
}
