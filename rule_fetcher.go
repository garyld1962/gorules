package gorules

import (
	"fmt"
	"io/ioutil"
	"os"
)

// RuleFetcher generalizes the way to get rule from any source
type RuleFetcher interface {
	Fetch() string
}

// RuleText is the RuleFetcher used when rules are directly given in api
type RuleText string

// Fetch makes RuleText implement RuleFetcher
func (rt RuleText) Fetch() string {
	return string(rt)
}

// RuleFromFile is the RuleFetcher used when rules names are specified
type RuleFromFile struct {
	ruleName string
}

// Fetch makes RuleTRuleFromFileext implement RuleFetcher
func (rt RuleFromFile) Fetch() string {
	file, e := ioutil.ReadFile("../rulez.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	x := parseStringToJSONObject(string(file))

	return selectValue(x, rt.ruleName).(string)
}

// NewRuleFromFile uses the rule name give to create a RuleFromFile Fetcher
func NewRuleFromFile(rule string) RuleFetcher {
	return RuleFromFile{ruleName: rule}
}

// NewRuleFetcher creates a RuleFetcher depending on the input string given
func NewRuleFetcher(input string) RuleFetcher {
	if containsSpace(input) {
		return RuleText(input)
	}

	return NewRuleFromFile(input)
}
