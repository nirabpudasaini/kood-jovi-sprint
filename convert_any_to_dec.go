package sprint

import (
	"math"
	"strings"
)

func ConvertAnyToDec(s, base string) int {
	result := 0
	baseNum := len([]rune(base))
	sLen := len([]rune(s))
	if BaseValidation(base) {
		for i, v := range []rune(s) {
			result += strings.Index(base, string(v)) * int(math.Pow(float64(baseNum), float64(sLen-1-i)))
		}
	} else {
		return 0
	}
	return result
}

func BaseValidation(base string) bool {
	if len([]rune(base)) < 2 {
		return false
	}
	countMap := make(map[string]int)
	for _, v := range []byte(base) {
		countMap[string(v)] = countMap[string(v)] + 1
		if countMap[string(v)] > 1 {
			return false
		}
		if v == '-' || v == '+' {
			return false
		}
	}
	return true
}
