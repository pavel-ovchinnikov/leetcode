package main

import "fmt"

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		return addBinary(b, a)
	}

	nA := len(a)
	nB := len(b)
	indexB := nB - 1
	result := make([]byte, nA)
	sum := 0

	for i := len(a) - 1; i >= 0; i-- {
		sum += int(a[i]) - int('0')
		if indexB >= 0 {
			sum += int(b[indexB]) - int('0')
			indexB--
		}

		result[i] = '0' + byte(sum%2)
		sum = sum / 2
	}
	if sum == 0 {
		return string(result)
	}

	return "1" + string(result)
}

func main() {
	fmt.Println(addBinary("11", "1"))
}
