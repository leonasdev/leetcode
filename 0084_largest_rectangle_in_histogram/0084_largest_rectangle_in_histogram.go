package leetcode

type Histogram struct {
	Index  int
	Height int
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := []Histogram{}

	for i, height := range heights {
		lastIdx := i
		for len(stack) > 0 && height < stack[len(stack)-1].Height {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			area := (i - top.Index) * top.Height
			maxArea = max(maxArea, area)
			lastIdx = top.Index
		}

		stack = append(stack, Histogram{
			Index:  lastIdx,
			Height: height,
		})
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		area := (len(heights) - top.Index) * top.Height
		maxArea = max(maxArea, area)
	}

	return maxArea
}
