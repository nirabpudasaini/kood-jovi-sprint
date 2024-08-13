package sprint

func SplitWhitespaces(s string) []string {
	result := []string{}
	word := ""
	for _, v := range []byte(s) {
		if v == ' ' {
			result = append(result, word)
			word = ""
		} else {
			word += string(v)
		}
	}
	result = append(result, word)

	return result
}
