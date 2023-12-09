package leetcode

import (
	"math"
	"slices"
)

func minEatingSpeed(piles []int, h int) int {
	maxPile := slices.Max(piles)
	lo, hi := 1, maxPile
	minRes := maxPile

	for lo <= hi {
		mid := lo + (hi-lo)/2
		spend := 0
		for _, p := range piles {
			spend += int(math.Ceil(float64(p) / float64(mid)))
		}
		if spend <= h {
			hi = mid - 1
			minRes = mid
		} else {
			lo = mid + 1
		}
	}

	return minRes
}
