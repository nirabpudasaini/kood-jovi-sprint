package sprint

func IsLower(s string) bool {
	isLowerCase := true
	for _, v := range s {
		if rune(v) < 'a' || rune(v) > 'z' {
			isLowerCase = false
			break
		}
	}
	return isLowerCase
}
