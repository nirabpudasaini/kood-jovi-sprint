package sprint

func LongestCommonSubstr(str1, str2 string) string {
	matrix := make([][]int, len(str1)+1)
	max := 0
	maxI := 0
	for i := range matrix {
		matrix[i] = make([]int, len(str2)+1)
	}
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				matrix[i][j] = matrix[i-1][j-1] + 1
				if max < matrix[i][j] {
					max = matrix[i][j]
					maxI = i
				}
			}
		}
	}
	return str1[maxI-max : maxI]
}
