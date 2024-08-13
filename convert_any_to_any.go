package sprint

import (
	"math"
	"strings"
)

func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {
	decimal := convertToDec(nbr, baseFrom)
	return converFromDec(decimal, baseTo)
}

func anyBaseValidation(base string) bool {
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

func convertToDec(s, base string) int {
	result := 0
	baseNum := len([]rune(base))
	sLen := len([]rune(s))
	if anyBaseValidation(base) {
		for i, v := range []rune(s) {
			result += strings.Index(base, string(v)) * int(math.Pow(float64(baseNum), float64(sLen-1-i)))
		}
	} else {
		return 0
	}
	return result
}

func converFromDec(n int, base string) string {
	result := ""
	baseNum := len([]rune(base))
	isNegative := false
	if n < 0 {
		isNegative = true
		n *= -1
	}
	if anyBaseValidation(base) {
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
