package sprint

func BalancedParentheses(input string) bool {
	parStack := []string{}
	for _, v := range input {
		lastIndex := len(parStack) - 1
		switch string(v) {
		case "(":
			parStack = append(parStack, "(")
		case "{":
			parStack = append(parStack, "{")
		case "[":
			parStack = append(parStack, "[")
		case ")":
			if lastIndex < 0 {
				return false
			}
			if parStack[lastIndex] == "(" {
				parStack = parStack[:lastIndex]
			} else {
				return false
			}
		case "}":
			if lastIndex < 0 {
				return false
			}
			if parStack[lastIndex] == "{" {
				parStack = parStack[:lastIndex]
			} else {
				return false
			}
		case "]":
			if lastIndex < 0 {
				return false
			}
			if parStack[lastIndex] == "[" {
				parStack = parStack[:lastIndex]
			} else {
				return false
			}
		}
	}

	return len(parStack) == 0

}
