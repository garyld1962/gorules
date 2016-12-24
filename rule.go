package gorules

// Rule is just a collection of expressions
type Rule struct {
	expressions []*Expression
}

//Evaluate all the expressions in the rule
func (r Rule) Evaluate() (bool, error) {
	result := evaluateExpressions(r.expressions)
	return result, nil
}

//Add Expressions to the Rule
func (r *Rule) Add(expression *Expression) {
	r.expressions = append(r.expressions, expression)
}

func reduceExpressions(accum Expression, expressions []*Expression) bool {
	if len(expressions) == 0 {
		value, _ := accum.Evaluate()
		return value
	}

	expr := *expressions[0]
	if isConjunctionExpression(expr) {
		conj, _ := expr.(ConjunctionExpression)
		isTrue, _ := accum.Evaluate()
		boolExpr := createBoolExpression(isTrue)
		conj.Add(&boolExpr)
		accum = conj
	} else {
		conj, _ := accum.(ConjunctionExpression)
		isTrue, _ := expr.Evaluate()
		boolExpr := createBoolExpression(isTrue)
		conj.Add(&boolExpr)
		accum = conj
	}

	return reduceExpressions(accum, expressions[1:])
}

func determineSeedAccum(expressions []*Expression) Expression {
	firstExpression := *expressions[0]
	if isConjunctionExpression(firstExpression) {
		conjExpr := firstExpression.(*ConjunctionExpression).Conjunction
		return identityBool(conjExpr)
	}
	return createOrConjunctionExpression(&FalseExpression)
}

func evaluateExpressions(expressions []*Expression) bool {
	seedValue := determineSeedAccum(expressions)
	return reduceExpressions(seedValue, expressions)
}

// // RuleSet ...
// type RuleSet struct {
// 	rules []Rule
// }

// // Add rules to the RuleSet
// func (r *RuleSet) Add(rule Rule) {
// 	r.rules = append(r.rules, rule)
// }

// //Evaluate all the rules in the Set
// func (r RuleSet) Evaluate() (bool, error) {
// 	return true, nil
// }
