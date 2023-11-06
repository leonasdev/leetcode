package main

import "fmt"

func isHappy(n int) bool {
	record := map[int]bool{}
	for n > 1 {
		sum := 0
		for n > 0 {
			d := n % 10
			square := d * d
			n /= 10
			sum += square
		}
		if v := record[sum]; v {
			return false
		}
		record[sum] = true
		n = sum
	}
	return true
}

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}
