package gorules

type ExpressionFunction func(string, string) (bool, error)

// type Operator string

// const (
// 	IsEqualTo Operator = "IsEqualTo"
// )

var FunctionList map[Operator]ExpressionFunction = map[Operator]ExpressionFunction{
	IsEqualTo: equals,
}
