package leetcode

import "slices"

func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)

	res := 0
	for i := 0; i < len(s) && res < len(g); i++ {
		if s[i] >= g[res] {
			res++
		}
	}

	return res
}
