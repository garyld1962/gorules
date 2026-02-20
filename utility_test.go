package gorules

import (
	"strings"
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

func TestLastButAllElement(t *testing.T) {
	tests := lastButAllElements(strings.Split("The quick brown", " "))
	assert.Equal(t, 2, len(tests))
}

func TestMakeLastWordFirst(t *testing.T) {
	tests := makeLastWordFirst("shipping.shippingweight  ISLESSTHAN 100 OR")
	out, _ := startsWith(tests, "OR")
	assert.True(t, out)
}

// Comparison operator boundary tests

func TestIsGreaterOrEqualBoundary(t *testing.T) {
	result, err := isGreaterOrEqual("10", "10")
	assert.NoError(t, err)
	assert.True(t, result, "equal values should return true")
}

func TestIsGreaterOrEqualGreater(t *testing.T) {
	result, err := isGreaterOrEqual("11", "10")
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestIsGreaterOrEqualLess(t *testing.T) {
	result, err := isGreaterOrEqual("9", "10")
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestIsGreaterOrEqualInvalidInput(t *testing.T) {
	_, err := isGreaterOrEqual("abc", "10")
	assert.Error(t, err)
}

func TestIsLessOrEqualBoundary(t *testing.T) {
	result, err := isLessOrEqual("10", "10")
	assert.NoError(t, err)
	assert.True(t, result, "equal values should return true")
}

func TestIsLessOrEqualLess(t *testing.T) {
	result, err := isLessOrEqual("9", "10")
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestIsLessOrEqualGreater(t *testing.T) {
	result, err := isLessOrEqual("11", "10")
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestIsGreaterPass(t *testing.T) {
	result, err := isGreater("11", "10")
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestIsGreaterFail(t *testing.T) {
	result, err := isGreater("10", "10")
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestIsGreaterInvalidInput(t *testing.T) {
	_, err := isGreater("abc", "10")
	assert.Error(t, err)
}

// Boolean operator tests

func TestIsTrueWithTrue(t *testing.T) {
	result, err := isTrue("", "true")
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestIsTrueWithFalse(t *testing.T) {
	result, err := isTrue("", "false")
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestIsTrueInvalidInput(t *testing.T) {
	_, err := isTrue("", "notabool")
	assert.Error(t, err)
}

func TestIsFalseWithTrue(t *testing.T) {
	result, err := isFalse("", "true")
	assert.NoError(t, err)
	assert.False(t, result, "isFalse('true') should return false")
}

func TestIsFalseWithFalse(t *testing.T) {
	result, err := isFalse("", "false")
	assert.NoError(t, err)
	assert.True(t, result, "isFalse('false') should return true")
}

// Null/NotNull tests

func TestIsNullEmpty(t *testing.T) {
	result, err := isNull("", "")
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestIsNullNonEmpty(t *testing.T) {
	result, err := isNull("value", "")
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestIsNotNullEmpty(t *testing.T) {
	result, err := isNotNull("", "")
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestIsNotNullNonEmpty(t *testing.T) {
	result, err := isNotNull("value", "")
	assert.NoError(t, err)
	assert.True(t, result)
}

// String operator tests

func TestContainsPass(t *testing.T) {
	result, err := contains("hello world", "world")
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestContainsFail(t *testing.T) {
	result, err := contains("hello world", "xyz")
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestStartsWithPass(t *testing.T) {
	result, err := startsWith("hello world", "hello")
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestStartsWithFail(t *testing.T) {
	result, err := startsWith("hello world", "world")
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestEndsWithPass(t *testing.T) {
	result, err := endsWith("hello world", "world")
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestEndsWithFail(t *testing.T) {
	result, err := endsWith("hello world", "hello")
	assert.NoError(t, err)
	assert.False(t, result)
}

// selectValue tests

func TestSelectValueSimplePath(t *testing.T) {
	data := map[string]interface{}{"country": "USA"}
	result := selectValue(data, "country")
	assert.Equal(t, "USA", result)
}

func TestSelectValueNestedPath(t *testing.T) {
	data := map[string]interface{}{
		"testobj": map[string]interface{}{
			"id": "3",
		},
	}
	result := selectValue(data, "testobj.id")
	assert.Equal(t, "3", result)
}

func TestSelectValueMissingKey(t *testing.T) {
	data := map[string]interface{}{"country": "USA"}
	result := selectValue(data, "missing")
	assert.Nil(t, result)
}

func TestSelectValueMissingNestedKey(t *testing.T) {
	data := map[string]interface{}{"country": "USA"}
	result := selectValue(data, "missing.nested")
	assert.Nil(t, result)
}

// isArray/isMap nil safety tests

func TestIsArrayNil(t *testing.T) {
	assert.False(t, isArray(nil))
}

func TestIsArraySlice(t *testing.T) {
	assert.True(t, isArray([]int{1, 2, 3}))
}

func TestIsMapNil(t *testing.T) {
	assert.False(t, isMap(nil))
}

func TestIsMapValid(t *testing.T) {
	assert.True(t, isMap(map[string]string{"a": "b"}))
}

// isMatch test

func TestIsMatchPass(t *testing.T) {
	result, err := isMatch("SOUTH AMERICA", "^SOUTH.*")
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestIsMatchFail(t *testing.T) {
	result, err := isMatch("SOUTH AMERICA", "^NORTH.*")
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestIsMatchInvalidRegex(t *testing.T) {
	_, err := isMatch("test", "[invalid")
	assert.Error(t, err)
}
