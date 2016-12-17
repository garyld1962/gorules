package gorules

import (
	"strings"

	"fmt"

	objects "github.com/stretchr/stew/objects"
)

type Expressionable interface {
	Parse(interface{}) (Expression, error)
}

type RuleStatement struct {
	Branch   string `json:"branch"`
	Selector string `json:"selector"`
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Target   string `json:"target"`
}

type ConjunctionStatement struct {
	Conjunction Conjunction `json:"type"`
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

func (s *RuleStatement) Parse(data interface{}) (Expression, error) {
	dat := data.(objects.Map)
	test := CreateValueExpressionWithTarget(s.Operator, "", GetKeyFromJSON(dat, s.Key), s.Target)
	return test, nil
}

func (c *ConjunctionStatement) Parse(_ interface{}) (Expression, error) {
	switch c.Conjunction {
	case And:
		return CreateAndConjunctionExpression(TrueExpression), nil
	case Or:
		return CreateAndConjunctionExpression(FalseExpression), nil
	default:
		return CreateAndConjunctionExpression(TrueExpression), nil
	}
}

func CreateRuleStatement(input string) *RuleStatement {
	parsed := StrSlice(reverse(strings.Split(input, " ")))
	rule := &RuleStatement{Target: parsed.GetOrEmpty(0),
		Operator: parsed.GetOrEmpty(1),
		Key:      parsed.GetOrDefault(2, "data"),
		Selector: parsed.GetOrDefault(3, "THIS"),
		Branch:   parsed.GetOrDefault(4, "IF")}
	return rule
}

// CreateConjunctionStatement Creates Conjunction Statement from string
func CreateConjunctionStatement(input string) *ConjunctionStatement {
	conjunction, err := ToConjunction(input)
	fmt.Println(conjunction)
	if err != nil {
		panic(err)
	}
	return &ConjunctionStatement{Conjunction: conjunction}
}
