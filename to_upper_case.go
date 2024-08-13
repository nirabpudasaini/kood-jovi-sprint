package sprint

func ToUpperCase(s string) string {
	sRuneArray := []rune(s)
	for i, v := range sRuneArray {
		if v >= 'a' && v <= 'z' {
			sRuneArray[i] -= 32
		}
	}
	return string(sRuneArray)
}
