package sprint

func Countdown(n int) string {
	result := ""
	for n > 0 {
		result += string(rune(n) + rune('0'))
		n -= 2

	}
	result += "0!"
	return result
}
