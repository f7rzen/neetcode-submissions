func maxProfit(prices []int) int {
	left := 0
	maxProf := 0

	for right := 1; right < len(prices); right++ {
		if prices[right] < prices[left] {
			left = right
		} else {
			profit := prices[right] - prices[left]
			if profit > maxProf {
				maxProf = profit
			}
		}
	}
	return maxProf
}