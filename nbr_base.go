package sprint

func NbrBase(n int, base string) string {
	result := ""
	baseNum := len([]rune(base))
	isNegative := false
	if n < 0 {
		isNegative = true
		n *= -1
	}
	if ValidBase(base) {
		for n >= baseNum {
			result = string(base[n%baseNum]) + result
			n /= baseNum
		}
		result = string(base[n]) + result
	} else {
		result = "NV"
	}
	if isNegative {
		result = "-" + result
	}
	return result
}

func ValidBase(base string) bool {
	if len([]rune(base)) < 2 {
		return false
	}
	countMap := make(map[string]int)
	for _, v := range []byte(base) {
		countMap[string(v)] = countMap[string(v)] + 1
		if countMap[string(v)] > 1 {
			return false
		}
	}
	return true
}
