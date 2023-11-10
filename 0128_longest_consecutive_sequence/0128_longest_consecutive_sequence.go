package leetcode

func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, v := range nums {
		set[v] = true
	}

	longest := 0
	for v := range set {
		if set[v-1] {
			continue
		}
		count := 0
		for set[v+count] {
			count++
		}
		if count > longest {
			longest = count
		}
	}

	return longest
}
