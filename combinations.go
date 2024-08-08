package sprint

import "fmt"

func Combinations() string {
	result := ""
	for i := 0; i < 1000; i++ {
		numberString := fmt.Sprintf("%03d", i)
		matchFlag := true
		highestValue := -1
		for _, v := range numberString {
			if int(v) > highestValue {
				highestValue = int(v)
			} else {
				matchFlag = false
				break
			}
		}
		if matchFlag {
			if numberString == "789" {
				result += fmt.Sprintf("%v", numberString)
				continue
			}
			result += fmt.Sprintf("%v,", numberString)
		}

	}
	return result
}
