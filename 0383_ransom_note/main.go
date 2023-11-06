package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	letterToCount := map[rune]int{}
	for _, c := range magazine {
		letterToCount[c]++
	}
	for _, c := range ransomNote {
		letterToCount[c]--
		if letterToCount[c] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))    // false
	fmt.Println(canConstruct("aa", "ab"))  // false
	fmt.Println(canConstruct("aa", "aab")) // true
}
