package main

import (
	"fmt"
	"math"
	"strings"
)

func mostWordsFound(sentences []string) int {
	var result float64 = 0.0

	for _, sentence := range sentences {
		result = math.Max(float64(strings.Count(sentence, " "))+1, float64(result))
	}

	return int(result)
}

func main() {
	fmt.Println(mostWordsFound([]string{
		"alice and bob love leetcode",
		"i think so too",
		"this is great thanks very much"}))
}
