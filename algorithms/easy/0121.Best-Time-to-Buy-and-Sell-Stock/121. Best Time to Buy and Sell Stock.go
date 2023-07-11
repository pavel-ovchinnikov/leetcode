package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	price := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-price > profit {
			profit = prices[i] - price
		} else if price > prices[i] {
			price = prices[i]
		}
	}
	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
