package sprint

func ReverseAlphabetValue(ch rune) rune {
	if ch >= 97 && ch <= 109 {
		return 110 + (109 - ch)
	} else if ch > 109 && ch <= 122 {
		return 97 + (122 - ch)
	}
	return 0
}
