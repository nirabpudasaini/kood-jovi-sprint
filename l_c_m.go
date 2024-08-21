package sprint

func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}

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
