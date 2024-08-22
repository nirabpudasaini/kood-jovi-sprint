package sprint

func DigitalRoot(n int) int {
	root := n
	for root > 10 {
		root = 0
		for n > 0 {
			root = root + n%10
			n = n / 10
		}
		n = root
	}
	return n
}
