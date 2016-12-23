package gorules

import "fmt"

// Selector type
type Selector int

const (
	This Selector = iota
	Any
	All
	maxSelectorFlag
)

var selectorNames = [...]string{
	This: "THIS",
	Any:  "ANY",
	All:  "ALL",
}

// ToSelector converts string to Selector
func ToSelector(s string) (Selector, error) {
	for i, r := range selectorNames {
		if s == r {
			return Selector(i), nil
		}
	}
	return maxSelectorFlag, fmt.Errorf("Invalid Selector value %q", s)
}

func (v Selector) String() string {
	if v >= maxSelectorFlag {
		return "Invalid Selector"
	}
	return selectorNames[v]
}

// IsSelector checks if the string is a valid Selector
func IsSelector(value string) bool {
	_, err := ToSelector(value)
	if err == nil {
		return true
	}
	return false
}
