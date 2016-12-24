package gorules

import "strings"

//Expressionable is the abstraction of any structure that can be converted to Expression
type Expressionable interface {
	ToExpression(interface{}) (Expression, error)
}

// RuleStatement holds a Expression with a Operator which can be parsed and evaluated
type RuleStatement struct {
	Branch   string `json:"branch"`
	Selector string `json:"selector"`
	Path     string `json:"path"`
	Operator string `json:"operator"`
	Target   string `json:"target"`
}

// ToExpression makes the RuleStatement Expressionable
func (r *RuleStatement) ToExpression(data interface{}) (Expression, error) {
	if isSelector(r.Selector) {
		selector, _ := toSelector(r.Selector)
		selectionFunction := selectorFunctions(selector)
		return selectionFunction(r, data.(map[string]interface{})), nil
	}
	_, err := toSelector(r.Selector)
	return nil, err
}

// createRuleStatement creates a RuleStatement with defaults
func createRuleStatement(input string) *RuleStatement {
	parsed := stringSlice(reverse(strings.Split(input, " ")))
	rule := &RuleStatement{Target: parsed.getOrEmpty(0),
		Operator: parsed.getOrEmpty(1),
		Path:     parsed.getOrDefault(2, "data"),
		Selector: parsed.getOrDefault(3, "THIS"),
		Branch:   parsed.getOrDefault(4, "IF")}
	return rule
}

// createRuleStatementFromExisting creates a RuleStatement with defaults
func createRuleStatementFromExisting(existingRule Expressionable, input string) *RuleStatement {
	parsed := stringSlice(reverse(strings.Split(input, " ")))
	var rule *RuleStatement
	if existingRule != nil {
		existingRulevalue := existingRule.(*RuleStatement)
		rule = &RuleStatement{Target: parsed.getOrEmpty(0),
			Operator: parsed.getOrDefault(1, existingRulevalue.Operator),
			Path:     parsed.getOrDefault(2, existingRulevalue.Path),
			Selector: parsed.getOrDefault(3, existingRulevalue.Selector),
			Branch:   parsed.getOrDefault(4, existingRulevalue.Branch)}

	} else {
		rule = createRuleStatement(input)
	}
	return rule
}

// ConjunctionStatement combines two RuleStatements
type ConjunctionStatement struct {
	Conjunction Conjunction `json:"conjunction"`
}

// Parse makes the ConjunctionStatement Expressionable
func (c *ConjunctionStatement) ToExpression(_ interface{}) (Expression, error) {
	switch c.Conjunction {
	case And:
		return createAndConjunctionExpression(&TrueExpression), nil
	case Or:
		return createOrConjunctionExpression(&FalseExpression), nil
	default:
		return createAndConjunctionExpression(&TrueExpression), nil
	}
}

// createConjunctionStatement Creates Conjunction Statement from string
func createConjunctionStatement(input string) *ConjunctionStatement {
	conjunction, err := toConjunction(input)
	if err != nil {
		panic(err)
	}
	return &ConjunctionStatement{Conjunction: conjunction}
}
