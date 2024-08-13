package sprint

func ToCapitalCase(s string) string {
	sRuneArray := []rune(s)
	isFirstLetter := true
	for i, v := range sRuneArray {
		if isFirstLetter {
			if v >= 'A' && v <= 'Z' {
				isFirstLetter = false
			} else if v >= 'a' && v <= 'z' {
				sRuneArray[i] -= 32
				isFirstLetter = false
			} else if v >= '0' && v <= '9' {
				isFirstLetter = false
			}
		} else {
			if v >= 'a' && v <= 'z' {
				continue
			} else if v >= 'A' && v <= 'Z' {
				sRuneArray[i] += 32
			} else if v >= '0' && v <= '9' {
				continue
			} else {
				isFirstLetter = true
			}
		}

	}
	return string(sRuneArray)
}
