package main

import (
	"fmt"
	"math/bits"
)

func minBitFlips(start int, goal int) int {
	return bits.OnesCount(uint(start ^ goal))
}

func main() {
	fmt.Println(minBitFlips(10, 7))
}
