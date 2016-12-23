package gorules

import (
	"fmt"
	"strings"

	objects "github.com/stretchr/stew/objects"
)

const dslSeperator = "|"

// EvaluateDSL evaluates DSL to a bool
func EvaluateDSL(dslText string, data string) bool {
	ruleToEvaluate := ParseString(dslText, data)
	result := EvaluateExpressions(CreateOrConjunctionExpression(&FalseExpression), ruleToEvaluate.expressions)
	return result
}

// ParseDSL parses simple DSL to Rule (array of Expressions)
// Will be deleted once all opeartors are tested with
func ParseDSL(dslText string, data string) *Rule {
	lst := strings.Split(dslText, dslSeperator)
	m, _ := objects.NewMapFromJSON(data)
	var rle = &Rule{}
	var exp Expressionable
	for _, x := range lst {
		if IsConjunction(x) {
			exp = CreateConjunctionStatement(x)
		} else {
			exp = CreateRuleStatement(strings.TrimSpace(x))
		}
		parsed, _ := exp.Parse(m)
		rle.Add(&parsed)
	}
	boo := EvaluateExpressions(CreateOrConjunctionExpression(&FalseExpression), rle.expressions)
	fmt.Println("Output", boo)
	return rle
}

//ParseString parses the DSL with space and creates Rule
func ParseString(dslText string, data string) *Rule {

	var ruleToEvaluate = &Rule{}
	var latestRuleStatement Expressionable
	var toParse string

	words := spiltWithSpace(dslText)
	jsonObj, _ := objects.NewMapFromJSON(data)

	for _, word := range words {
		if !IsConjunction(word) {
			toParse = concatStrings(toParse, word, " ")
		} else {
			latestRuleStatement = CreateRuleStatementFromExisting(latestRuleStatement, trim(toParse))
			toParse = ""
			valueExpr, _ := latestRuleStatement.Parse(jsonObj)
			conjExpr, _ := CreateConjunctionStatement(word).Parse(jsonObj)
			ruleToEvaluate.Add(&valueExpr)
			ruleToEvaluate.Add(&conjExpr)
		}

	}

	latestRuleStatement = CreateRuleStatementFromExisting(latestRuleStatement, trim(toParse))
	lastRuleExpr, _ := latestRuleStatement.Parse(jsonObj)
	ruleToEvaluate.Add(&lastRuleExpr)

	result, _ := ruleToEvaluate.Evalute()
	fmt.Println("Out", result)
	return ruleToEvaluate
}

// EvaluateExpressions loops through all Expressions and evaluates it
func EvaluateExpressions(accum Expression, expressions []*Expression) bool {
	if len(expressions) == 0 {
		value, _ := accum.Evaluate()
		return value
	}

	expr := *expressions[0]
	if isConjunctionExpression(expr) {
		conj, _ := expr.(*ConjunctionExpression)
		isTrue, _ := accum.Evaluate()
		boolExpr := CreateBoolExpression(isTrue)
		conj.Add(&boolExpr)
		accum = conj
	} else {
		conj, _ := accum.(*ConjunctionExpression)
		isTrue, _ := expr.Evaluate()
		boolExpr := CreateBoolExpression(isTrue)
		conj.Add(&boolExpr)
		accum = conj
	}

	return EvaluateExpressions(accum, expressions[1:])
}

//GetKeyFromJSON get the value for key in JSON
// func GetKeyFromJSON(obj objects.Map, key string) string {
// 	value := obj.Get(key)
// 	//fmt.Println(key,value)
// 	switch x := value.(type) {
// 	case map[string]interface{}:
// 		return "JSONObject"
// 	case []interface{}:
// 		return "JSONArray"
// 	case string:
// 		return x
// 	case int:
// 		return fmt.Sprintf("%d", x)
// 	default:
// 		return "default"
// 	}
// }

// type Accum struct {
// 	Rules    *Rule          `json:"rules"`
// 	LastRule Expressionable `json:"lastRule"`
// 	ToParse  string         `json:"ToParse"`
// 	Data     objects.Map
// }
// func ParseStringRecursion(dslText string, data string) *Rule {
// 	// " IF THIS country IsEqualTo USA |AND| IF THIS country IsEqualTo USA "
// 	m, _ := objects.NewMapFromJSON(data)
// 	var initialAccum = Accum{Rules: &Rule{expressions: make([]*Expression, 0)}, LastRule: nil, ToParse: "", Data: m}
// 	lst := strings.Split(dslText, " ")
// 	final := StringToStatement(initialAccum, lst)
// 	return final.Rules
// }
// func StringToStatement(accum Accum, textArray []string) Accum {
// 	var latestvalueStatement, latestConjunctionStatement Expressionable

// 	if len(textArray) == 0 {
// 		latestvalueStatement = CreateRuleStatementFromExisting(strings.TrimSpace(accum.ToParse), accum.LastRule)
// 		valparsed, _ := latestvalueStatement.Parse(accum.Data)
// 		accum.Rules.Add(&valparsed)
// 		return accum
// 	}
// 	valueToCheck := textArray[0]
// 	if !IsConjunction(valueToCheck) {
// 		fmt.Println(accum)
// 		return StringToStatement(Accum{Rules: accum.Rules, LastRule: accum.LastRule, ToParse: accum.ToParse + " " + valueToCheck}, textArray[1:])
// 	}

// 	latestvalueStatement = CreateRuleStatementFromExisting(strings.TrimSpace(accum.ToParse), accum.LastRule)
// 	latestConjunctionStatement = CreateConjunctionStatement(valueToCheck)
// 	valparsed, _ := latestvalueStatement.Parse(accum.Data)
// 	conjparsed, _ := latestConjunctionStatement.Parse(accum.Data)
// 	accum.Rules.Add(&valparsed)
// 	accum.Rules.Add(&conjparsed)
// 	fmt.Println(accum)
// 	return StringToStatement(Accum{Rules: accum.Rules, LastRule: accum.LastRule, ToParse: " ", Data: accum.Data}, textArray[1:])
// }
