package gorules

type OrExpressionEvaluator struct {
}

func (o OrExpressionEvaluator) Evaluate(ex Expression) (bool, error) {
	return true, nil
}
