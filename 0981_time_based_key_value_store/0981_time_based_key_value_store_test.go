package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeMap(t *testing.T) {
	obj := Constructor()
	obj.Set("foo", "bar", 1)

	assert.Equal(t, "bar", obj.Get("foo", 1))
	assert.Equal(t, "bar", obj.Get("foo", 3))

	obj.Set("foo", "bar2", 4)

	assert.Equal(t, "bar2", obj.Get("foo", 4))
	assert.Equal(t, "bar2", obj.Get("foo", 5))

	obj.Set("foo", "bar3", 5)

	assert.Equal(t, "bar3", obj.Get("foo", 5))

	obj.Set("foo", "bar3", 5)

	assert.Equal(t, "bar3", obj.Get("foo", 5))
}
