package sprint

func IsPalindrome(s string) bool {
	return s == StrReverse(s)
}

func StrReverse(s string) string {
	sRuneArray := []rune(s)
	result := ""
	for i := len(sRuneArray) - 1; i >= 0; i-- {
		result += string(sRuneArray[i])
	}
	return result
}
