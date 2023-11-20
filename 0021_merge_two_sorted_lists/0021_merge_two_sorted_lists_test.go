package leetcode

import (
	"testing"

	structrue "github.com/leonasdev/leetcode/structure"
	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := structrue.IntsToList([]int{1, 2, 4})
	list2 := structrue.IntsToList([]int{1, 3, 4})
	res := mergeTwoLists(list1, list2)
	assert.Equal(t, []int{1, 1, 2, 3, 4, 4}, structrue.ListToInts(res))
}
