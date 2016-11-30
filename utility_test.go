package gorules_test

import (
	"testing"
	//"fmt"
	"github.com/stretchr/testify/assert"
)

func TestGetNumber(t *testing.T) {

	n, _ := getNumber("32")

	assert.Equal(t, n, float64(32))

	f, err := getNumber("abc")

	assert.Error(t, err)

	assert.Equal(t, f, float64(0))

}

func TestEquals (t *testing.T) {

  v := "test"
  s := "test"
  assert.Equal(t,v,s,"values are not equal")
  s = "me"
  assert.NotEqual(t,v,s,"values are should not be equal")

}

