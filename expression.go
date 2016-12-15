package gorules

type Expression interface {
	Evaluate() (bool, error)
}

type ConjunctionExpression interface {
	Add(*Expression)
}

type ValueExpression struct {
	Operator Operator `json:"operator"`
	Path     string   `json:"path"`
	Value    string   `json:"value"`
	Target   string   `json:"target"`
}

// Evaluate ...
func (v ValueExpression) Evaluate() (bool, error) {
	fun := FunctionList[v.Operator]
	result, err := fun(v.Value, v.Target)
	return result, err
}

func CreateValueExpression(operatorText string, path string, value string) *ValueExpression {
	operator, err := ToOperator(operatorText)
	if err == nil {
		return &ValueExpression{Operator: operator, Path: path, Value: value}
	}
	panic(err)
}

func CreateValueExpressionWithTarget(operatorText string, path string, value string, target string) *ValueExpression {
	operator, err := ToOperator(operatorText)
	if err == nil {
		return &ValueExpression{Operator: operator, Path: path, Value: value, Target: target}
	}
	panic(err)
}

func CreateAndExpression(e Expression) *AndExpression {
	a := &AndExpression{}
	a.expressions = make([]*Expression, 1)
	a.Add(&e)
	return a
}

func CreateOrExpression(e Expression) *OrExpression {
	o := &OrExpression{}
	o.expressions = make([]*Expression, 1)
	o.Add(&e)
	return o
}

type AndExpression struct {
	expressions []*Expression
}

func (v AndExpression) Evaluate() (bool, error) {
	return true, nil
}

func (v *AndExpression) Add(e *Expression) {
	v.expressions = append(v.expressions, e)
}

type OrExpression struct {
	expressions []*Expression
}

func (o OrExpression) Evaluate() (bool, error) {
	return true, nil
}

func (v *OrExpression) Add(e *Expression) {
	v.expressions = append(v.expressions, e)
}

//-------------------------------------

type True struct {
}

func (v True) Evaluate() (bool, error) {
	return true, nil
}

type False struct {
}

func (v False) Evaluate() (bool, error) {
	return false, nil
}

//-------------------------------------
