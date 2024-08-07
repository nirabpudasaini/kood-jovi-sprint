package sprint

func Season(month string) string {
	months := [12]string{"jan", "feb", "dec", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov"}
	switch whatIndexItIsIn(month, months[:]) {
	case 0, 1, 2:
		return "winter"
	case 3, 4, 5:
		return "spring"
	case 6, 7, 8:
		return "summer"
	case 9, 10, 11:
		return "autumn"
	default:
		return "invalid input: " + month
	}
}

func whatIndexItIsIn(month string, months []string) int {
	for i, v := range months {
		if v == month {
			return i
		}
	}
	return -1
}
