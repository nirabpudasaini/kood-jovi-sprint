package sprint

func ToThePowerIterative(n int, power int) int {
	result := 1
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	for power >= 1 {
		result = int(uint(result * n))
		power -= 1
	}
	return result
}
