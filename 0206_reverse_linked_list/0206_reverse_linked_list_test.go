package leetcode

import (
	"testing"

	structrue "github.com/leonasdev/leetcode/structure"
	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	t.Run("NormalCase", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		list := structrue.IntsToList(nums)

		list = reverseList(list)

		assert.Equal(t, []int{5, 4, 3, 2, 1}, structrue.ListToInts(list))
	})

	t.Run("Empty", func(t *testing.T) {
		nums := []int{}
		list := structrue.IntsToList(nums)

		list = reverseList(list)

		assert.Equal(t, []int{}, structrue.ListToInts(list))
	})
}
