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
			if parStack[lastIndex] == "(" {
				parStack = parStack[:lastIndex]
			}
		case "}":
			if parStack[lastIndex] == "{" {
				parStack = parStack[:lastIndex]
			}
		case "]":
			if parStack[lastIndex] == "[" {
				parStack = parStack[:lastIndex]
			}
		}
	}

	return len(parStack) == 0

}
