package sprint

func ArrMap(f func(int) bool, a []int) []bool {
	result := []bool{}
	for _, v := range a {
		result = append(result, f(v))
	}
	return result
}
