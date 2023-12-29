package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanAttendMeetings(t *testing.T) {
	assert.False(t, canAttendMeetings([]Interval{
		{
			Start: 0,
			End:   30,
		},
		{
			Start: 5,
			End:   10,
		},
		{
			Start: 15,
			End:   20,
		},
	}))

	assert.True(t, canAttendMeetings([]Interval{
		{
			Start: 5,
			End:   8,
		},
		{
			Start: 9,
			End:   15,
		},
	}))
}
