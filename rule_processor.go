package gorules

import (
	"fmt"
	"io/ioutil"
	"os"
)

// RuleProcessor generalizes the way to get rule from any source
type RuleProcessor interface {
	Process(map[string]interface{}) interface{}
}

// RuleText is the RuleFetcher used when rules are directly given in api
type RuleText string

// Fetch gets string out of RuleText
func (rt RuleText) Fetch() string {
	return string(rt)
}

// Process makes RuleText implement RuleProcessor
func (rt RuleText) Process(data map[string]interface{}) interface{} {
	rule := rt.Fetch()
	return EvaluateRules(rule, data)
}

// RuleFromFile is the RuleFetcher used when rules names are specified
type RuleFromFile struct {
	ruleName string
}

// Fetch makes RuleFromFile implement RuleFetcher and gets data from rulez file
func (rt RuleFromFile) Fetch() string {
	file, e := ioutil.ReadFile("./rulez.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	x := parseStringToJSONObject(string(file))

	return selectValue(x, rt.ruleName).(string)
}

// Process makes RuleFromFile implement RuleProcessor
func (rt RuleFromFile) Process(data map[string]interface{}) interface{} {
	rule := rt.Fetch()
	return EvaluateRules(rule, data)
}

// NewRuleFromFile uses the rule name give to create a RuleFromFile Fetcher
func NewRuleFromFile(rule string) RuleProcessor {
	return RuleFromFile{ruleName: rule}
}

// RuleCollection can be a collection of RuleFromFile/RuleText
type RuleCollection []string

// Fetch makes RuleFromFile implement RuleFetcher and gets data from rulez file
func (rt RuleCollection) Fetch() []RuleProcessor {
	ruleCollection := make([]RuleProcessor, 0)
	for _, toProcess := range rt {
		fmt.Println(toProcess, len(rt))
		ruleCollection = append(ruleCollection, NewRuleFetcher(toProcess))
	}
	fmt.Println(ruleCollection, len(rt))

	return ruleCollection
}

// Process makes RuleCollection implement RuleProcessor
func (rt RuleCollection) Process(data map[string]interface{}) interface{} {
	rules := rt.Fetch()
	var resultCollection []bool
	for _, toProcess := range rules {
		resultCollection = append(resultCollection, toProcess.Process(data).(bool))
	}

	return resultCollection
}

//NewRuleCollection creates a new RuleCollection
func NewRuleCollection(rule []string) RuleProcessor {
	return RuleCollection(rule)
}

type RuleObjectCollection map[string]string

// Fetch makes RuleObjectCollection implement RuleFetcher and gets data from rulez file
func (rt RuleObjectCollection) Fetch() map[string]RuleProcessor {
	var ruleCollection = make(map[string]RuleProcessor)

	for k, v := range rt {
		ruleCollection[k] = NewRuleFetcher(v)
	}

	fmt.Println(ruleCollection, len(rt))

	return ruleCollection
}

// Process makes RuleCollection implement RuleProcessor
func (rt RuleObjectCollection) Process(data map[string]interface{}) interface{} {
	fmt.Println(rt)
	rules := rt.Fetch()
	var resultCollection = make(map[string]bool)
	for k, v := range rules {
		fmt.Println(v)
		resultCollection[k] = v.Process(data).(bool)
	}

	return resultCollection
}

//NewRuleCollection creates a new RuleCollection
func NewRuleObjectCollection(rule map[string]string) RuleProcessor {
	return RuleObjectCollection(rule)
}

// NewRuleFetcher creates a RuleFetcher depending on the input string given
func NewRuleFetcher(input interface{}) RuleProcessor {
	_, ok := input.(map[string]string)
	if isArray(input) {
		fmt.Println(input, isArray(input), len(input.([]string)))
		return NewRuleCollection(input.([]string))
	} else if ok {
		return RuleObjectCollection(input.(map[string]string))
	} else if containsSpace(input.(string)) {
		return RuleText(input.(string))
	}

	return NewRuleFromFile(input.(string))
}
