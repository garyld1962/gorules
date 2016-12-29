package gorules

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStringSlice = `shipping.shippingweight ISGREATERTHAN |shipping.totalshippingweight + 2|  AND
                       shipping.shippingweight ISLESSTHAN 100 OR
                       shipping.country EQUALS USA OR
					   shipping.country EQUALS CANADA
                       AND
                       state EQUALS FL AND 
					   country EQUALS USA`

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

func TestLastButAllElement(t *testing.T) {
	tests := lastButAllElements(strings.Split("The quick brown", " "))
	fmt.Println(tests)
	assert.Equal(t, 2, len(tests))
}

func TestMakeLastWordFirst(t *testing.T) {
	tests := makeLastWordFirst("shipping.shippingweight  ISLESSTHAN 100 OR")
	fmt.Println(tests)
	out, _ := startsWith(tests, "OR")
	assert.True(t, out)
}

// func TestLines(t *testing.T) {
// 	lineList := lines(testStringSlice)
// 	fmt.Println(lineList)
// 	assert.Equal(t, 5, len(lineList))
// }

// func TestPrecedene(t *testing.T) {
// 	lineList := markPrecedence(testStringSlice)
// 	fmt.Println(lineList)
// 	assert.True(t, true)
// }
