package main

import (
	"fmt"
	"math"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlateFreq, wordsFreq := [26]int{}, [26]int{}
	freq(&licensePlateFreq, licensePlate)
	fmt.Println(licensePlateFreq)

	shortestLen := math.MaxInt
	res := ""
	for _, word := range words {
		clear(&wordsFreq)
		freq(&wordsFreq, word)
		if check(licensePlateFreq, wordsFreq) && len(word) < shortestLen {
			res = word
			fmt.Println(wordsFreq)
			shortestLen = len(word)
		}
	}

	return res
}

func check(freq1, freq2 [26]int) bool {
	for i := 0; i < 26; i++ {
		if freq1[i] > freq2[i] {
			return false
		}
	}
	return true
}

func freq(freq *[26]int, s string) {
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			char += 32
		}
		if char >= 'a' && char <= 'z' {
			freq[char-'a']++
		}
	}
}

func clear(freq *[26]int) {
	for i := 0; i < 26; i++ {
		freq[i] = 0
	}
}

func main() {
	fmt.Println(shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
	fmt.Println(shortestCompletingWord("Ah71752", []string{"suggest", "letter", "of", "husband", "easy", "education", "drug", "prevent", "writer", "old"}))
}
