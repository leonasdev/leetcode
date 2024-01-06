package leetcode

type ValTime struct {
	Val       string
	Timestamp int
}

type TimeMap struct {
	Map map[string][]*ValTime
}

func Constructor() TimeMap {
	return TimeMap{
		Map: map[string][]*ValTime{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.Map[key] = append(this.Map[key], &ValTime{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	valTimes, exists := this.Map[key]
	if !exists {
		return ""
	}

	lo, hi, res := 0, len(valTimes)-1, ""
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if valTimes[mid].Timestamp <= timestamp {
			lo = mid + 1
			res = valTimes[mid].Val
		} else {
			hi = mid - 1
		}
	}

	return res
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
