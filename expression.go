package gorules

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
	operatorFunc := operatorFuncList[v.Operator]
	result, err := operatorFunc(v.Value, v.Target)
	// fmt.Println("Evaluate", v, result)
	return result, err
}

func createRuleExpression(operatorText string, value string) Expression {
	operator, err := toOperator(operatorText)
	if err == nil {
		return RuleExpression{Operator: operator, Value: value}
	}
	panic(err)
}

func createRuleExpressionWithTarget(operatorText string, value string, target string) Expression {
	operator, err := toOperator(operatorText)
	if err == nil {
		return RuleExpression{Operator: operator, Value: value, Target: target}
	}
	panic(err)
}

func createRuleExpressionFromRuleStmt(rule RuleStatement, data map[string]interface{}) Expression {
	// fmt.Println("source", rule)
	source, _ := rule.Source.Evaluate(data)
	target, _ := rule.Target.Evaluate(data)

	return createRuleExpressionWithTarget(rule.Operator, decodeSpace(source), decodeSpace(target))
}

// ConjunctionExpression used to combine any type of Expressions
type ConjunctionExpression struct {
	Conjunction Conjunction  `json:"conjunction"`
	Expressions []Expression `json:"expressions"`
}

// Evaluate used to get the combined evaluated value of all the expressions using the Conjunction
func (c ConjunctionExpression) Evaluate() (bool, error) {

	evaluator, accumlator := conjunctionExprProperties(c.Conjunction)
	// fmt.Println(c.Expressions)
	for _, e := range c.Expressions {
		var resultBool, _ = evaluator(accumlator, (e))
		accumlator = createBooleanExpression(resultBool)
	}
	return accumlator.Evaluate()
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

func createConjuntionExprFromCollectionStmt(ruleStmt RuleStatement, data map[string]interface{}) Expression {
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
		valExp := createRuleExpressionWithTarget(ruleStmt.Operator, valueToCompare, target)

		conjExpr = conjExpr.Add(valExp)
	}
	return conjExpr
}

func isConjunctionExpression(expr Expression) bool {
	_, ok := expr.(ConjunctionExpression)
	// fmt.Println("conv", x, ok)
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
