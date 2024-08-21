package sprint

func LongestClimb(arr []int) []int {
	longestArr := []int{arr[0]}
	currentArr := []int{arr[0]}
	lastElement := arr[0]
	for i := 1; i < len(arr); i++ {
		if lastElement <= arr[i] {
			currentArr = append(currentArr, arr[i])
		} else {
			currentArr = []int{arr[i]}
		}
		if len(currentArr) > len(longestArr) {
			longestArr = currentArr
		}
		lastElement = arr[i]
	}
	return longestArr
}
