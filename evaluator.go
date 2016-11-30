package gorules

import "fmt"

// Evaluator ...
type Evaluator interface {
	Evaluate(ex Expression) (bool, error)
}

// ValueExpressionEvaluator ...
type ValueExpressionEvaluator struct {
}

// Evaluate ...
func (ve ValueExpressionEvaluator) Evaluate(ex Expression) (bool, error) {

	e := ex.(ValueExpression)
	fmt.Println(e.Operator)
	switch e.Operator {
	case "IsEqualTo":
		return equals(e.Value, e.Target), nil

	case "IsNotEqualTo":
		return !equals(e.Value, e.Target), nil

	case "IsGreater":
		r, err := isGreater(e.Value, e.Target)
		if err != nil {
			return false, err
		}
		return r, nil
	case "IsGreaterThanOrEquals":

		r, err := isGreaterOrEqual(e.Target, e.Value)
		if err != nil {
			return false, err
		}
		return r, nil
	case "IsLessThanOrEquals":

		r, err := isLessOrEqual(e.Target, e.Value)
		if err != nil {
			return false, err
		}
		return r, nil

	case "IsLessThan":
		r, err := isGreater(e.Value, e.Target)
		if err != nil {
			return false, err
		}
		return !r, nil

	case "IsNull":

		return isNull(e.Value, e.Target)

	case "IsNotNull":
		return isNotNull(e.Value, e.Target)
	case "IsFalse":
		return isFalse(e.Value, e.Target)
	case "IsTrue":
		return isTrue(e.Value, e.Target)
	case "StartsWith":
		return startsWith(e.Value, e.Target)
	case "EndsWith":
		return endsWith(e.Value, e.Target)
	case "Contains":
		return contains(e.Value, e.Target)

	case "In":
		r := in(e.Value, e.Target)
		return r, nil
	case "NotIn":
		r := !in(e.Value, e.Target)
		return r, nil

	default:
	}
	panic("undefined Operator " + e.Operator)
}
