package sprint

import "fmt"

func StrCompress(s string) string {
	if s == "" {
		return ""
	}
	lastLetter := s[0]
	consucutiveCount := 1
	result := ""
	for i := 1; i < len(s); i++ {
		if s[i] == lastLetter {
			consucutiveCount++
		} else {
			if consucutiveCount == 1 {
				result = result + string(s[i-1])
			} else {
				result = result + fmt.Sprintf("%d%s", consucutiveCount, string(s[i-1]))
				consucutiveCount = 1
			}
		}
		lastLetter = s[i]
	}
	if consucutiveCount == 1 {
		result = result + string(lastLetter)
	} else {
		result = result + fmt.Sprintf("%d%s", consucutiveCount, string(lastLetter))
	}

	return result
}
