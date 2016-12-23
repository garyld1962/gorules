package gorules

import "strings"
import "fmt"

//Expressionable is the abstraction of any structure that can be converted to Expression
type Expressionable interface {
	Parse(interface{}) (Expression, error)
}

// RuleStatement holds a Expression with a Operator which can be parsed and evaluated
type RuleStatement struct {
	Branch   string `json:"branch"`
	Selector string `json:"selector"`
	Path     string `json:"path"`
	Operator string `json:"operator"`
	Target   string `json:"target"`
}

// Parse makes the RuleStatement Expressionable
func (ruleStatement *RuleStatement) Parse(data interface{}) (Expression, error) {
	if IsSelector(ruleStatement.Selector) {
		selector, _ := ToSelector(ruleStatement.Selector)
		selectionFunction := selectorFunctions(selector)
		return selectionFunction(ruleStatement, data), nil
	}
	_, err := ToSelector(ruleStatement.Selector)
	return nil, err
}

// CreateRuleStatement creates a RuleStatement with defaults
func CreateRuleStatement(input string) *RuleStatement {
	parsed := StrSlice(reverse(strings.Split(input, " ")))
	rule := &RuleStatement{Target: parsed.GetOrEmpty(0),
		Operator: parsed.GetOrEmpty(1),
		Path:     parsed.GetOrDefault(2, "data"),
		Selector: parsed.GetOrDefault(3, "THIS"),
		Branch:   parsed.GetOrDefault(4, "IF")}
	fmt.Println(parsed)
	return rule
}

// CreateRuleStatementFromExisting creates a RuleStatement with defaults
func CreateRuleStatementFromExisting(existingRule Expressionable, input string) *RuleStatement {
	parsed := StrSlice(reverse(strings.Split(input, " ")))
	var rule *RuleStatement
	if existingRule != nil {
		existingRulevalue := existingRule.(*RuleStatement)
		rule = &RuleStatement{Target: parsed.GetOrEmpty(0),
			Operator: parsed.GetOrDefault(1, existingRulevalue.Operator),
			Path:     parsed.GetOrDefault(2, existingRulevalue.Path),
			Selector: parsed.GetOrDefault(3, existingRulevalue.Selector),
			Branch:   parsed.GetOrDefault(4, existingRulevalue.Branch)}

	} else {
		rule = CreateRuleStatement(input)
	}
	return rule
}

// ConjunctionStatement combines two RuleStatements
type ConjunctionStatement struct {
	Conjunction Conjunction `json:"conjunction"`
}

// Parse makes the ConjunctionStatement Expressionable
func (c *ConjunctionStatement) Parse(_ interface{}) (Expression, error) {
	switch c.Conjunction {
	case And:
		return CreateAndConjunctionExpression(&TrueExpression), nil
	case Or:
		return CreateOrConjunctionExpression(&FalseExpression), nil
	default:
		return CreateAndConjunctionExpression(&TrueExpression), nil
	}
}

// CreateConjunctionStatement Creates Conjunction Statement from string
func CreateConjunctionStatement(input string) *ConjunctionStatement {
	conjunction, err := ToConjunction(input)
	if err != nil {
		panic(err)
	}
	return &ConjunctionStatement{Conjunction: conjunction}
}

type CollectionStatement struct {
	Type string        `json:"type"`
	Rule RuleStatement `json:"rule"`
}

type StrSlice []string

func (s StrSlice) GetOrDefault(index int, defaultValue string) string {
	if index >= 0 && index < len(s) {
		return s[index]
	}
	return defaultValue
}

func (s StrSlice) GetOrEmpty(index int) string {
	return s.GetOrDefault(index, "")
}
