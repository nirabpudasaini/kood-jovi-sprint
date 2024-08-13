package sprint

func ToCapitalCase(s string) string {
	sRuneArray := []rune(s)
	for i, v := range sRuneArray {
		if v == ' ' || v == '-' || i == 0 {
			if i != len(sRuneArray)-1 {
				if sRuneArray[i+1] >= 'a' && sRuneArray[i+1] <= 'z' {
					sRuneArray[i+1] = sRuneArray[i+1] - 32
				}
			}
		}
	}
	return string(sRuneArray)
}
