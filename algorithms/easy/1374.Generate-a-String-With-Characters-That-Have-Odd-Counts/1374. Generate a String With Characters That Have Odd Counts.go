package main

import "fmt"

func generateTheString(n int) string {
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = 'a'
	}

	if n%2 == 0 {
		res[0] = 'b'
	}
	return string(res)
}

func main() {
	fmt.Println(generateTheString(4))
}
