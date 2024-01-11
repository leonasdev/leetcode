package leetcode

func subsets(nums []int) [][]int {
	res := [][]int{}
	subset := []int{}
	dfs := func(int) {}
	dfs = func(i int) {
		if i == len(nums) {
			copied := make([]int, len(subset))
			copy(copied, subset)
			res = append(res, copied)
			return
		}
		subset = append(subset, nums[i])
		dfs(i + 1)
		subset = subset[:len(subset)-1]
		dfs(i + 1)
	}
	dfs(0)

	return res
}
