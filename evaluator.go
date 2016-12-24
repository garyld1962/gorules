package gorules

func evaluator(parser ruleParserFunc) func(string, string) bool {

	return func(dslText string, data string) bool {
		result, _ := parser(dslText, parseStringToJSONObject(data)).Evaluate()
		// fmt.Println("result", result)
		return result
	}
}

// DSLEvaluator evaluates DSL to a bool with function ParseDSL
var DSLEvaluator = evaluator(ParseDSL)
