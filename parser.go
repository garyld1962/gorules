package gorules

import (
	"fmt"
	objects "github.com/stretchr/stew/objects"
	"strings"
)

func ParseDSL(dslExpr string, data string) []string {
	lst := strings.Split(dslExpr, "|")
	m, _ := objects.NewMapFromJSON(data)
	var rle =&Rule{}
	var exp Expressionable
	for l, x := range lst {
                // Even elements are rules; odd are conjunctions
		if isEven(l) {
			exp = CreateRuleStatement(strings.TrimSpace(x))
		} else {
			exp = CreateConjunctionStatement(x)
		}
		parsed, _ := exp.Parse(m)
                fmt.Println("pared",parsed)
		rle.Add(&parsed)
	}
        fmt.Println(rle)
	return lst
}

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
