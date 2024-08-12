package sprint

func GenerateRange(min, max int) []int {
	results := []int{}
	for min < max {
		results = append(results, min)
		min++
	}
	return results
}
