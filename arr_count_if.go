package sprint

func ArrCountIf(f func(string) bool, a []string) int {
	count := 0
	for _, v := range a {
		if f(v) {
			count++
		}
	}
	return count
}
