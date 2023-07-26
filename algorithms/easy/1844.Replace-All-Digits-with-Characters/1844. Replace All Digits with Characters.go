package main

import "fmt"

func replaceDigits(s string) string {
	res := []byte(s)
	for i := 1; i < len(s); i += 2 {
		res[i] = s[i-1] + s[i] - byte('0')
	}
	return string(res)
}

func main() {
	fmt.Println(replaceDigits("a1c1e1"))
	fmt.Println(replaceDigits("a1b2c3d4e"))
}
