package gorules

import "fmt"

// StringOperator to hold various String operations
type StringOperator int

const (
	//TakeTill two operands
	TakeTill StringOperator = iota
	maxStringOperatorFlag
)

var stringOperatorNames = [...]string{
	TakeTill: "TAKETILL",
}

// String makes MathOperator implement Stringer
func (o StringOperator) String() string {
	if o >= maxStringOperatorFlag {
		return "Invalid String Operator"
	}
	return stringOperatorNames[o]
}

// toMathOperator converts string to one of the operatorNames in const declaration
func toStringOperator(s string) (StringOperator, error) {
	for i, r := range stringOperatorNames {
		if s == r {
			return StringOperator(i), nil
		}
	}
	return maxStringOperatorFlag, fmt.Errorf("Invalid String Operator value %q", s)
}

// isOperator check if a string is a Operator Names
func isStringOperator(value string) bool {
	_, err := toStringOperator(value)
	if err == nil {
		return true
	}
	return false
}

// StringOperatorList returns the Operators
func StringOperatorList() []string {
	operatorArray := make([]string, 0)
	for _, r := range stringOperatorNames {
		operatorArray = append(operatorArray, r)
	}
	return operatorArray
}

type stringOperatorFunc func(string, string) (string, error)

var stringOperatorFuncList = map[StringOperator]stringOperatorFunc{
	TakeTill: takeTill,
}
