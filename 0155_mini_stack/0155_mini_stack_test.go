package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func TestMiniStack(t *testing.T) {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)

	assert.Equal(t, -3, obj.GetMin())

	obj.Pop()

	assert.Equal(t, 0, obj.Top())

	assert.Equal(t, -2, obj.GetMin())
}
