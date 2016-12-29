package gorules

import "fmt"

// Selector defines whether the JSON key is a value or an array
type Selector int

const (
	// This refers to the value of the key placed after this
	This Selector = iota
	// Any lets the key given treated as a array and figures out how to operate
	Any
	// All lets the key given treated as a array and figures out how to operate
	All
	maxSelectorFlag
)

var selectorNames = [...]string{
	This: "THIS",
	Any:  "ANY",
	All:  "ALL",
}

// toSelector converts string to Selector
func toSelector(s string) (Selector, error) {
	for i, r := range selectorNames {
		if s == r {
			return Selector(i), nil
		}
	}
	return maxSelectorFlag, fmt.Errorf("Invalid Selector value %q", s)
}

// String makes Operator implement Stringer
func (v Selector) String() string {
	if v >= maxSelectorFlag {
		return "Invalid Selector"
	}
	return selectorNames[v]
}

// isSelector checks if the string is a valid Selector
func isSelector(value string) bool {
	_, err := toSelector(value)
	if err == nil {
		return true
	}
	return false
}

type selectorFunc func(*RuleStatement, map[string]interface{}) Expression

var selectorFuncList = map[Selector]selectorFunc{
	This: createValueExpressionFromRuleStmt,
	Any:  createConjuntionExprFromCollectionStmt,
	All:  createConjuntionExprFromCollectionStmt,
}

var selectorConjExprList = map[Selector]ConjunctionExpression{
	This: createOrConjunctionExpression(FalseExpression).(ConjunctionExpression),
	Any:  createOrConjunctionExpression(FalseExpression).(ConjunctionExpression),
	All:  createAndConjunctionExpression(TrueExpression).(ConjunctionExpression),
}

func selectorFunctions(selector Selector) selectorFunc {
	return selectorFuncList[selector]
}

func selectorConjExprMap(selector Selector) ConjunctionExpression {
	return selectorConjExprList[selector]
}
