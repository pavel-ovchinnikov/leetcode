package main

import "fmt"

func minPartitions(n string) int {
	res := 0
	for _, c := range n {
		res = max(res, (int(c) - int('0')))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minPartitions("32"))
}
