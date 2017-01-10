package gorules

import "strings"

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

func square(operand, _ int) (int, error) {
	return operand * operand, nil
}

func takeTill(delimiter, input string) (string, error) {
	return strings.Split(input, delimiter)[0], nil
}

func toUpper(input, _ string) (string, error) {
	return strings.ToUpper(input), nil
}

func toLower(input, _ string) (string, error) {
	return strings.ToLower(input), nil
}
