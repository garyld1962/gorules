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
