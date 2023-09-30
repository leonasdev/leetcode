package main

import "fmt"

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

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))         // 4
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})) // 9
}
