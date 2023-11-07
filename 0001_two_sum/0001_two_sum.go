package leetcode

func twoSum(nums []int, target int) []int {
	visit := map[int]int{}
	for i, num := range nums {
		if v, exist := visit[target-num]; exist {
			return []int{v, i}
		}
		visit[num] = i
	}
	return nil
}
