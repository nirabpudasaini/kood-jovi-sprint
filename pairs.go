package sprint

import "fmt"

func Pairs() string {
	result := ""
	for i := 0; i <= 99; i++ {
		for j := 1; j <= 99; j++ {
			if i == j {
				continue
			}
			if i == 99 && j == 98 {
				result += fmt.Sprintf("%02d %02d", i, j)
				continue
			}
			result += fmt.Sprintf("%02d %02d, ", i, j)
		}
	}
	fmt.Println(result)
	return result
}
