package gorules

import (
	"fmt"
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

// RuleFilePath is the path to the rules JSON file. Override via GORULES_FILE_PATH env var.
var RuleFilePath = func() string {
	if p := os.Getenv("GORULES_FILE_PATH"); p != "" {
		return p
	}
	return "./rulez.json"
}()

// Fetch makes RuleFromFile implement RuleFetcher and gets data from rulez file
func (rt RuleFromFile) Fetch() string {
	file, err := os.ReadFile(RuleFilePath)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		return ""
	}
	x, err := parseStringToJSONObject(string(file))
	if err != nil {
		fmt.Printf("JSON parse error: %v\n", err)
		return ""
	}
	raw := selectValue(x, rt.ruleName)
	if raw == nil {
		return ""
	}
	if str, ok := raw.(string); ok {
		return str
	}
	return fmt.Sprintf("%v", raw)
}

// Process makes RuleFromFile implement RuleProcessor
func (rt RuleFromFile) Process(data map[string]interface{}) interface{} {
	rule := rt.Fetch()
	if rule == "" {
		return false
	}
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
		ruleCollection = append(ruleCollection, NewRuleProcessor(toProcess))
	}

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
		ruleCollection[k] = NewRuleProcessor(v)
	}

	return ruleCollection
}

// Process makes RuleCollection implement RuleProcessor
func (rt RuleObjectCollection) Process(data map[string]interface{}) interface{} {
	rules := rt.Fetch()
	var resultCollection = make(map[string]bool)
	for k, v := range rules {
		resultCollection[k] = v.Process(data).(bool)
	}

	return resultCollection
}

//NewRuleCollection creates a new RuleCollection
func NewRuleObjectCollection(rule map[string]string) RuleProcessor {
	return RuleObjectCollection(rule)
}

// NewRuleProcessor creates a RuleFetcher depending on the input string given
func NewRuleProcessor(input interface{}) RuleProcessor {
	// Check for map[string]string first
	if mapVal, ok := input.(map[string]string); ok {
		return RuleObjectCollection(mapVal)
	}
	// Check for []string (direct Go callers)
	if strSlice, ok := input.([]string); ok {
		return NewRuleCollection(strSlice)
	}
	// Check for string
	if str, ok := input.(string); ok {
		if containsSpace(str) {
			return RuleText(str)
		}
		return NewRuleFromFile(str)
	}
	// Fallback: return a RuleText that will evaluate to false
	return RuleText("")
}
