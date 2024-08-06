package sprint

func ShiftBy(r rune, step int) rune {
	shifted := rune(int(r) + step)
	for shifted > 122 {
		shifted = 97 + (shifted - 122)
	}
	return shifted
}
