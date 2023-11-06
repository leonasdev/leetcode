package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi]
	index := lo
	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[index], arr[hi] = arr[hi], arr[index]
	return index
}

func qs(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}
	index := partition(arr, lo, hi)
	qs(arr, lo, index-1)
	qs(arr, index+1, hi)
}

func quickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}

func main() {
	for k := 0; k < 10; k++ {
		s := []int{}
		s2 := make([]int, 100)
		for i := 0; i < 100; i++ {
			s = append(s, rand.Intn(100))
		}
		copy(s2, s)
		sort.Ints(s)
		quickSort(s2)
		for j := 0; j < len(s); j++ {
			if s[j] != s2[j] {
				fmt.Println("wrong!!")
				return
			}
		}
		fmt.Println("ok")
	}
}
