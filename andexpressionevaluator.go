package gorules


type AndExpressionEvaluator struct {
}

func (a AndExpressionEvaluator) Evaluate(ex Expression) (bool, error) {
	return true, nil
}