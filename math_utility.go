package gorules

import (
	"strings"
)

func multiply(operand1, operand2 int) (int, error) {
	return operand1 * operand2, nil
}

func divide(operand1, operand2 int) (int, error) {
	return operand1 / operand2, nil
}

func add(operand1, operand2 int) (int, error) {
	return operand1 + operand2, nil
}

func subtract(operand1, operand2 int) (int, error) {
	return operand1 - operand2, nil
}

func takeTill(delimiter, input string) (string, error) {
	return strings.Split(input, delimiter)[0], nil
}
