package main

import (
	"fmt"
	"math"
)

func minSpeedOnTime(dist []int, hour float64) int {
	l, r := 0, 10000000
	for l < r {
		mid := (l + r) / 2
		if timeSpeed(dist, mid) > hour {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if hour < timeSpeed(dist, l) {
		l = -1
	}
	return l
}

func timeSpeed(dist []int, speed int) (curTime float64) {
	n := len(dist)
	for i := 0; i < n; i++ {
		curTime = math.Ceil(curTime)
		curTime += float64(dist[i]) / float64(speed)
	}
	return curTime
}

func main() {
	fmt.Println(minSpeedOnTime([]int{1, 3, 2}, 6))
	fmt.Println(minSpeedOnTime([]int{1, 3, 2}, 2.7))
	fmt.Println(minSpeedOnTime([]int{1, 3, 2}, 1.9))
}
