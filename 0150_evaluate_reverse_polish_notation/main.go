package main

import (
	"fmt"
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

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))                                             // 9
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))                                            // 6
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})) // 22
}
