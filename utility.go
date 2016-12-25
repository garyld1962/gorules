package gorules

import (
	"bytes"
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/stretchr/stew/objects"
)

func in(v, t string) bool {
	return strings.Contains(v, t)
}

func equals(value, target string) (bool, error) {
	return value == target, nil
}

func isMatch(value, target string) (bool, error) {

	reg, err := regexp.Compile(value)
	if err != nil {
		return false, err
	}
	return reg.MatchString(target), nil
}

func isGreater(v, t string) (bool, error) {

	var val1, val2 float64
	var err error
	if val1, err = getNumber(v); err != nil {
		return false, err
	}
	if val2, err = getNumber(t); err != nil {
		return false, err
	}

	return val1 > val2, nil

}

func isGreaterOrEqual(v, t string) (bool, error) {

	var val1, val2 float64
	var err error
	if val1, err = getNumber(v); err != nil {
		return false, err
	}
	if val2, err = getNumber(t); err != nil {
		return false, err
	}

	return val1 > val2, nil

}

func isLessOrEqual(v, t string) (bool, error) {

	var val1, val2 float64
	var err error
	if val1, err = getNumber(v); err != nil {
		return false, err
	}
	if val2, err = getNumber(t); err != nil {
		return false, err
	}

	return val1 < val2, nil

}

func getNumber(v string) (float64, error) {

	n, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, err
	}
	return n, err
}

func isNull(value, target string) (bool, error) {

	return target == "", nil
}

func isNotNull(value, target string) (bool, error) {

	return target != "", nil
}

func isTrue(value, target string) (bool, error) {

	r, err := strconv.ParseBool(target)
	if err != nil {
		return false, err
	}
	return r, nil
}

func isFalse(value, target string) (bool, error) {

	r, err := strconv.ParseBool(target)
	if err != nil {
		return false, err
	}
	return r, nil
}

func contains(value, target string) (bool, error) {

	return strings.Contains(value, target), nil
}

func startsWith(value, target string) (bool, error) {
	return strings.HasPrefix(value, target), nil
}

func endsWith(value, target string) (bool, error) {

	return strings.HasSuffix(value, target), nil
}

func getInterfaceType(t interface{}) reflect.Kind {
	switch t.(type) {
	case bool:
		return reflect.Bool
	case uint8, uint16, uint32, uint64, uint:
		return reflect.Uint64
	case int8, int16, int32, int64, int:
		return reflect.Int64
	case float32, float64:
		return reflect.Float64
	case string:
		return reflect.String
	default:
		return reflect.TypeOf(t).Kind()
	}
}

func isArray(i interface{}) bool {
	return reflect.TypeOf(i).Kind() == reflect.Array || reflect.TypeOf(i).Kind() == reflect.Slice
}
func isMap(i interface{}) bool {
	return reflect.TypeOf(i).Kind() == reflect.Map
}
func getArrayType(arr interface{}) (reflect.Type, error) {

	if (isArray(arr)) == false {
		return nil, errors.New("value is not an arry or slice")
	}
	return reflect.TypeOf(arr).Elem(), nil
}
func isInArray(arr, val interface{}) (bool, error) {
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice, reflect.Array:
		// get the value of "target". It should not be a collection type

		switch val.(type) {
		case uint8, uint16, uint32, uint64, uint:
			tv := reflect.ValueOf(val).Uint()
			ar := reflect.ValueOf(arr)
			for i := 0; i < ar.Len(); i++ {
				el := ar.Index(i).Uint()
				if el == tv {
					return true, nil
				}
			}
			return false, nil

		case int8, int16, int32, int64, int:
			tv := reflect.ValueOf(val).Int()
			ar := reflect.ValueOf(arr)
			for i := 0; i < ar.Len(); i++ {
				el := ar.Index(i).Int()
				if el == tv {
					return true, nil
				}
			}
			return false, nil

		case float32, float64:
			tv := reflect.ValueOf(val).Float()
			ar := reflect.ValueOf(arr)
			for i := 0; i < ar.Len(); i++ {
				el := ar.Index(i).Float()
				if el == tv {
					return true, nil
				}
			}
			return false, nil

		case string:
			tv := reflect.ValueOf(val).String()
			ar := reflect.ValueOf(arr)
			for i := 0; i < ar.Len(); i++ {
				el := ar.Index(i).String()
				if el == tv {
					return true, nil
				}
			}
			return false, nil

		default:

		}
	}
	return true, nil
}
func selectValue(m map[string]interface{}, path string) interface{} {

	propertyNames := strings.Split(path, ".")
	// we only have one path element so return what we have
	if len(propertyNames) == 1 {
		return m[propertyNames[0]]
	}
	// create map using propertyName as a key and the []interface{} as a value
	m1 := make(map[string]interface{})
	m1[propertyNames[0]] = m[propertyNames[0]]
	// create a new path with the remainder of the parts.
	newPath := strings.Join(propertyNames[2:], ",")
	return selectValue(m1, newPath)
}

