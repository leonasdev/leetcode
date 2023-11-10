package leetcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	actual := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	want := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}
	sort.SliceStable(actual, func(i, j int) bool {
		return len(actual[i]) < len(actual[j])
	})
	sort.SliceStable(want, func(i, j int) bool {
		return len(actual[i]) < len(actual[j])
	})
	for i := range want {
		assert.ElementsMatch(t, want[i], actual[i])
	}
	assert.Equal(t, [][]string{}, groupAnagrams([]string{}))
}
