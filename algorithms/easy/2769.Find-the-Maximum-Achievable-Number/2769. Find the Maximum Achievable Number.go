package main

import "fmt"

func theMaximumAchievableX(num int, t int) int {
	return (t * 2) + num
}

func main() {
	fmt.Println(theMaximumAchievableX(4, 1))
}
