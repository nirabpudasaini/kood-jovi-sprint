package sprint

func SubstrIndex(s, toFind string) int {
	sArr := []rune(s)
	sLen := len(sArr)
	toFindLen := len(toFind)
	for i := 0; i < sLen; i++ {
		if s[i] == toFind[0] {
			if i+toFindLen < sLen {
				if s[i:i+toFindLen] == toFind {
					return i
				}
			}

		}
	}
	return -1
}
