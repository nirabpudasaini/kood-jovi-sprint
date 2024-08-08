package sprint

func AlphaNumber(n int) string {
	result := ""
	negative := false
	if n < 0 {
		negative = true
		n = n * -1
	}
	for n > 1 {
		result = string(rune(n%10+95)) + result
		n /= 10
	}
	if negative {
		result = "-" + result
	}
	return result
}
