package sprint

func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}
