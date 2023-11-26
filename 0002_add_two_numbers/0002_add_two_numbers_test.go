package leetcode

import (
	"testing"

	"github.com/leonasdev/leetcode/structure"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		l1 := structure.IntsToList([]int{2, 4, 3})
		l2 := structure.IntsToList([]int{5, 6, 4})
		assert.Equal(t, []int{7, 0, 8}, structure.ListToInts(addTwoNumbers(l1, l2)))
	})

	t.Run("2", func(t *testing.T) {
		l1 := structure.IntsToList([]int{0})
		l2 := structure.IntsToList([]int{0})
		assert.Equal(t, []int{0}, structure.ListToInts(addTwoNumbers(l1, l2)))
	})

	t.Run("3", func(t *testing.T) {
		l1 := structure.IntsToList([]int{9, 9, 9, 9, 9, 9, 9})
		l2 := structure.IntsToList([]int{9, 9, 9, 9})
		assert.Equal(t, []int{8, 9, 9, 9, 0, 0, 0, 1}, structure.ListToInts(addTwoNumbers(l1, l2)))
	})
}
