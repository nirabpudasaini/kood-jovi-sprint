package sprint

func StrReverse(s string) string {
	sRuneArray := []rune(s)
	result := ""
	for i := len(sRuneArray) - 1; i >= 0; i-- {
		result += string(sRuneArray[i])
	}
	return result
}
