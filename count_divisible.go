package sprint

func CountDivisible(from, to, step, divisor int) int {
	count := 0
	if step <= 0 || divisor == 0 {
		return count
	}
	for i := from; i < to; i += step {
		if i%divisor == 0 {
			count += 1
		}
	}
	return count
}
