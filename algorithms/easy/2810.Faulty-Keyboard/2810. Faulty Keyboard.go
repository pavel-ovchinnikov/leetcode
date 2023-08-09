package main

import "fmt"

func finalString(s string) string {
	res := make([]byte, 0, len(s))
	for _, c := range s {
		if c == 'i' {
			reverse(res)
		} else {
			res = append(res, byte(c))
		}
	}
	return string(res)
}

func reverse(arr []byte) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func main() {
	// fmt.Println(finalString("string"))
	fmt.Println(finalString("poiinter"))
	// fmt.Println(finalString("ipoiinter"))
}
