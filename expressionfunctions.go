package gorules

type OperatorFunc func(string, string) (bool, error)

var OperatorFuncList map[Operator]OperatorFunc = map[Operator]OperatorFunc{
	IsEqualTo: equals,
}

type ConjunctionFunc func(Expression, Expression) (bool, error)

var conjunctionFuncList map[Conjunction]ConjunctionFunc = map[Conjunction]ConjunctionFunc{
	And: AndEvaluator,
	Or:  OrEvaluator,
}

var IdentityBoolForConjunction map[Conjunction]Expression = map[Conjunction]Expression{
	And: TrueExpression,
	Or:  FalseExpression,
}

func AndEvaluator(expr_one Expression, expr_two Expression) (bool, error) {
	isOneTrue, _ := expr_one.Evaluate()
	isTwoTrue, _ := expr_two.Evaluate()

	return isOneTrue && isTwoTrue, nil
}

func OrEvaluator(expr_one Expression, expr_two Expression) (bool, error) {
	isOneTrue, _ := expr_one.Evaluate()
	isTwoTrue, _ := expr_two.Evaluate()

	return isOneTrue || isTwoTrue, nil
}

func ConjunctionFunction(conjunction Conjunction) ConjunctionFunc {
	return conjunctionFuncList[conjunction]
}

func IdentityBool(conjuntion Conjunction) Expression {
	return IdentityBoolForConjunction[conjuntion]
}

func ConjunctionExpressionProps(conjunction Conjunction) (ConjunctionFunc, Expression) {
	return ConjunctionFunction(conjunction), IdentityBool(conjunction)
}
