package gorules

import (
	"fmt"
	"strings"

	objects "github.com/stretchr/stew/objects"
)

const dslSeperator = "|"

//ParseDSL parses simple DSL to Rule (array of Expressions)
func ParseDSL(dslText string, data string) *Rule {
	lst := strings.Split(dslText, dslSeperator)
	m, _ := objects.NewMapFromJSON(data)
	var rle = &Rule{}
	var exp Expressionable
	for _, x := range lst {
		if IsConjunction(x) {
			exp = CreateConjunctionStatement(x)
		} else {
			exp = CreateRuleStatement(strings.TrimSpace(x))
		}
		parsed, _ := exp.Parse(m)
		rle.Add(&parsed)
	}
	boo := ReduceRuleToBool(CreateOrConjunctionExpression(&FalseExpression), rle.expressions)
	fmt.Println("Output", boo)
	return rle
}

//GetKeyFromJSON get the value for key in JSON
func GetKeyFromJSON(obj objects.Map, key string) string {
	value := obj.Get(key)
	//fmt.Println(key,value)
	switch x := value.(type) {
	case map[string]interface{}:
		return "JSONObject"
	case []interface{}:
		return "JSONArray"
	case string:
		return x
	case int:
		return fmt.Sprintf("%d", x)
	default:
		return "default"
	}
}

func ReduceRuleToBool(accum Expression, expressions []*Expression) bool {
	if len(expressions) == 0 {
		fmt.Println("accum", accum)
		value, _ := accum.Evaluate()
		return value
	}

	expr := *expressions[0]

	if isConjunctionExpression(expr) {
		conj, _ := expr.(*ConjunctionExpression)
		isTrue, _ := accum.Evaluate()
		boolExpr := CreateBoolExpression(isTrue)
		conj.Add(&boolExpr)
		accum = conj
	} else {
		conj, _ := accum.(*ConjunctionExpression)
		isTrue, _ := expr.Evaluate()
		boolExpr := CreateBoolExpression(isTrue)
		conj.Add(&boolExpr)
		accum = conj
	}

	return ReduceRuleToBool(accum, expressions[1:])
}
