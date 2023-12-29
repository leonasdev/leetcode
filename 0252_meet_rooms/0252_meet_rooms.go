package leetcode

import "sort"

type Interval struct {
	Start, End int
}

func canAttendMeetings(intervals []Interval) bool {
	// Write your code here
	sort.Slice(intervals, func(i, j int) bool { return intervals[i].Start < intervals[j].End })

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i].End > intervals[i+1].Start {
			return false
		}
	}

	return true
}
