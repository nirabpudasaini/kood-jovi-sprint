package sprint

func FactorialIterative(n int) int {
	if n <= 0 {
		return 0
	}
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial = int(uint(i * factorial))
	}
	return factorial
}
