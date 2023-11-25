package leetcode

import (
	"testing"

	"github.com/leonasdev/leetcode/structure"
	"github.com/stretchr/testify/assert"
)

func TestReorderList(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		head := structure.IntsToList([]int{1, 2, 3, 4})
		reorderList(head)

		assert.Equal(t, []int{1, 4, 2, 3}, structure.ListToInts(head))
	})

	t.Run("2", func(t *testing.T) {
		head := structure.IntsToList([]int{1, 2, 3, 4, 5})
		reorderList(head)

		assert.Equal(t, []int{1, 5, 2, 4, 3}, structure.ListToInts(head))
	})
}
