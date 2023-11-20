package leetcode

func twoSum(numbers []int, target int) []int {
	lo, hi := 0, len(numbers)-1

	for { // The tests are generated such that there is exactly one solution.
		sum := numbers[lo] + numbers[hi]
		if sum > target {
			hi--
		} else if sum < target {
			lo++
		} else {
			return []int{lo + 1, hi + 1}
		}
	}
}
