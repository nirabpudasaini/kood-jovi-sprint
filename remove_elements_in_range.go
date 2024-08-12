package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	results := []float64{}
	if from > to {
		from, to = to, from
	}
	for i, v := range arr {
		if i < from || i >= to {
			results = append(results, v)
		}
	}
	return results
}
