package gorules

import "fmt"

// Conjunction type
type Conjunction int

const (
	// And used && to evaluate the Expressions provided
	And Conjunction = iota
	// Or uses || to evaluate the expressions provided
	Or
	maxConjunctionFlag
)

var conjunctionNames = [...]string{
	And: "AND",
	Or:  "OR",
}

// String makes Conjunction implement Stringer interface
func (c Conjunction) String() string {
	if c >= maxConjunctionFlag {
		return "Invalid Conjunction"
	}
	return conjunctionNames[c]
}

// toConjunction converts string to Conjunction
func toConjunction(s string) (Conjunction, error) {
	for i, r := range conjunctionNames {
		if s == r {
			return Conjunction(i), nil
		}
	}
	return maxConjunctionFlag, fmt.Errorf("Invalid Conjunction value %q", s)
}

func isConjunction(value string) bool {
	_, err := toConjunction(value)
	if err == nil {
		return true
	}
	return false
}

type conjunctionFunc func(Expression, Expression) (bool, error)

var conjunctionFuncList = map[Conjunction]conjunctionFunc{
	And: andEvaluator,
	Or:  orEvaluator,
}

var identityBoolForConjunction = map[Conjunction]Expression{
	And: TrueExpression,
	Or:  FalseExpression,
}

func andEvaluator(exprOne Expression, exprTwo Expression) (bool, error) {
	isOneTrue, _ := exprOne.Evaluate()
	isTwoTrue, _ := exprTwo.Evaluate()

	return isOneTrue && isTwoTrue, nil
}

func orEvaluator(exprOne Expression, exprTwo Expression) (bool, error) {
	isOneTrue, _ := exprOne.Evaluate()
	isTwoTrue, _ := exprTwo.Evaluate()

	return isOneTrue || isTwoTrue, nil
}

func conjunctionFunction(conjunction Conjunction) conjunctionFunc {
	return conjunctionFuncList[conjunction]
}

func identityBool(conjuntion Conjunction) Expression {
	return identityBoolForConjunction[conjuntion]
}

// conjunctionExprProperties returns the conjuntion function used for evaluation and the seed value
func conjunctionExprProperties(conjunction Conjunction) (conjunctionFunc, Expression) {
	return conjunctionFunction(conjunction), identityBool(conjunction)
}

func endsWithConjunction(input string) bool {

	strList := spiltWithSpace(trim(input))

	if len(strList) > 1 {
		lstElmt := lastElement(strList)

		if isConjunction(lstElmt) {
			return true
		}

	}
	return false
}
