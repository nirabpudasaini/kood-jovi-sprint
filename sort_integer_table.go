package sprint

func SortIntegerTable(table []int) []int {
	size := len(table)
	for i := 0; i < (size - 1); i++ {
		for j := 0; j < size-i-1; j++ {
			if table[j] > table[j+1] {
				table[j+1], table[j] = table[j], table[j+1]
			}
		}
	}
	return table
}
