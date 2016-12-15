package gorules

import "fmt"

// Operator to maintain the list of operatorNames as a single entity
type Operator int

const (
	IsEqualTo Operator = iota
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

func (v Operator) String() string {
	if v >= maxOperatorFlag {
		return "Invalid Operator"
	}
	return operatorNames[v]
}

// ToOperator converts string to one of the operatorNames in const declaration
func ToOperator(s string) (Operator, error) {
	for i, r := range operatorNames {
		if s == r {
			return Operator(i), nil
		}
	}
	return maxOperatorFlag, fmt.Errorf("Invalid Operator value %q", s)
}

// IsOperator check if a string is a Operator Names
func IsOperator(value string) bool {
	_, err := ToOperator(value)
	if err == nil {
		return true
	}
	return false
}
