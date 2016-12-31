package gorules

// StringSlice is a type to add additional functionality to string slice
type StringSlice []string

func (s StringSlice) getOrDefault(index int, defaultValue string) string {
	if index >= 0 && index < len(s) {
		return s[index]
	}
	return defaultValue
}

func (s StringSlice) getOrEmpty(index int) string {
	return s.getOrDefault(index, "")
}

// NewStringSlice creates a new stringSlice from string array
func NewStringSlice(stringArray []string) StringSlice {
	return StringSlice(stringArray)
}

// func (s StringSlice) SFilter(f func(string) bool) StringSlice {
// 	vsf := make([]string, 0)
// 	for _, v := range s {
// 		if f(v) {
// 			vsf = append(vsf, v)
// 		}
// 	}
// 	return vsf
// }

func (s StringSlice) Map(f func(string) string) StringSlice {
	vsm := make([]string, len(s))
	for i, v := range s {
		vsm[i] = f(v)
	}
	return StringSlice(vsm)
}
