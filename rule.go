package gorules

// Rule is just a collection of expressions
type Rule []Expression

// Evaluate all the expressions in the rule
func (r Rule) Evaluate() (bool, error) {
	result := evaluateExpressions(r)
	return result, nil
}

// Add Expressions to the Rule
func (r Rule) Add(expression Expression) Rule {
	r = append(r, expression)
	return r
}

func isRule(expr Expression) bool {
	_, ok := expr.(Rule)
	return ok
}

func reduceExpressions(accum Expression, expressions []Expression) bool {
	if len(expressions) == 0 {
		value, _ := accum.Evaluate()
		return value
	}

	expr := firstExpr(expressions)

	if isConjunctionExpression(expr) {
		conj, _ := expr.(ConjunctionExpression)
		isTrue, _ := accum.Evaluate()
		boolExpr := createBooleanExpression(isTrue)
		conj = conj.Add(boolExpr)
		accum = conj
	} else {
		conj, _ := accum.(ConjunctionExpression)
		isTrue, _ := expr.Evaluate()
		boolExpr := createBooleanExpression(isTrue)
		conj = conj.Add(boolExpr)
		accum = conj
	}

	return reduceExpressions(accum, expressions[1:])
}

func determineSeedAccum(expressions []Expression) Expression {
	firstExpression := firstExpr(expressions)
	if isConjunctionExpression(firstExpression) {
		conjExpr := firstExpression.(ConjunctionExpression).Conjunction
		return identityBool(conjExpr)
	}
	return createOrConjunctionExpression(FalseExpression)
}

func evaluateExpressions(expressions []Expression) bool {
	seedValue := determineSeedAccum(expressions)
	return reduceExpressions(seedValue, expressions)
}
