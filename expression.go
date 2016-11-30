package gorules

type Expression interface {
	Evaluate() (bool, error)
}

type ValueExpression struct {
	Operator string `json:"operator"`
	Path     string `json:"path"`
	Value    string `json:"value"`
	Target   string `json:"target"`
}

// Evaluate ...
func (v ValueExpression) Evaluate() (bool, error) {
	return true, nil
}

func CreateValueExpression(operator string, path string, value string) *ValueExpression {
	expression := &ValueExpression{Operator: operator, Path: path, Value: value}
	return expression
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

func (v AndExpression) Evaluate(target interface{}) (bool, error) {
	return true, nil
}

func (v AndExpression) Add(e *Expression) {
	v.expressions = append(v.expressions, e)
}

type OrExpression struct {
	expressions []*Expression
}

func (o OrExpression) Evaluate(target interface{}) (bool, error) {
	return true, nil
}

func (v OrExpression) Add(e *Expression) {
	v.expressions = append(v.expressions, e)
}
