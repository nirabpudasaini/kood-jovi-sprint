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
	for i := 0; i < len(s)-1; i++ {
		if numMap[string(s[i])] >= numMap[string(s[i+1])] {
			result = result + numMap[string(s[i])]
		} else {
			result = result - numMap[string(s[i])]
		}
	}
	result = result + numMap[string(s[len(s)-1])]

	return result
}
