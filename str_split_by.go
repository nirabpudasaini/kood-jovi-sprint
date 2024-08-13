package sprint

func StrSplitBy(s, sep string) []string {
	sArr := []rune(s)
	sepArr := []rune(sep)
	result := []string{}
	sepLen := len(sepArr)
	sLen := len(sArr)
	word := ""

	for i := 0; i < len(sArr); i++ {
		if s[i] == sep[0] {
			if i+sepLen < sLen {
				if s[i:(i+sepLen)] == sep {
					result = append(result, word)
					i = i + sepLen - 1
					word = ""
					continue
				}
			}
		}
		word += string(sArr[i])
	}
	result = append(result, word)
	return result
}
