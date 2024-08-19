package sprint

func NextPrime(n int) int {
	if isPrime(n) {
		return n
	}
	return NextPrime(n + 1)
}

func isPrime(n int) bool {
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
