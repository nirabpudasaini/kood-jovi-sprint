package sprint

func FromRoman(s string) int {
	numMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	result := 0
	for _, v := range s {
		result = result + numMap[string(v)]
	}

	return result
}
