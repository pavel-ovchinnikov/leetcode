package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xx := x
	y := 0
	for x != 0 {
		y = y*10 + x%10
		x /= 10
	}

	return y == xx
}
