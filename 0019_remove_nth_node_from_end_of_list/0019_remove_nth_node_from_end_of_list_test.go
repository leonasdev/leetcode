package leetcode

import (
	"testing"

	"github.com/leonasdev/leetcode/structure"
	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		head := structure.IntsToList([]int{1, 2, 3, 4, 5})
		assert.Equal(t, []int{1, 2, 3, 5}, structure.ListToInts(removeNthFromEnd(head, 2)))
	})

	t.Run("2", func(t *testing.T) {
		head := structure.IntsToList([]int{1})
		assert.Equal(t, []int{}, structure.ListToInts(removeNthFromEnd(head, 1)))
	})

	t.Run("1", func(t *testing.T) {
		head := structure.IntsToList([]int{1, 2})
		assert.Equal(t, []int{1}, structure.ListToInts(removeNthFromEnd(head, 1)))
	})
}
