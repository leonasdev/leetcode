package leetcode

func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]
	for _, p := range prices {
		if p-minPrice > maxProfit {
			maxProfit = p - minPrice
		}
		if p < minPrice {
			minPrice = p
		}
	}
	return maxProfit
}
