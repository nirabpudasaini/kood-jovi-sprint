package sprint

func Sqrt(n int) int {
	if n == 1 {
		return 1
	}
	for i := 0; i <= n/2; i++ {
		if i*i == n {
			return i
		} else if i*i > n {
			return 0
		}
	}
	return 0
}
