package main

import "fmt"

func convertToBase7(num int) string {
	res := make([]byte, 0, 100)
	if num == 0 {
		return "0"
	}

	sign := num < 0
	num = abs(num)

	for num != 0 {
		res = append(res, byte(num%7+'0'))
		num /= 7
	}

	if sign {
		res = append(res, '-')
	}

	reverse(res)

	return string(res)
}

func reverse(arr []byte) {
	l, r := 0, len(arr)-1
	for l != r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func main() {
	fmt.Println(convertToBase7(100))
}
