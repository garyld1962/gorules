package gorules

type stringSlice []string

func (s stringSlice) getOrDefault(index int, defaultValue string) string {
	if index >= 0 && index < len(s) {
		return s[index]
	}
	return defaultValue
}

func (s stringSlice) getOrEmpty(index int) string {
	return s.getOrDefault(index, "")
}
