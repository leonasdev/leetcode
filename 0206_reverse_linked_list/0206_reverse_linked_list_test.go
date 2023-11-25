package leetcode

import (
	"testing"

	"github.com/leonasdev/leetcode/structure"
	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	t.Run("NormalCase", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		list := structure.IntsToList(nums)

		list = reverseList(list)

		assert.Equal(t, []int{5, 4, 3, 2, 1}, structure.ListToInts(list))
	})

	t.Run("Empty", func(t *testing.T) {
		nums := []int{}
		list := structure.IntsToList(nums)

		list = reverseList(list)

		assert.Equal(t, []int{}, structure.ListToInts(list))
	})
}
