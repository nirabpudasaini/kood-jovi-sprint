package sprint

func GenerateRange(min, max int) []int {
	if min > max {
		return nil
	}
	results := make([]int, 0)
	for min < max {
		results = append(results, min)
		min++
	}
	return results
}
