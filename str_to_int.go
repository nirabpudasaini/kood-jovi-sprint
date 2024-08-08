package sprint

func StrToInt(s string) int {
	var result int
	var negative = false
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		s = s[1:]
		negative = true
	}
	for _, i := range []byte(s) {
		if i >= '0' && i <= '9' {
		} else {
			return 0
		}
		result = result*10 + int(i-'0')
	}
	if negative {
		result = result * -1
	}
	return result
}
