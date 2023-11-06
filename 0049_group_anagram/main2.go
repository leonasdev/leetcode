package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	arrToStrs := map[[26]int][]string{}
	for _, str := range strs {
		letterArr := [26]int{}
		for _, ch := range str {
			letterArr[ch-'a']++
		}
		arrToStrs[letterArr] = append(arrToStrs[letterArr], str)
	}

	result := [][]string{}
	for _, strs := range arrToStrs {
		result = append(result, strs)
	}
	return result
}

func main() {
	fmt.Println(groupAnagram([]string{"eat", "tea", "tan", "ate", "nat", "bat"})) // [["bat"],["nat","tan"],["ate","eat","tea"]]
	fmt.Println(groupAnagram([]string{}))                                         // []
	fmt.Println(groupAnagram([]string{"a"}))                                      // [["a"]]
}
