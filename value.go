package gorules

import (
	"errors"
	"fmt"
	"strconv"
)

// Value refers to anytype that can be evaluated to a concrete string value
type Value interface {
	Evaluate(interface{}) (string, error)
	String() string
}

// Constant is used to hold the string value
type Constant struct {
	value string
}

// EmptyConstant is the string form of empty Constant value
var EmptyConstant = "''"

// Evaluate returns the string from the Constant
func (c Constant) Evaluate(_ interface{}) (string, error) {
	if startsWithSingleQuotes(c.value) {
		return stringBetweenSingleQuotes(c.value), nil
	}
	return "", errors.New("Not a Constant")
}

// String makes Constant implement Stringer
func (c Constant) String() string {
	return c.value
}

// NewConstant creates new Constant which is within single quotes.Creates an empty string if value has no quotes
func NewConstant(value string) Constant {
	return Constant{value: value}
}

// Path has the JSON Path. Needs data to be evaluated to the final string value
type Path struct {
	jsonPath string
}

// Evaluate returns the string from the Constant
func (p Path) Evaluate(data interface{}) (string, error) {
	// fmt.Println(p, data)
	return selectValue(data.(map[string]interface{}), p.jsonPath).(string), nil
}

// String makes Path implement Stringer
func (p Path) String() string {
	return p.jsonPath
}

// NewPath creates new JSON path which can be evaluated with supplied data
func NewPath(value string) Path {
	return Path{jsonPath: value}
}

// NewValue used to create any of the value type
func NewValue(value string) Value {

	if startsWithSingleQuotes(value) {
		return NewConstant(value)
	} else if startsWithPipe(value) {
		x := spiltWithSpace(stringBetweenPipe(value))[0]
		if isMathOperator(x) {
			return NewMathExpression(decodeSpace(stringBetweenPipe(value)))
		}

		return NewStringExpression(decodeSpace(stringBetweenPipe(value)))
	}
	return NewPath(value)
}

// MathExpression is used to evaluate mathematical expressions on json values
type MathExpression struct {
	operand1 Value
	operand2 Value
	operator MathOperator
}

//NewMathExpression is a wrapper around MathExpression
func NewMathExpression(expression string) MathExpression {
	parsedOperandsAndOperatorValue := StringSlice(spiltWithSpace(expression))
	mathOperator, _ := toMathOperator(parsedOperandsAndOperatorValue.getOrEmpty(0))
	return MathExpression{operand1: NewValue(trim(parsedOperandsAndOperatorValue.getOrDefault(1, EmptyConstant))),
		operand2: NewValue(trim(parsedOperandsAndOperatorValue.getOrDefault(2, EmptyConstant))),
		operator: mathOperator}
}

// Evaluate works out the expression and returns the result as a string
func (m MathExpression) Evaluate(data interface{}) (string, error) {
	operand1, _ := m.operand1.Evaluate(data)
	operand2, _ := m.operand2.Evaluate(data)
	firstOperand, _ := strconv.Atoi(operand1)
	secondOperand, _ := strconv.Atoi(operand2)
	mathOperatorFunc := mathOperatorFuncList[m.operator]
	result, err := mathOperatorFunc(firstOperand, secondOperand)
	return strconv.Itoa(result), err
}

func (m MathExpression) String() string {
	dummyValue, _ := m.operand1.Evaluate(make([]interface{}, 0))
	return dummyValue
}

// StringExpression is used to evaluate strings expressions on json values
type StringExpression struct {
	operand1 Value
	operand2 Value
	operator StringOperator
}

//NewStringExpression is a wrapper around StringExpression
func NewStringExpression(expression string) StringExpression {
	parsedOperandsAndOperatorValue := StringSlice(spiltWithSpace(expression))
	fmt.Println(expression, parsedOperandsAndOperatorValue)
	stringOperator, _ := toStringOperator(parsedOperandsAndOperatorValue.getOrEmpty(0))
	return StringExpression{operand1: NewValue(trim(parsedOperandsAndOperatorValue.getOrDefault(1, EmptyConstant))),
		operand2: NewValue(trim(parsedOperandsAndOperatorValue.getOrDefault(2, EmptyConstant))),
		operator: stringOperator}
}

func (m StringExpression) String() string {
	dummyValue, _ := m.operand1.Evaluate(make([]interface{}, 0))
	return dummyValue
}

// Evaluate works out the expression and returns the result as a string
func (m StringExpression) Evaluate(data interface{}) (string, error) {
	operand1, _ := m.operand1.Evaluate(data)
	operand2, _ := m.operand2.Evaluate(data)
	fmt.Println(operand1, "re", operand2)
	stringOperatorFunc := stringOperatorFuncList[m.operator]
	result, err := stringOperatorFunc(operand1, operand2)
	return result, err
}
