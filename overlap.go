package sprint

func Overlap(arr1, arr2 []int) []int {
	arrMap := map[int]int{}
	result := []int{}
	for _, v := range arr1 {
		arrMap[v] += 1
	}
	for _, v := range arr2 {
		if arrMap[v] > 0 {
			result = append(result, v)
			arrMap[v] -= 1
		}
	}
	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}
