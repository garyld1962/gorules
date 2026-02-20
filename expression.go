package gorules

import "fmt"

// Expression refers to anytype that can be evaluated
type Expression interface {
	Evaluate() (bool, error)
}

// RuleExpression stores the value and target be Operated on.Can act with different operates
type RuleExpression struct {
	Operator Operator `json:"operator"`
	Value    string   `json:"value"`
	Target   string   `json:"target"`
}

// Evaluate used to eval an expression to a bool
func (v RuleExpression) Evaluate() (bool, error) {
	operatorFunc, ok := operatorFuncList[v.Operator]
	if !ok || operatorFunc == nil {
		return false, fmt.Errorf("no implementation for operator %s", v.Operator)
	}
	result, err := operatorFunc(v.Value, v.Target)
	return result, err
}

func createRuleExpression(operatorText string, value string) (Expression, error) {
	operator, err := toOperator(operatorText)
	if err != nil {
		return FalseExpression, fmt.Errorf("createRuleExpression: %w", err)
	}
	return RuleExpression{Operator: operator, Value: value}, nil
}

func createRuleExpressionWithTarget(operatorText string, value string, target string) (Expression, error) {
	operator, err := toOperator(operatorText)
	if err != nil {
		return FalseExpression, fmt.Errorf("createRuleExpressionWithTarget: %w", err)
	}
	return RuleExpression{Operator: operator, Value: value, Target: target}, nil
}

func createRuleExpressionFromRuleStmt(rule RuleStatement, data map[string]interface{}) Expression {

	source, _ := rule.Source.Evaluate(data)
	target, _ := rule.Target.Evaluate(data)
	expr, _ := createRuleExpressionWithTarget(rule.Operator, decodeSpace(source), decodeSpace(target))
	return expr
}

// ConjunctionExpression used to combine any type of Expressions
type ConjunctionExpression struct {
	Conjunction Conjunction  `json:"conjunction"`
	Expressions []Expression `json:"expressions"`
}

// Evaluate used to get the combined evaluated value of all the expressions using the Conjunction
func (c ConjunctionExpression) Evaluate() (bool, error) {

	evaluator, accumulator := conjunctionExprProperties(c.Conjunction)
	for _, e := range c.Expressions {
		var resultBool, _ = evaluator(accumulator, (e))
		accumulator = createBooleanExpression(resultBool)
	}
	return accumulator.Evaluate()
}

// Add the expression to be evaluated into the conjunction espression
func (c ConjunctionExpression) Add(expr Expression) ConjunctionExpression {
	c.Expressions = append(c.Expressions, expr)
	return c
}

var createAndConjunctionExpression = createConjunctionExpression(And)

var createOrConjunctionExpression = createConjunctionExpression(Or)

func createConjunctionExpression(conjunction Conjunction) func(Expression) Expression {
	return func(expr Expression) Expression {
		conj := ConjunctionExpression{Conjunction: conjunction}
		conj = conj.Add(expr)
		return conj
	}
}

func createConjunctionExprFromCollectionStmt(ruleStmt RuleStatement, data map[string]interface{}) Expression {
	selector, err := toSelector(ruleStmt.Selector)
	if err != nil {
		return FalseExpression
	}

	conjExpr := selectorConjExprMap(selector)
	arrayPath, key := getArrayPathAndKey(ruleStmt.Source.String())
	rawArray := selectValue(data, arrayPath)
	arrayValue, ok := rawArray.([]interface{})
	if !ok {
		return FalseExpression
	}

	for _, x := range arrayValue {
		xMap, ok := x.(map[string]interface{})
		if !ok {
			continue
		}
		rawVal := selectValue(xMap, key)
		var valueToCompare string
		if str, ok := rawVal.(string); ok {
			valueToCompare = str
		} else {
			valueToCompare = fmt.Sprintf("%v", rawVal)
		}
		target, _ := ruleStmt.Target.Evaluate(x)
		valExp, _ := createRuleExpressionWithTarget(ruleStmt.Operator, valueToCompare, target)
		conjExpr = conjExpr.Add(valExp)
	}
	return conjExpr
}

func isConjunctionExpression(expr Expression) bool {
	_, ok := expr.(ConjunctionExpression)

	return ok
}

// BooleanExpression stores either true or false value as an Expression
type BooleanExpression bool

// Evaluate makes BooleanExpression implement Expression
func (v BooleanExpression) Evaluate() (bool, error) {
	if v {
		return true, nil
	}
	return false, nil
}

// createBooleanExpression creates a BooleanExpression with a bool
func createBooleanExpression(boolType bool) Expression {
	if boolType {
		return BooleanExpression(true)
	}
	return BooleanExpression(false)
}

// TrueExpression always evaluates to True
var TrueExpression = createBooleanExpression(true)

// FalseExpression always evaluates to False
var FalseExpression = createBooleanExpression(false)

//-------------------------------------
