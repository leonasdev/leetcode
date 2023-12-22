package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	codec := Codec{}
	strs := []string{"lint", "code", "love", "you"}
	assert.Equal(t, strs, codec.Decode(codec.Encode(strs)))

	codec = Codec{}
	strs = []string{"lintabcdefghijklmn", "code", "love", "you"}
	assert.Equal(t, strs, codec.Decode(codec.Encode(strs)))
}
