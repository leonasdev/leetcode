package leetcode

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, t := range tokens {
		switch t {
		case "+":
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			stack = append(stack, b+a)
		case "-":
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			stack = append(stack, b-a)
		case "*":
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			stack = append(stack, b*a)
		case "/":
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			stack = append(stack, b/a)

		default: // numbers
			num, _ := strconv.Atoi(t)
			stack = append(stack, num)
		}
	}
	return stack[len(stack)-1]
}
