package sprint

func SortWordArr(a []string) []string {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if rune(a[j][0]) > rune(a[j+1][0]) {
				a[j], a[j+1] = a[j+1], a[j]
			} else if rune(a[j][0]) == rune(a[j+1][0]) {
				if rune(a[j][1]) > rune(a[j+1][1]) {
					a[j], a[j+1] = a[j+1], a[j]
				}
			}
		}
	}
	return a
}
