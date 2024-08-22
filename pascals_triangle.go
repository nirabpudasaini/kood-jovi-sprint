package sprint

func PascalsTriangle(n int) [][]int {
	result := [][]int{}
	for i := 0; i < n; i++ {
		result = append(result, []int{})
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				result[i] = append(result[i], 1)
			} else {
				result[i] = append(result[i], result[i-1][j-1]+result[i-1][j])
			}
		}
	}
	return result

}
