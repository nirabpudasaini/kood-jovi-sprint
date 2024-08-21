package sprint

func RemoveDuplicates(arr []int) []int {
	arrMap := make(map[int]int)
	result := []int{}
	for _, v := range arr {
		if arrMap[v] == 0 {
			result = append(result, v)
		}
		arrMap[v]++
	}
	return result
}
