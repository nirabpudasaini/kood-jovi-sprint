package sprint

import "fmt"

func Countdown(n int) string {
	result := ""
	for n > 0 {
		result += fmt.Sprintf("%d,", n)
		n -= 2

	}
	result += "0!"
	return result
}
