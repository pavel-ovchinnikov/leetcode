package main

import "fmt"

func average(salary []int) float64 {
	maxSalary := salary[0]
	minSalary := salary[0]
	var sum float64
	for _, sal := range salary {
		sum += float64(sal)
		maxSalary = max(maxSalary, sal)
		minSalary = min(minSalary, sal)
	}
	return (sum - float64(maxSalary) - float64(minSalary)) / (float64(len(salary)) - 2)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(average([]int{4000, 3000, 1000, 2000}))
}
