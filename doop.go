package sprint

func Doop(a int, op string, b int) int {
	if b == 0 {
		return 0
	}
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	case "%":
		return a % b
	default:
		return 0
	}
}
