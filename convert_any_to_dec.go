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
