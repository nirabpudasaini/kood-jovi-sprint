package sprint

func BalanceOut(arr []bool) []bool {
	count := 0
	for _, v := range arr {
		if v {
			count++
		} else {
			count--
		}
	}
	if count == 0 {
		return arr
	} else if count > 0 {
		for i := 0; i <= count; i++ {
			arr = append(arr, false)
		}
	} else {
		count = count * -1
		for i := 0; i < count; i++ {
			arr = append(arr, true)
		}
	}
	return arr
}
