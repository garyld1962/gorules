package gorules

import "fmt"

//go:generate stringer -type=Conjunction

// Conjunction type
type Conjunction int

const (
	And Conjunction = iota
	Or
	maxConjunctionFlag
)

var conjunctionNames = [...]string{
	And: "AND",
	Or:  "OR",
}

// ToConjunction converts string to Conjunction
func ToConjunction(s string) (Conjunction, error) {
	for i, r := range conjunctionNames {
		if s == r {
			return Conjunction(i), nil
		}
	}
	return maxConjunctionFlag, fmt.Errorf("Invalid Conjunction value %q", s)
}

func (v Conjunction) String() string {
	if v >= maxConjunctionFlag {
		return "Invalid Conjunction"
	}
	return conjunctionNames[v]
}

// IsConjunction checks if the string is a valid Conjunction
func IsConjunction(value string) bool {
	_, err := ToConjunction(value)
	if err == nil {
		return true
	}
	return false
}
