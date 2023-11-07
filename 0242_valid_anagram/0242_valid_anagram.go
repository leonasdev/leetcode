package leetcode

import "fmt"

func isAnagram(s string, t string) bool {
	letterToCount := map[rune]int{}
	for _, v := range s {
		letterToCount[v]++
	}
	for _, v := range t {
		letterToCount[v]--
	}
	for _, v := range letterToCount {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // true
	fmt.Println(isAnagram("rat", "car"))         // false
	fmt.Println(isAnagram("你好不好", "好不好你"))       // true
	fmt.Println(isAnagram("你好不", "好不好你"))        // false
	fmt.Println(isAnagram("你好不", "不好你"))         // true
	fmt.Println(isAnagram("你好不", "不你"))          // false
}
