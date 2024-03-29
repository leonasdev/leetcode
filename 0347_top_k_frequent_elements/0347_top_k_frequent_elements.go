package leetcode

func topKFrequent(nums []int, k int) []int {
	numToFre := map[int]int{}
	for _, n := range nums {
		numToFre[n]++
	}

	countSlice := make([][]int, len(nums)+1) // +1 since the number may appear zero times
	for num, fre := range numToFre {
		countSlice[fre] = append(countSlice[fre], num)
	}

	result := []int{}
	for i := len(countSlice) - 1; i >= 0 && len(result) < k; i-- {
		result = append(result, countSlice[i]...)
	}

	return result
}
