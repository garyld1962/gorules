package gorules

import "fmt"

// Rule is just a collection of expressions
type Rule []Expression

//Evaluate all the expressions in the rule
func (r Rule) Evaluate() (bool, error) {
	// fmt.Println("5", r.expressions)
	result := evaluateExpressions(r)
	return result, nil
}

//Add Expressions to the Rule
func (r Rule) Add(expression Expression) Rule {
	// fmt.Println("4", expression)
	r = append(r, expression)
	// fmt.Println("4", r.expressions[0])
	return r
}

func isRule(expr Expression) bool {
	_, ok := expr.(Rule)
	return ok
}

func reduceExpressions(accum Expression, expressions []Expression) bool {
	if len(expressions) == 0 {
		value, _ := accum.Evaluate()
		if isConjunctionExpression(accum) {
			for _, x := range accum.(ConjunctionExpression).Expressions {
				fmt.Println("Expressions", x)
			}
		}

		fmt.Println("accum.Evaluate", accum, value)
		return value
	}

	expr := firstExpr(expressions)

	fmt.Println("----------------------------------------------")
	fmt.Println("reduceExpressions- Starts", expr)
	fmt.Println("reduceExpressions-accum At START", accum)
	if isConjunctionExpression(expr) {
		fmt.Println("Conjuction Loop", expr)
		conj, _ := expr.(ConjunctionExpression)
		isTrue, _ := accum.Evaluate()
		boolExpr := createBooleanExpression(isTrue)
		conj = conj.Add(boolExpr)
		accum = conj
	} else {
		fmt.Println("Value Loop")
		conj, _ := accum.(ConjunctionExpression)
		isTrue, _ := expr.Evaluate()
		boolExpr := createBooleanExpression(isTrue)
		conj = conj.Add(boolExpr)
		accum = conj
	}
	fmt.Println("reduceExpressions-accum At END", accum)
	fmt.Println("----------------------------------------------")

	return reduceExpressions(accum, expressions[1:])
}

func determineSeedAccum(expressions []Expression) Expression {
	// fmt.Println("5", expressions)
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
