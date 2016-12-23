package gorules

type Expression interface {
	Evaluate() (bool, error)
}

type ValueExpression struct {
	Operator Operator `json:"operator"`
	Path     string   `json:"path"`
	Value    string   `json:"value"`
	Target   string   `json:"target"`
}

// Evaluate ...
func (v ValueExpression) Evaluate() (bool, error) {
	operatorFunc := OperatorFuncList[v.Operator]
	result, err := operatorFunc(v.Value, v.Target)
	return result, err
}

func CreateValueExpression(operatorText string, path string, value string) Expression {
	operator, err := ToOperator(operatorText)
	if err == nil {
		return &ValueExpression{Operator: operator, Path: path, Value: value}
	}
	panic(err)
}

func CreateValueExpressionWithTarget(operatorText string, path string, value string, target string) Expression {
	operator, err := ToOperator(operatorText)
	if err == nil {
		return &ValueExpression{Operator: operator, Path: path, Value: value, Target: target}
	}
	panic(err)
}

type ConjunctionExpression struct {
	Conjunction Conjunction   `json:"conjunction"`
	Expressions []*Expression `json:"expressions"`
}

// Evaluate ...
func (v ConjunctionExpression) Evaluate() (bool, error) {

	evaluator, accumlator := ConjunctionExprProperties(v.Conjunction)

	for _, e := range v.Expressions {
		var resultBool, _ = evaluator(accumlator, (*e))
		accumlator = CreateBoolExpression(resultBool)
	}
	return accumlator.Evaluate()
}

func CreateConjunctionExpression(conjunction Conjunction) func(*Expression) Expression {
	return func(expr *Expression) Expression {
		conj := &ConjunctionExpression{Conjunction: conjunction}
		conj.Add(expr)
		return conj
	}
}

var CreateAndConjunctionExpression func(*Expression) Expression = CreateConjunctionExpression(And)

var CreateOrConjunctionExpression func(*Expression) Expression = CreateConjunctionExpression(Or)

func (conj *ConjunctionExpression) Add(expr *Expression) {
	conj.Expressions = append(conj.Expressions, expr)
}

func isConjunctionExpression(expr Expression) bool {
	_, ok := expr.(*ConjunctionExpression)
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

// CreateBoolExpression creates a BoolValueExpression with a bool
func CreateBoolExpression(boolType bool) Expression {
	if boolType {
		return BoolValueExpression(true)
	}
	return BoolValueExpression(false)
}

// TrueExpression always evaluates to True
var TrueExpression = CreateBoolExpression(true)

// FalseExpression always evaluates to False
var FalseExpression = CreateBoolExpression(false)

//-------------------------------------
