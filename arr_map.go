package sprint

func ArrMap(f func(int) bool, a []int) []bool {
	result := []bool{}
	for _, v := range a {
		result = append(result, f(v))
	}
	return result
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
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

func IsNegative(n int) bool {
	return n < 0
}
