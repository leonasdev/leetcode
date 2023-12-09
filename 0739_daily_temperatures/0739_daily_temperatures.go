package leetcode

func dailyTemperatures(temperatures []int) []int {
	idxStack := []int{}
	res := make([]int, len(temperatures))

	for i, temp := range temperatures {
		for len(idxStack) > 0 && temp > temperatures[idxStack[len(idxStack)-1]] {
			res[idxStack[len(idxStack)-1]] = i - idxStack[len(idxStack)-1]
			idxStack = idxStack[:len(idxStack)-1]
		}
		idxStack = append(idxStack, i)
	}

	return res
}
