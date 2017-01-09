package gorules

import "fmt"

//Expressionable is the abstraction of any structure that can be converted to Expression
type Expressionable interface {
	ToExpression(interface{}) (Expression, error)
}

// RuleStatement holds a Expression with a Operator which can be parsed and evaluated
type RuleStatement struct {
	Branch   string `json:"branch"`
	Selector string `json:"selector"`
	Source   Value  `json:"source"`
	Operator string `json:"operator"`
	Target   Value  `json:"target"`
}

// ToExpression makes the RuleStatement Expressionable
func (r RuleStatement) ToExpression(data interface{}) (Expression, error) {
	//fmt.Println("Statement", *r)
	selector, err := toSelector(r.Selector)
	if err != nil {
		return nil, err
	}
	selectionFunction := selectorFunctions(selector)
	return selectionFunction(r, data.(map[string]interface{})), nil
}

// createRuleStmt creates a RuleStatement with defaults
func createRuleStmt(input string) RuleStatement {

	parsed := StringSlice(reverse(spiltWithSpace(encodeString(input))))
	fmt.Println("createRuleStmt", input, parsed, encodeString(input), NewValue(parsed.getOrEmpty(0)))
	ruleStmt := RuleStatement{
		Target:   NewValue(parsed.getOrEmpty(0)),
		Source:   NewValue(parsed.getOrDefault(2, "data")),
		Operator: parsed.getOrEmpty(1),
		Selector: parsed.getOrDefault(3, "THIS"),
		Branch:   parsed.getOrDefault(4, "IF")}
	fmt.Println("ruleStmt", ruleStmt)
	return ruleStmt
}

// createRuleStmtFromExisting creates a RuleStatement and fills the missing values from the existingRule provided
func createRuleStmtFromExisting(existingRule Expressionable, input string) RuleStatement {

	parsed := StringSlice(reverse(spiltWithSpace(encodeString(input))))
	var rule RuleStatement
	if existingRule != nil {
		existRuleVal := existingRule.(RuleStatement)
		rule = RuleStatement{
			Target:   NewValue(parsed.getOrEmpty(0)),
			Source:   NewValue(parsed.getOrDefault(2, existRuleVal.Source.String())),
			Operator: parsed.getOrDefault(1, existRuleVal.Operator),
			Selector: parsed.getOrDefault(3, existRuleVal.Selector),
			Branch:   parsed.getOrDefault(4, existRuleVal.Branch)}

	} else {
		rule = createRuleStmt(input)
	}
	fmt.Println("rule", rule)
	return rule
}

// ConjunctionStatement combines two RuleStatements
type ConjunctionStatement struct {
	Conjunction Conjunction `json:"conjunction"`
}

// ToExpression makes the ConjunctionStatement Expressionable
func (c ConjunctionStatement) ToExpression(_ interface{}) (Expression, error) {
	switch c.Conjunction {
	case And:
		return createAndConjunctionExpression(TrueExpression), nil
	case Or:
		return createOrConjunctionExpression(FalseExpression), nil
	default:
		return createAndConjunctionExpression(TrueExpression), nil
	}
}

// createConjunctionStmt Creates Conjunction Statement from string
func createConjunctionStmt(input string) ConjunctionStatement {
	conjunction, err := toConjunction(input)
	if err != nil {
		panic(err)
	}
	return ConjunctionStatement{Conjunction: conjunction}
}
