package sprint

func IsSorted(f func(a, b string) int, arr []string) bool {
	isAscending := f(arr[0], arr[1]) == -1
	for i := 0; i < len(arr)-1; i++ {
		if isAscending {
			if f(arr[i], arr[i+1]) == 1 {
				return false
			}
		} else {
			if f(arr[i], arr[i+1]) == -1 {
				return false
			}
		}
	}
	return true
}
