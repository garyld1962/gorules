package gorules

import "fmt"

// StringOperator to hold various String operations
type StringOperator int

const (
	//TakeTill two operands
	TakeTill StringOperator = iota
	//ToUpper makes the string to uppercase
	ToUpper
	//ToLower makes the string to lowercase
	ToLower
	maxStringOperatorFlag
)

var stringOperatorNames = [...]string{
	TakeTill: "TAKETILL",
	ToUpper:  "TOUPPER",
	ToLower:  "TOLOWER",
}

// String makes StringOperator implement Stringer
func (o StringOperator) String() string {
	if o >= maxStringOperatorFlag {
		return "Invalid String Operator"
	}
	return stringOperatorNames[o]
}

// toStringOperator converts string to one of the operatorNames in const declaration
func toStringOperator(s string) (StringOperator, error) {
	for i, r := range stringOperatorNames {
		if s == r {
			return StringOperator(i), nil
		}
	}
	return maxStringOperatorFlag, fmt.Errorf("Invalid String Operator value %q", s)
}

// isStringOperator checks if a string is a StringOperator name
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
	ToUpper:  toUpper,
	ToLower:  toLower,
}
