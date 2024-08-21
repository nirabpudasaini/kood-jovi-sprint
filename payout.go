package sprint

func Payout(amount int, denominations []int) (payout []int) {
	payout = []int{}
	denominations = sortDescending(denominations)
	for amount > 0 {
		if amount < denominations[len(denominations)-1] {
			return []int{}
		}
		for i := 0; i < len(denominations); i++ {
			if amount >= denominations[i] {
				payout = append(payout, denominations[i])
				amount -= denominations[i]
				break
			}
		}
	}
	return payout
}

func sortDescending(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
