package leetcode

import "sort"

type Car struct {
	Position int
	Speed    int
}

func carFleet(target int, position []int, speed []int) int {
	cars := []Car{}
	for i := 0; i < len(position); i++ {
		cars = append(cars, Car{
			Position: position[i],
			Speed:    speed[i],
		})
	}
	sort.Slice(cars, func(i, j int) bool { return cars[i].Position > cars[j].Position })

	maxTime := 0.0
	count := 0
	for _, car := range cars {
		timeToTarget := float64(target-car.Position) / float64(car.Speed)
		if timeToTarget > maxTime {
			count++
			maxTime = timeToTarget
		}
	}

	return count
}
