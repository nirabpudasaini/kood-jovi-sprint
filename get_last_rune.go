package sprint

func GetLastRune(s string) rune {
	sRuneArray := []rune(s)
	return sRuneArray[len(sRuneArray)-1]
}
