package leetcode

func longestConsecutive(nums []int) int {
	longest := 0
	m := map[int]bool{}

	for _, num := range nums {
		m[num] = true
	}

	for num := range m {
		if m[num-1] {
			continue
		}

		count := 1
		for m[num+count] {
			count++
		}
		longest = max(longest, count)
	}

	return longest
}
