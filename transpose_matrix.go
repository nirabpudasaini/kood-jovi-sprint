package sprint

func TransposeMatrix(matrix [][]int) [][]int {
	transpose := make([][]int, len(matrix[0]))
	for i := range transpose {
		transpose[i] = make([]int, len(matrix))
		for j := range transpose[i] {
			transpose[i][j] = matrix[j][i]
		}
	}
	return transpose
}
