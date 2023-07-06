package main

import "fmt"

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(i int) int {
	return 0
}

func guessNumber(n int) int {
	l, r, mid := 0, n, 0
	for l <= r {
		mid = l + (r-l)/2
		num := guess(mid)
		if num == 0 {
			break
		}
		if num == 1 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return mid
}

func main() {
	fmt.Println(guessNumber(10))
}
