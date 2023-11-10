package leetcode

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	for k := 0; k < 10; k++ {
		s := []int{}
		s2 := make([]int, 100)
		for i := 0; i < 100; i++ {
			s = append(s, rand.Intn(100))
		}

		copy(s2, s)

		sort.Ints(s)
		quickSort(s2)

		assert.Equal(t, s, s2)
	}
}
