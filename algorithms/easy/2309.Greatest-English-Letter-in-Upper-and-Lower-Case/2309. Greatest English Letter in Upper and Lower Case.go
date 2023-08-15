package main

import "fmt"

func greatestLetter(s string) string {
	res := make([]bool, 'z'+1)
	diff := 'a' - 'A'

	for _, letter := range s {
		res[letter] = true
	}

	for i := 'Z'; i >= 'A'; i-- {
		if res[i] && res[i+diff] {
			return string(i)
		}
	}

	return ""

}

func main() {
	fmt.Println(greatestLetter("lEeTcOdE"))
}
