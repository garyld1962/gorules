package gorules

import "fmt"

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

var conjunctionFuncList map[Conjunction]conjunctionFunc = map[Conjunction]conjunctionFunc{
	And: andEvaluator,
	Or:  orEvaluator,
}

var IdentityBoolForConjunction map[Conjunction]Expression = map[Conjunction]Expression{
	And: TrueExpression,
	Or:  FalseExpression,
}

func andEvaluator(expr_one Expression, expr_two Expression) (bool, error) {
	isOneTrue, _ := expr_one.Evaluate()
	isTwoTrue, _ := expr_two.Evaluate()

	return isOneTrue && isTwoTrue, nil
}

func orEvaluator(expr_one Expression, expr_two Expression) (bool, error) {
	isOneTrue, _ := expr_one.Evaluate()
	isTwoTrue, _ := expr_two.Evaluate()

	return isOneTrue || isTwoTrue, nil
}

func conjunctionFunction(conjunction Conjunction) conjunctionFunc {
	return conjunctionFuncList[conjunction]
}

func identityBool(conjuntion Conjunction) Expression {
	return IdentityBoolForConjunction[conjuntion]
}

// conjunctionExprProperties returns the conjuntion function used for evaluation and the seed value
func conjunctionExprProperties(conjunction Conjunction) (conjunctionFunc, Expression) {
	return conjunctionFunction(conjunction), identityBool(conjunction)
}
