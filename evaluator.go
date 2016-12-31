package gorules

import "fmt"

// RuleEvaluator has to be implemented by anything that needs to be converted to a rule
type RuleEvaluator interface {
	Evaluate(string, map[string]interface{}) bool
}

// ruleEvaluatorFunc is to enable any function to implemente RuleEvaluator interface
type ruleEvaluatorFunc func(RuleFetcher, map[string]interface{}) bool

// Parse Make any function that has RuleParserFunc type signature to become RuleParser
func (fn ruleEvaluatorFunc) Evaluate(rule RuleFetcher, dataAsJSON map[string]interface{}) bool {
	return fn(rule, dataAsJSON)
}

func evaluator(parser ruleParserFunc) func(RuleFetcher, map[string]interface{}) bool {

	return func(rule RuleFetcher, data map[string]interface{}) bool {
		ruleText := rule.Fetch()
		ruleParsed := parser(ruleText, data)
		result, _ := ruleParsed.Evaluate()
		fmt.Println("result", rule, result)
		return result
	}
}

// DSLEvaluator evaluates DSL to a bool with function ParseDSL
var DSLEvaluator = evaluator(ParseDSL)

var EvaluateRules = ruleEvaluatorFunc(evaluator(ParseDSLWithPrecedence))
