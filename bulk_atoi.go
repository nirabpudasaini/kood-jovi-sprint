package sprint

func BulkAtoi(arr []string) []int {
	results := []int{}
	for _, v := range arr {
		results = append(results, strToInt(v))
	}
	return results
}

func strToInt(s string) int {
	negative := false
	if s[0] == '-' {
		negative = true
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	result := 0
	for _, v := range []byte(s) {
		if v < '0' || v > '9' {
			return 0
		}
		result = result*10 + int(v-'0')
	}
	if negative {
		result = result * -1
	}

	return result
}
