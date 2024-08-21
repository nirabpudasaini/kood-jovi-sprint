package sprint

func ToRoman(num int) string {
	result := ""
	numMap := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	dem := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []int{}
	for num > 0 {
		for i := 0; i < len(dem); i++ {
			if num >= dem[i] {
				roman = append(roman, dem[i])
				num -= dem[i]
				break
			}
		}
	}
	for _, v := range roman {
		result = result + numMap[v]
	}

	return result

}
