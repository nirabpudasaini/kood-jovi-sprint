package sprint

func ArrCountIf(f func(string) bool, a []string) int {
	count := 0
	for _, v := range a {
		if f(v) {
			count++
		}
	}
	return count
}

func IsUpper(s string) bool {
	for _, v := range s {
		if v < 'A' || v > 'Z' {
			return false
		}
	}
	return true
}

func IsAlphanumeric(s string) bool {
	for _, v := range s {
		if !((v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9')) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, v := range s {
		if v < 'a' || v > 'z' {
			return false
		}
	}
	return true
}

func IsNumeric(s string) bool {
	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}
