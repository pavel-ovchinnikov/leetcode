package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	morseCode := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.",
		"--.", "....", "..", ".---", "-.-", ".-..",
		"--", "-.", "---", ".--.", "--.-", ".-.",
		"...", "-", "..-", "...-", ".--", "-..-",
		"-.--", "--..",
	}

	set := make(map[string]bool)
	for _, w := range words {
		key := ""
		for _, c := range w {
			key += morseCode[c-'a']
		}
		set[key] = true
	}
	return len(set)
}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
