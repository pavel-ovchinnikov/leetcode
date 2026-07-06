package task0013

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}

	hash := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	last := 0
	for _, ch := range s {
		d := hash[ch]
		sum += d
		if last < d {
			sum -= 2 * last
		}

		last = d
	}
	return sum
}
