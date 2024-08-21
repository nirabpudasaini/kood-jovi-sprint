package sprint

func GCD(a, b int) int {
	high := b
	low := a
	if a > b {
		high, low = low, high
	}
	for low != 0 {
		remainder := high % low
		high = low
		low = remainder
	}
	return high
}
