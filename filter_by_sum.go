package sprint

func FilterBySum(arr [][]int, limit int) [][]int {
	result := [][]int{}
	for _, subaarr := range arr {
		sum := 0
		for _, v := range subaarr {
			sum += v
		}
		if sum >= limit {
			result = append(result, subaarr)
		}
	}
	return result
}
