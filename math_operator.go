package gorules

import "fmt"

// MathOperator to hold various arithematic operations
type MathOperator int

const (
	//Multiply two operands
	Multiply MathOperator = iota
	// Divide two operands
	Divide
	// Add two operands
	Add
	// Subtract two operands
	Subtract
	maxMathOperatorFlag
)

var mathOperatorNames = map[MathOperator]string{
	Multiply: "MUL",
	Divide:   "DIV",
	Add:      "ADD",
	Subtract: "SUB",
}

// String makes MathOperator implement Stringer
func (o MathOperator) String() string {
	if o >= maxMathOperatorFlag {
		return "Invalid Math Operator"
	}
	return mathOperatorNames[o]
}

// toMathOperator converts string to one of the operatorNames in const declaration
func toMathOperator(s string) (MathOperator, error) {
	fmt.Println("to Math Operator invoked ", s)
	for i, r := range mathOperatorNames {
		if s == r {
			return MathOperator(i), nil
		}
	}
	return maxMathOperatorFlag, fmt.Errorf("Invalid Math Operator value %q", s)
}

// isOperator check if a string is a Operator Names
func isMathOperator(value string) bool {
	_, err := toMathOperator(value)
	if err == nil {
		return true
	}
	return false
}

// MathOperatorList returns the Operators
func MathOperatorList() []string {
	operatorArray := make([]string, 0)
	for _, r := range mathOperatorNames {
		operatorArray = append(operatorArray, r)
	}
	return operatorArray
}

type mathOperatorFunc func(int, int) (int, error)

var mathOperatorFuncList = map[MathOperator]mathOperatorFunc{
	Multiply: multiply,
	Divide:   divide,
	Add:      add,
	Subtract: subtract,
}
