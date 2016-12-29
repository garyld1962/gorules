package gorules

// Expression refers to anytype that can be evaluated
type Expression interface {
	Evaluate() (bool, error)
}

// ValueExpression stores the value and target be Operated on.Can act with different operates
type ValueExpression struct {
	Operator Operator `json:"operator"`
	Value    string   `json:"value"`
	Target   string   `json:"target"`
}

// Evaluate used to eval an expression to a bool
func (v ValueExpression) Evaluate() (bool, error) {
	operatorFunc := operatorFuncList[v.Operator]
	result, err := operatorFunc(v.Value, v.Target)
	// fmt.Println("Evaluate", v, result)
	return result, err
}

func createValueExpression(operatorText string, value string) Expression {
	operator, err := toOperator(operatorText)
	if err == nil {
		return &ValueExpression{Operator: operator, Value: value}
	}
	panic(err)
}

func createValueExpressionWithTarget(operatorText string, value string, target string) Expression {
	operator, err := toOperator(operatorText)
	if err == nil {
		return &ValueExpression{Operator: operator, Value: value, Target: target}
	}
	panic(err)
}

func createValueExpressionFromRuleStmt(rule *RuleStatement, data map[string]interface{}) Expression {
	// fmt.Println("source", rule)
	source, _ := rule.Source.Evaluate(data)
	target, _ := rule.Target.Evaluate(data)

	return createValueExpressionWithTarget(rule.Operator, decodeSpace(source), decodeSpace(target))
}

// ConjunctionExpression used to combine any type of Expressions
type ConjunctionExpression struct {
	Conjunction Conjunction   `json:"conjunction"`
	Expressions []*Expression `json:"expressions"`
}

// Evaluate used to get the combined evaluated value of all the expressions using the Conjunction
func (c ConjunctionExpression) Evaluate() (bool, error) {

	evaluator, accumlator := conjunctionExprProperties(c.Conjunction)
	// fmt.Println(c.Expressions)
	for _, e := range c.Expressions {
		var resultBool, _ = evaluator(accumlator, (*e))
		accumlator = createBoolExpression(resultBool)
	}
	return accumlator.Evaluate()
}

// Add the expression to be evaluated into the conjunction espression
func (c *ConjunctionExpression) Add(expr *Expression) {
	c.Expressions = append(c.Expressions, expr)
}

var createAndConjunctionExpression = createConjunctionExpression(And)

var createOrConjunctionExpression = createConjunctionExpression(Or)

func createConjunctionExpression(conjunction Conjunction) func(*Expression) Expression {
	return func(expr *Expression) Expression {
		conj := ConjunctionExpression{Conjunction: conjunction}
		conj.Add(expr)
		return conj
	}
}

func createConjuntionExprFromCollectionStmt(ruleStmt *RuleStatement, data map[string]interface{}) Expression {
	selector, err := toSelector(ruleStmt.Selector)

	if err != nil {
		panic(err)
	}

	conjExpr := selectorConjExprMap(selector)
	arrayPath, key := getArrayPathAndKey(ruleStmt.Source.String())
	arrayValue := selectValue(data, arrayPath).([]interface{})

	for _, x := range arrayValue {
		valueToCompare := selectValue(x.(map[string]interface{}), key).(string)
		target, _ := ruleStmt.Target.Evaluate(x)
		valExp := createValueExpressionWithTarget(ruleStmt.Operator, valueToCompare, target)

		conjExpr.Add(&valExp)
	}
	return conjExpr
}

func isConjunctionExpression(expr Expression) bool {
	_, ok := expr.(ConjunctionExpression)
	// fmt.Println("conv", x, ok)
	return ok
}

// BoolValueExpression stores either true or false value as an Expression
type BoolValueExpression bool

// Evaluate makes BoolValueExpression implement Expression
func (v BoolValueExpression) Evaluate() (bool, error) {
	if v {
		return true, nil
	}
	return false, nil
}

// createBoolExpression creates a BoolValueExpression with a bool
func createBoolExpression(boolType bool) Expression {
	if boolType {
		return BoolValueExpression(true)
	}
	return BoolValueExpression(false)
}

// TrueExpression always evaluates to True
var TrueExpression = createBoolExpression(true)

// FalseExpression always evaluates to False
var FalseExpression = createBoolExpression(false)

//-------------------------------------
