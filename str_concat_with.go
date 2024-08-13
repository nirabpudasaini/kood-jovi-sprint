package sprint

func StrConcatWith(strs []string, sep string) string {
	result := ""
	lastIndex := len(strs) - 1
	for i, v := range strs {
		if i == lastIndex {
			result += v
			break
		}
		result = result + v + sep
	}
	return result
}
