package structrue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListToInts(t *testing.T) {
	t.Run("NormalCase", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
				},
			},
		}

		assert.Equal(t, []int{1, 2, 3}, ListToInts(head))
	})

	t.Run("OneElement", func(t *testing.T) {
		head := &ListNode{}

		assert.Equal(t, []int{0}, ListToInts(head))
	})

	t.Run("Nil", func(t *testing.T) {
		assert.Equal(t, []int{}, ListToInts(nil))
	})
}

func TestIntsToList(t *testing.T) {
	t.Run("NormalCase", func(t *testing.T) {
		nums := []int{1, 2, 3}
		head := IntsToList(nums)

		i := 0
		for head != nil {
			assert.Equal(t, nums[i], head.Val)
			head = head.Next
			i++
		}
	})

	t.Run("OneElement", func(t *testing.T) {
		nums := []int{0}
		head := IntsToList(nums)

		i := 0
		for head != nil {
			assert.Equal(t, nums[i], head.Val)
			head = head.Next
			i++
		}
	})

	t.Run("Nil", func(t *testing.T) {
		nums := []int{}
		head := IntsToList(nums)

		assert.Nil(t, head)
	})
}
