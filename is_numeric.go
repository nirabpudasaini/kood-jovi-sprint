package sprint

func IsNumeric(s string) bool {
	isNumeric := true
	for _, v := range s {
		if rune(v) < '0' || rune(v) > '9' {
			isNumeric = false
			break
		}
	}
	return isNumeric
}
