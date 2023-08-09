package main

import (
	"fmt"
	"strings"
)

// func splitWordsBySeparator(words []string, separator byte) []string {
// 	res := []string{}
// 	for _, w := range words {
// 		for _, splitWord := range strings.Split(w, string(separator)) {
// 			if len(splitWord) != 0 {
// 				res = append(res, splitWord)
// 			}
// 		}

// 	}
// 	return res
// }

func splitWordsBySeparator(words []string, separator byte) []string {
	res := make([]string, 0)
	builder := strings.Builder{}

	for _, word := range words {
		for _, c := range word {
			if c == rune(separator) {
				if builder.Len() == 0 {
					continue
				}
				res = append(res, builder.String())
				builder.Reset()
			} else {
				builder.WriteByte(byte(c))
			}
		}

		if builder.Len() != 0 {
			res = append(res, builder.String())
			builder.Reset()
		}
	}

	return res
}

func main() {
	fmt.Println(splitWordsBySeparator([]string{"one.two.three", "four.five", "six"}, '.'))
}
