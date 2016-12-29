package gorules

import "fmt"

// RuleParser has to be implemented by anything that needs to be converted to a rule
type RuleParser interface {
	Parse(string, map[string]interface{}) Expression
}

// ruleParserFunc is to enable any function to implemente RuleParser interface
type ruleParserFunc func(string, map[string]interface{}) Expression

// Parse Make any function that has RuleParserFunc type signature to become RuleParser
func (fn ruleParserFunc) Parse(dslText string, dataAsJSON map[string]interface{}) Expression {
	return fn(dslText, dataAsJSON)
}

//ParseDSL parses the DSL with space and creates Rule
var ParseDSL = ruleParserFunc(parseDSLToExpr)

var ParseDSLWithPrecedence = ruleParserFunc(callPrecedenceParser)

func parseDSLToExpr(dslText string, jsonObj map[string]interface{}) Expression {
	fmt.Println("2", dslText)
	var ruleToEvaluate = Rule{}
	var latestRuleStmt Expressionable
	var toParse string

	words := spiltWithSpace(dslText)

	for _, word := range words {

		word = trim(word)

		if !isConjunction(word) {
			toParse = concatStrings(toParse, word, " ")
			fmt.Println("toParse", word, toParse)
		} else {
			if notEmpty(toParse) {
				latestRuleStmt = createRuleStmtFromExisting(latestRuleStmt, trim(toParse))
				fmt.Println("latestRuleStmt", latestRuleStmt)
				valueExpr, _ := latestRuleStmt.ToExpression(jsonObj)
				ruleToEvaluate = ruleToEvaluate.Add(valueExpr)
				toParse = ""
			}
			conjExpr, _ := createConjunctionStmt(word).ToExpression(jsonObj)
			fmt.Println("conjExpr", conjExpr)
			ruleToEvaluate = ruleToEvaluate.Add(conjExpr)
		}
		// fmt.Println("2", ruleToEvaluate)
	}

	if notEmpty(toParse) {
		latestRuleStmt = createRuleStmtFromExisting(latestRuleStmt, trim(toParse))
		lastRuleExpr, _ := latestRuleStmt.ToExpression(jsonObj)
		// fmt.Println("1", lastRuleExpr)
		ruleToEvaluate = ruleToEvaluate.Add(lastRuleExpr)
	}
	for x, test := range ruleToEvaluate.expressions {
		fmt.Println(x, "Expression Inner", test)
	}
	out, _ := ruleToEvaluate.Evaluate()
	fmt.Println("3", ruleToEvaluate, out)
	return ruleToEvaluate
}

func precedenceParser(accum Rule, doPush bool, lines []string, jsonObj map[string]interface{}) Expression {

	if len(lines) == 0 {
		return accum
	}

	var linetoWorkOn = first(lines)
	if doPush {
		if isConjunction(linetoWorkOn) {
			conjExpr, _ := createConjunctionStmt(linetoWorkOn).ToExpression(jsonObj)
			accum = accum.Add(conjExpr)
			return precedenceParser(accum, true, lines[1:], jsonObj)
		} else if endsWithConjunction(linetoWorkOn) {
			rule := parseDSLToExpr(linetoWorkOn, jsonObj).(Rule)
			rule = rule.Add(precedenceParser(rule, false, lines[1:], jsonObj))
			return rule
		} else {
			expr := parseDSLToExpr(linetoWorkOn, jsonObj)
			accum = accum.Add(expr)
			return precedenceParser(accum, true, lines[1:], jsonObj)
		}
	} else {
		rule := parseDSLToExpr(linetoWorkOn, jsonObj).(Rule)
		return rule
	}
}

func callPrecedenceParser(dslText string, jsonObj map[string]interface{}) Expression {
	return precedenceParser(Rule{}, true, lines(dslText), jsonObj)
}
