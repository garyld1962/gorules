package gorules

func evaluator(parser ruleParserFunc) func(string, map[string]interface{}) bool {

	return func(rule string, data map[string]interface{}) bool {
		ruleParsed := parser(rule, data)
		result, _ := ruleParsed.Evaluate()
		return result
	}
}

var dslEvaluator = evaluator(parseDSL)
var evaluateRulesFunc = evaluator(parseDSLWithPrecedence)

// DSLEvaluator evaluates DSL to a bool with function parseDSL
func DSLEvaluator(rule string, data map[string]interface{}) bool {
	return dslEvaluator(rule, data)
}

// EvaluateRules evaluates DSL with precedence parsing
func EvaluateRules(rule string, data map[string]interface{}) bool {
	return evaluateRulesFunc(rule, data)
}
