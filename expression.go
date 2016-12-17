package gorules

type Expression interface {
	Evaluate() (bool, error)
}

// type ConjunctionExpression interface {
// 	Add(*Expression)
// }

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

	evaluator, accumlator := ConjunctionExpressionProps(v.Conjunction)

	for _, e := range v.Expressions {
		var isTrue, _ = evaluator(accumlator, (*e))
		accumlator = CreateBoolExpression(isTrue)
	}
	return accumlator.Evaluate()
}

func CreateConjunctionExpression(conjunction Conjunction) func(Expression) Expression {
	return func(expr Expression) Expression {
		conj := &ConjunctionExpression{Conjunction: conjunction}
		conj.Expressions = make([]*Expression, 1)
		conj.Add(&expr)
		return conj
	}
}

var CreateAndConjunctionExpression func(Expression) Expression = CreateConjunctionExpression(And)

var CreateOrConjunctionExpression func(Expression) Expression = CreateConjunctionExpression(Or)

func (conj *ConjunctionExpression) Add(expr *Expression) {
	conj.Expressions = append(conj.Expressions, expr)
}

func IsConjunctionExpression(expr Expression) bool {
	_, ok := expr.(*ConjunctionExpression)
	return ok
}

// func CreateAndExpression(e Expression) *AndExpression {
// 	a := &AndExpression{}
// 	a.expressions = make([]*Expression, 1)
// 	a.Add(&e)
// 	return a
// }

// func CreateOrExpression(e Expression) *OrExpression {
// 	o := &OrExpression{}
// 	o.expressions = make([]*Expression, 1)
// 	o.Add(&e)
// 	return o
// }

// type AndExpression struct {
// 	expressions []*Expression
// }

// func (v AndExpression) Evaluate() (bool, error) {
// 	return true, nil
// }

// func (v *AndExpression) Add(e *Expression) {
// 	v.expressions = append(v.expressions, e)
// }

// type OrExpression struct {
// 	expressions []*Expression
// }

// func (o OrExpression) Evaluate() (bool, error) {
// 	return true, nil
// }

// func (v *OrExpression) Add(e *Expression) {
// 	v.expressions = append(v.expressions, e)
// }

//-------------------------------------

type BoolValueExpression struct {
	Type bool `json:"type"`
}

func (v BoolValueExpression) Evaluate() (bool, error) {
	if v.Type {
		return true, nil
	}
	return false, nil
}

func CreateBoolExpression(bool_type bool) Expression {
	if bool_type {
		return &BoolValueExpression{Type: true}
	}
	return &BoolValueExpression{Type: false}
}

// TrueExpression always evaluates to True
var TrueExpression = CreateBoolExpression(true)

// FalseExpression always evaluates to False
var FalseExpression = CreateBoolExpression(false)

//-------------------------------------
