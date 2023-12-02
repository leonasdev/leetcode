package leetcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{-1, 0, 1, 2, -1, -4}
		want := [][]int{{-1, -1, 2}, {-1, 0, 1}}
		actual := threeSum(nums)

		assert.Equal(t, len(want), len(actual))
		sort.Slice(want, func(i, j int) bool {
			return want[i][0] < want[j][0]
		})
		sort.Slice(actual, func(i, j int) bool {
			return actual[i][0] < actual[j][0]
		})
		for i := 0; i < len(want); i++ {
			assert.ElementsMatch(t, want[i], actual[i])
		}
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{0, 1, 1}
		want := [][]int{}
		actual := threeSum(nums)

		assert.Equal(t, len(want), len(actual))
		sort.Slice(want, func(i, j int) bool {
			return want[i][0] < want[j][0]
		})
		sort.Slice(actual, func(i, j int) bool {
			return actual[i][0] < actual[j][0]
		})
		for i := 0; i < len(want); i++ {
			assert.ElementsMatch(t, want[i], actual[i])
		}
	})

	t.Run("3", func(t *testing.T) {
		nums := []int{0, 0, 0}
		want := [][]int{{0, 0, 0}}
		actual := threeSum(nums)

		assert.Equal(t, len(want), len(actual))
		sort.Slice(want, func(i, j int) bool {
			return want[i][0] < want[j][0]
		})
		sort.Slice(actual, func(i, j int) bool {
			return actual[i][0] < actual[j][0]
		})
		for i := 0; i < len(want); i++ {
			assert.ElementsMatch(t, want[i], actual[i])
		}
	})
}
