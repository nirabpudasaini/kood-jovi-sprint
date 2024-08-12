package sprint

import "fmt"

func CombN(n int) []string {
	result := []string{}
	maxValue := 1
	minValue := 0
	for i := 1; i <= n; i++ {
		maxValue = maxValue * 10
		if i >= 1 {
			minValue = minValue*10 + i
		}
	}
	maxValue -= 1
	minValue = minValue / 10
	isAscendingDigits := true
	for i := minValue; i <= maxValue; i++ {
		if i > 123456789 {
			break
		}
		currentNumber := fmt.Sprintf("%0*d", n, i)
		maxdigit := '0' - 1
		for _, v := range []byte(currentNumber) {
			if rune(v) > maxdigit {
				maxdigit = rune(v)
				isAscendingDigits = true
			} else {
				isAscendingDigits = false
				break
			}
		}
		if isAscendingDigits {
			result = append(result, fmt.Sprintf("%0*d", n, i))
		}
	}

	return result
}
