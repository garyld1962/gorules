package gorules

import "fmt"

func evaluator(parser ruleParserFunc) func(string, string) bool {

	return func(dslText string, data string) bool {
		rule := parser(dslText, parseStringToJSONObject(data))
		fmt.Println("result", rule)
		result, _ := rule.Evaluate()
		return result
	}
}

// DSLEvaluator evaluates DSL to a bool with function ParseDSL
var DSLEvaluator = evaluator(ParseDSL)

var DSLEvaluatorWithP = evaluator(ParseDSLWithPrecedence)
