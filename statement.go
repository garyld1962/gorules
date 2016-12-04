package gorules

import (
	objects "github.com/stretchr/stew/objects"
	"strings"
)

type Expressionable interface {
	Parse(interface{}) (Expression, error)
}

type ruleStatement struct {
	Branch   string `json:"branch"`
	Selector string `json:"selector"`
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Target   string `json:"target"`
}

type conjunctionStatement struct {
	Type Conjunction `json:"type"`
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


func (s *ruleStatement) Parse(data interface{}) (Expression, error) {
	dat := data.(objects.Map)
	test := CreateValueExpressionWithTarget(s.Operator, "", GetKeyFromJSON(dat, s.Key), s.Target)
	return test, nil
}

func (c *conjunctionStatement) Parse(_ interface{}) (Expression, error) {
	return CreateAndExpression(True{}), nil
}

func CreateRuleStatement(input string) *ruleStatement {
	parsed := StrSlice(reverse(strings.Split(input, " ")))
	rule := &ruleStatement{Target: parsed.GetOrEmpty(0),
		Operator: parsed.GetOrEmpty(1),
		Key:      parsed.GetOrDefault(2, "data"),
		Selector: parsed.GetOrDefault(3, "THIS"),
		Branch:   parsed.GetOrDefault(4, "IF")}
	return rule
}


func CreateConjunctionStatement(input string) *conjunctionStatement {
	if input == "AND" {
		return &conjunctionStatement{Type: 1}
	}
	return &conjunctionStatement{Type: 2}
}

