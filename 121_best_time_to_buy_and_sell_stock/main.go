package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

func maxProfit2(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
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

func maxProfit3(prices []int) int {
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

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit3(prices)) // 5

	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit3(prices)) // 0

	prices = []int{2, 4, 1}
	fmt.Println(maxProfit3(prices)) // 2

	prices = []int{2, 4, 1, 2, 2, 2}
	fmt.Println(maxProfit3(prices)) // 2
}
