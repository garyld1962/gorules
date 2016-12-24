package gorules

// RuleParser has to be implemented by anything that needs to be converted to a rule
type RuleParser interface {
	Parse(string, map[string]interface{}) *Rule
}

// ruleParserFunc is to enable any function to implemente RuleParser interface
type ruleParserFunc func(string, map[string]interface{}) *Rule

// Parse Make any function that has RuleParserFunc type signature to become RuleParser
func (fn ruleParserFunc) Parse(dslText string, dataAsJSON map[string]interface{}) *Rule {
	return fn(dslText, dataAsJSON)
}

//ParseDSL parses the DSL with space and creates Rule
var ParseDSL = ruleParserFunc(parseStringAsDSL)

func parseStringAsDSL(dslText string, jsonObj map[string]interface{}) *Rule {

	var ruleToEvaluate = &Rule{}
	var latestRuleStmt Expressionable
	var toParse string

	words := spiltWithSpace(dslText)

	for _, word := range words {
		if !isConjunction(word) {
			toParse = concatStrings(toParse, word, " ")
		} else {
			latestRuleStmt = createRuleStmtFromExisting(latestRuleStmt, trim(toParse))
			toParse = ""
			valueExpr, _ := latestRuleStmt.ToExpression(jsonObj)
			conjExpr, _ := createConjunctionStmt(word).ToExpression(jsonObj)
			ruleToEvaluate.Add(&valueExpr)
			ruleToEvaluate.Add(&conjExpr)
		}

	}

	latestRuleStmt = createRuleStmtFromExisting(latestRuleStmt, trim(toParse))
	lastRuleExpr, _ := latestRuleStmt.ToExpression(jsonObj)
	ruleToEvaluate.Add(&lastRuleExpr)
	return ruleToEvaluate
}
