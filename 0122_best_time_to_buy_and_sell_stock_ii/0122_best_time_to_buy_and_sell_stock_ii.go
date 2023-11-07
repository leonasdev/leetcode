package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	maxProfit := 0
	minPrice := prices[0]
	for _, p := range prices {
		if p-minPrice > profit {
			profit = p - minPrice
			maxProfit += profit
		}

		if p < minPrice {
			minPrice = p
			profit = 0
		}
	}

	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices)) // 7

	prices = []int{7, 1, 3, 5, 6, 4}
	fmt.Println(maxProfit(prices)) // 5

	prices = []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices)) // 4

	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices)) // 0
}
