package main

import "fmt"

func repeatedCharacter(s string) byte {
	var res byte
	hash := map[byte]bool{}
	for _, c := range s {
		_, ok := hash[byte(c)]
		if ok {
			res = byte(c)
			break
		}
		hash[byte(c)] = true
	}
	return res
}

func main() {
	fmt.Println(string(repeatedCharacter("eesll")))
}
