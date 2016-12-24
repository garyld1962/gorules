package gorules

import "fmt"

// Operator to maintain the list of operatorNames as a single entity
type Operator int

const (
	// IsEqualTo compares values for equals
	IsEqualTo Operator = iota
	// IsGreaterThan compares values for greater than value
	IsGreaterThan
	maxOperatorFlag
	// IsNotEqualTo
	// IsGreaterThan
	// In
	// NotIn
	// IsLessThan
	// IsGreaterThanOrEqualTo
	// IsLessThanOrEqualTo
	// IsLike
	// IsNotLike
	// IsNull
	// IsNotNull
)

var operatorNames = [...]string{
	IsEqualTo:     "IsEqualTo",
	IsGreaterThan: "IsGreaterThan",
}

// String makes Operator implement Stringer
func (o Operator) String() string {
	if o >= maxOperatorFlag {
		return "Invalid Operator"
	}
	return operatorNames[o]
}

// toOperator converts string to one of the operatorNames in const declaration
func toOperator(s string) (Operator, error) {
	for i, r := range operatorNames {
		if s == r {
			return Operator(i), nil
		}
	}
	return maxOperatorFlag, fmt.Errorf("Invalid Operator value %q", s)
}

// isOperator check if a string is a Operator Names
func isOperator(value string) bool {
	_, err := toOperator(value)
	if err == nil {
		return true
	}
	return false
}

type operatorFunc func(string, string) (bool, error)

var operatorFuncList = map[Operator]operatorFunc{
	IsEqualTo: equals,
}
