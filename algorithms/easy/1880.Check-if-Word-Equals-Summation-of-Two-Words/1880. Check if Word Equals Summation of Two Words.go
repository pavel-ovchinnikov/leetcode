package main

import "fmt"

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	string2int := func(word string) (res int) {
		for _, c := range word {
			res = res*10 + int(c-'a')
		}
		return
	}

	return string2int(firstWord)+string2int(secondWord) == string2int(targetWord)
}

func main() {
	fmt.Println(isSumEqual("aa", "ab", "aaab"))
}
