package main

import "fmt"

func commonChars(words []string) []string {
	letters := make(map[rune]int)
	res := []string{}

	for _, с := range words[0] {
		letters[с]++
	}

	words = words[1:]

	for _, word := range words {
		temp := make(map[rune]int)
		for _, с := range word {
			if letters[с] > 0 {
				temp[с]++
				letters[с]--
			}
		}

		letters = temp
	}

	for k, v := range letters {
		for i := 0; i < v; i++ {
			res = append(res, string(k))
		}
	}

	return res
}

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
}