func getValue(m map[string]interface{}, path string) (bool, interface{}) {
	propertyNames := strings.Split(path, ".")
	if len(propertyNames) == 1 && m[propertyNames[0]] != nil {
		return true, m[propertyNames[0]]
	}
	return false, nil
}

func reverse(value []string) []string {
	result := []string{}
	for i := len(value) - 1; i >= 0; i-- {
		result = append(result, value[i])
	}
	return result
}

func isEven(value int) bool {
	return value%2 == 0
}

func concatStrings(strgs ...string) string {
	var buffer bytes.Buffer
	for _, strg := range strgs {
		buffer.WriteString(strg)
	}
	return buffer.String()
}

func trim(input string) string {
	return strings.TrimSpace(input)
}

func splitString(delimiter string) func(string) []string {

	return func(input string) []string {
		return strings.Split(input, delimiter)
	}
}

var spiltWithSpace = splitString(" ")

var spiltWithDot = splitString(".")

func getArrayPathAndKey(path string) (string, string) {
	s := spiltWithDot(path)
	lengt := len(s)
	final := s[lengt-1]
	return strings.Join(s[0:lengt-1], "."), final
}

func parseStringToJSONObject(jsonAsString string) map[string]interface{} {
	result, err := objects.NewMapFromJSON(jsonAsString)
	if err != nil {
		panic(err)
	}
	return result
}

func startsWithIdentifier(toStartValue string) func(stringToCheck string) bool {
	return func(stringToCheck string) bool {
		return strings.HasPrefix(stringToCheck, toStartValue)
	}
}

func hasStringBetween(delimiter string, input string) bool {
	return strings.Count(input, delimiter) >= 2
}

var startsWithSingleQuotes = startsWithIdentifier("'")

func stringBetween(delimiter string) func(input string) string {
	return func(input string) string {
		if !hasStringBetween(delimiter, input) {
			return ""
		}
		return strings.Split(input, delimiter)[1]
	}
}

var stringBetweenSingleQuotes = stringBetween("'")

func surroundBy(delimiter string) func(input string) string {
	return func(input string) string {
		return concatStrings(delimiter + input + delimiter)
	}
}

var surroundBySingleQuotes = surroundBy("'")

func encodeSpace(input string) string {
	return strings.Replace(input, " ", "!+!", -1)
}

func decodeSpace(input string) string {
	return strings.Replace(input, "!+!", " ", -1)
}

func encodeString(input string) string {
	return recursiveEncode("", input)
}

func recursiveEncode(accum, input string) string {
	v := strings.Count(input, "'")
	if v == 0 {
		return accum + input
	}

	arr := strings.SplitN(input, "'", 2)
	current := arr[0]
	remaining := arr[1]
	if v%2 != 0 {
		current = surroundBySingleQuotes(encodeSpace(current))
	}
	return recursiveEncode(accum+current, remaining)
}

// func Totest(input string) string {
// 	v := strings.Count(input, "'")
// 	var result = ""
// 	var temp = ""
// 	var arr []string
// 	for i := 1; i <= v; i++ {
// 		arr = strings.SplitN(input, "'", 2)
// 		temp = arr[0]
// 		input = arr[1]
// 		if i%2 == 0 {
// 			temp = surroundBySingleQuotes(encodeSpace(temp))
// 		}

// 		result = concatStrings(result, temp)
// 	}
// 	return result
// }
