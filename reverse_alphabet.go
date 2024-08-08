package sprint

func ReverseAlphabet(step int) string {
	if step <= 0 || step > 26 {
		step = 1
	}
	result := ""
	for i := 0; i < 26; i += step {
		result += string(rune(122 - i))
	}
	return result

}
