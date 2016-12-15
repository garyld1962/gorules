package gorules

type OperatorFunc func(string, string) (bool, error)

var OperatorFuncList map[Operator]OperatorFunc = map[Operator]OperatorFunc{
	IsEqualTo: equals,
}

type ConjunctionFunc func(Expression, Expression) (bool, error)

var ConjunctionFuncList map[Conjunction]ConjunctionFunc = map[Conjunction]ConjunctionFunc{
	And: AndEvaluator,
}

func AndEvaluator(expr_one Expression, expr_two Expression) (bool, error) {
	return true, nil
}
