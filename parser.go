package gorules

// ruleParserFunc is a function type for DSL parsing
type ruleParserFunc func(string, map[string]interface{}) Expression

// parseDSL parses the DSL with space and creates Rule
var parseDSL = ruleParserFunc(parseDSLToExpr)

// parseDSLWithPrecedence parses DSL with new lines and precedence
var parseDSLWithPrecedence = ruleParserFunc(precedenceParser)

func parseDSLToExpr(dslText string, jsonObj map[string]interface{}) Expression {
	var ruleToEvaluate = Rule{}
	var latestRuleStmt Expressionable
	var toParse string

	words := splitWithSpace(dslText)

	for _, word := range words {

		word = trim(word)

		if !isConjunction(word) {
			toParse = concatStrings(toParse, word, " ")
		} else {
			if notEmpty(toParse) {
				latestRuleStmt = createRuleStmtFromExisting(latestRuleStmt, trim(toParse))
				valueExpr, _ := latestRuleStmt.ToExpression(jsonObj)
				ruleToEvaluate = ruleToEvaluate.Add(valueExpr)
				toParse = ""
			}
			conjStmt, _ := createConjunctionStmt(word)
			conjExpr, _ := conjStmt.ToExpression(jsonObj)
			ruleToEvaluate = ruleToEvaluate.Add(conjExpr)
		}
	}

	if notEmpty(toParse) {
		latestRuleStmt = createRuleStmtFromExisting(latestRuleStmt, trim(toParse))
		lastRuleExpr, _ := latestRuleStmt.ToExpression(jsonObj)
		ruleToEvaluate = ruleToEvaluate.Add(lastRuleExpr)
	}
	return ruleToEvaluate
}

func precedenceParserRecursive(accum Rule, pushInExistingRule bool, lines []string, jsonObj map[string]interface{}) Expression {

	if len(lines) == 0 {
		return accum
	}

	var lineToWorkOn = first(lines)

	if pushInExistingRule {
		if isConjunction(lineToWorkOn) {
			conjStmt, _ := createConjunctionStmt(lineToWorkOn)
			conjExpr, _ := conjStmt.ToExpression(jsonObj)
			accum = accum.Add(conjExpr)
			return precedenceParserRecursive(accum, true, lines[1:], jsonObj)
		} else if endsWithConjunction(lineToWorkOn) {
			rule := parseDSLToExpr(lineToWorkOn, jsonObj).(Rule)
			rule = rule.Add(precedenceParserRecursive(rule, false, lines[1:], jsonObj))
			return rule
		} else {
			expr := parseDSLToExpr(lineToWorkOn, jsonObj)
			accum = accum.Add(expr)
			return precedenceParserRecursive(accum, true, lines[1:], jsonObj)
		}
	} else {
		rule := parseDSLToExpr(lineToWorkOn, jsonObj).(Rule)
		return rule
	}
}

func precedenceParser(dslText string, jsonObj map[string]interface{}) Expression {
	return precedenceParserRecursive(Rule{}, true, lines(dslText), jsonObj)
}
