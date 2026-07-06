package task0009

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	rem := 0
	rev := x

	for x != 0 {
		d := x % 10
		rem = rem*10 + d
		x = x / 10
	}

	return rem == rev
}
