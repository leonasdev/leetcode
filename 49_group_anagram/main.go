package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	arrToAnagrams := map[[26]int][]string{}

	for _, str := range strs {
		arr := [26]int{}
		for _, ch := range str {
			arr[ch-'a']++
		}
		arrToAnagrams[arr] = append(arrToAnagrams[arr], str)
	}

	result := [][]string{}
	for _, strs := range arrToAnagrams {
		result = append(result, strs)
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs)) // [["bat"],["nat","tan"],["ate","eat","tea"]]

	strs = []string{}
	fmt.Println(groupAnagrams(strs)) // []
}
