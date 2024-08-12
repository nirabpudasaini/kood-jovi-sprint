package sprint

func GenerateRange(min, max int) []int {
	results := make([]int, 0)
	for min < max {
		results = append(results, min)
		min++
	}
	return results
}
