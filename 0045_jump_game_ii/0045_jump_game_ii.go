package leetcode

import "fmt"

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func jump(nums []int) int {
	farest := 0
	end := 0
	jumps := 0
	for i := 0; i < len(nums)-1; i++ {
		farest = maxInt(farest, i+nums[i])

		if i == end {
			jumps++
			end = farest
		}
	}
	return jumps
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums)) // 2

	nums = []int{1, 2, 3}
	fmt.Println(jump(nums)) // 2

	nums = []int{1, 1, 1, 1}
	fmt.Println(jump(nums)) // 3

	nums = []int{6, 1, 1, 1}
	fmt.Println(jump(nums)) // 1

	nums = []int{3, 2, 3, 1, 1, 1, 1}
	fmt.Println(jump(nums)) // 3

	nums = []int{7, 2, 3, 1, 1, 1, 1}
	fmt.Println(jump(nums)) // 1

	nums = []int{0}
	fmt.Println(jump(nums)) // 0

	nums = []int{8, 1}
	fmt.Println(jump(nums)) // 1
}
