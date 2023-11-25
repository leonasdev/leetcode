package leetcode

import (
	"testing"

	"github.com/leonasdev/leetcode/structure"
	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := structure.IntsToList([]int{1, 2, 4})
	list2 := structure.IntsToList([]int{1, 3, 4})
	res := mergeTwoLists(list1, list2)
	assert.Equal(t, []int{1, 1, 2, 3, 4, 4}, structure.ListToInts(res))
}
