package main

import "fmt"

func decodeMessage(key string, message string) string {
	table := map[rune]rune{' ': ' '}
	for _, c := range key {
		if _, ok := table[c]; !ok {
			table[c] = rune(len(table) - 1 + 'a')
		}
	}

	index := 0
	buf := make([]rune, len(message))
	for _, ch := range message {
		buf[index] = table[ch]
		index++
	}

	return string(buf)
}

func main() {
	fmt.Println(decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
}
