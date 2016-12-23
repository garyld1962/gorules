package gorules

// RuleSet ...
type RuleSet struct {
	rules []Rule
}

func (r *RuleSet) Add(rule Rule) {
	r.rules = append(r.rules, rule)
}

func (r RuleSet) Evalute() (bool, error) {
	return true, nil
}

type Rule struct {
	expressions []*Expression
}

func (r Rule) Evalute() (bool, error) {
	result := EvaluateExpressions(CreateOrConjunctionExpression(&FalseExpression), r.expressions)
	return result, nil
}

func (r *Rule) Add(expression *Expression) {
	r.expressions = append(r.expressions, expression)
}
