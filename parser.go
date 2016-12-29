package gorules

import (
	"fmt"
)

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

var ParseDSLWithPrecedence = ruleParserFunc(parseDSLWithPrecedence)

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
func parseDSLWithPrecedence(dslText string, jsonObj map[string]interface{}) Expression {

	var ruleToEvaluate = Rule{}
	var latestRuleStmt Expressionable
	var expr Expression
	var toParse string

	lines := reverse(spiltWithNewLine(dslText))
	fmt.Println("lines", lines)
	for _, line := range lines {

		line = trim(line)
		fmt.Println("line", line)

		if isConjunction(line) {
			conjExpr, _ := createConjunctionStmt(line).ToExpression(jsonObj)
			ruleToEvaluate = ruleToEvaluate.Add(conjExpr)
			ruleToEvaluate = ruleToEvaluate.Add(expr)
			expr = nil
		} else if endsWithConjunction(line) {
			if notEmpty(toParse) {
				toParse = concatStrings(toParse, " ", makeLastWordFirst(line))
				expr = parseDSLToExpr(toParse, jsonObj)
			} else {
				expr = parseDSLToExpr2(expr, line, jsonObj)
			}
			toParse = ""

		} else {
			toParse = line
		}
	}

	if notEmpty(toParse) {
		latestRuleStmt = createRuleStmtFromExisting(latestRuleStmt, trim(toParse))
		lastRuleExpr, _ := latestRuleStmt.ToExpression(jsonObj)
		ruleToEvaluate = ruleToEvaluate.Add(lastRuleExpr)
	}

	if expr != nil {
		ruleToEvaluate = ruleToEvaluate.Add(expr)
	}

	for x, test := range ruleToEvaluate.expressions {
		fmt.Println(x, " Expression", test)
	}

	return ruleToEvaluate
}
func parseDSLToExpr2(seedExpression Expression, dslText string, jsonObj map[string]interface{}) Expression {
	fmt.Println("3", dslText, seedExpression)
	var ruleToEvaluate = Rule{}
	ruleToEvaluate = ruleToEvaluate.Add(seedExpression)
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
		fmt.Println(x, " Expression Inner", test)
	}
	// fmt.Println("3", toParse, ruleToEvaluate)
	return ruleToEvaluate
}

func precedenceParser(accum Rule, doPush bool, lines []string) Expression {
	var rule = Rule{}
	var linetoWorkOn = lines[0]
	var expr = parseDSLToExprNew(line)

	if doPush {

	}

}
