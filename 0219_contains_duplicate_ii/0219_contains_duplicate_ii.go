package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	record := map[int]int{}
	for i, v := range nums {
		if j, ok := record[v]; ok && i-j <= k {
			return true
		}
		record[v] = i
	}
	return false
}
