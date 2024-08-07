package sprint

func BetweenLimits(from, to rune) string {
	tempto := 'a'
	result := ""
	if from > to {
		tempto = to
		to = from
		from = tempto
	}
	for i := 1; i < int(to-from); i++ {
		result += string(from + rune(i))
	}
	return result
}
