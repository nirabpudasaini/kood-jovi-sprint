package sprint

func ArrAny(f func(string) bool, a []string) bool {
	for _, v := range a {
		if f(v) {
			return true
		}
	}
	return false
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
