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
		fmt.Println("pared", parsed)
		rle.Add(&parsed)
	}
	fmt.Println(rle)
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

// func selectValue(m map[string]interface{}, path string) interface{} {

// 	propertyNames := strings.Split(path, ".")
// 	// we only have one path element so return what we have
// 	if len(propertyNames) == 1 {
// 		return m[propertyNames[0]]
// 	}
// 	// create map using propertyName as a key and the []interface{} as a value
// 	m1 := make(map[string]interface{})
// 	m1[propertyNames[0]] = m[propertyNames[0]]
// 	// create a new path with the remainder of the parts.
// 	newPath := strings.Join(propertyNames[2:], ",")
// 	return selectValue(m1, newPath)
// }

// func ReduceRuleToBool(accum Expression, expressions []*Expression) bool {
// 	var expr Expression
// 	if len(expressions) == 0 {
// 		value, _ := accum.Evaluate()
// 		return value
// 	}

// 	expr = expressions[0]
// 	if IsConjunctionExpression(expr) {

// 	}

// 	// var result bool
// 	// var current_conjunction ConjunctionExpression
// 	// accum = CreateOrExpression(False{})

// 	// for _, e := range rule.expressions {
// 	// 	expr := *e
// 	// 	if IsConjunctionExpression(expr) {
// 	// 		result, _ = accum.Evaluate()
// 	// 		if result {
// 	// 			accum = True{}
// 	// 		} else {
// 	// 			accum = False{}
// 	// 		}
// 	// 	} else {
// 	// 		accum.Add(&expr)
// 	// 	}
// 	// }
// 	return true
// }
