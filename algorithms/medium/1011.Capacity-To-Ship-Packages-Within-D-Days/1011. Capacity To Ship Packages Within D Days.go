package leetcode

func shipWithinDays(weights []int, days int) int {
	minWeightShip, maxWeightShip := 0, 0
	weightShip := 0

	for _, weight := range weights {
		minWeightShip = Max(weight, minWeightShip)
		maxWeightShip += weight
	}

	for minWeightShip <= maxWeightShip {
		mid := minWeightShip + (maxWeightShip-minWeightShip)/2
		if canShip(weights, days, mid) {
			weightShip = mid
			maxWeightShip = mid - 1
		} else {
			minWeightShip = mid + 1
		}

	}
	return weightShip
}

func canShip(weights []int, expectedDays int, weightShip int) bool {
	sumWeight := 0
	days := 1
	for _, weight := range weights {
		if weight > weightShip || days > expectedDays {
			return false
		}

		if weight+sumWeight > weightShip {
			sumWeight = weight
			days += 1
		} else {
			sumWeight += weight
		}
	}
	return days <= expectedDays
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
