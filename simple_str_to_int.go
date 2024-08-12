package sprint

func SimpleStrToInt(s string) int {
	var result int
	for _, i := range []byte(s) {
		if i < '0' && i > '9' {
			return 0
		}
		result = result*10 + int(i-'0')
	}
	return result
}
