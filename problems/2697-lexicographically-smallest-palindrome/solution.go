package lexicographicallysmallestpalindrome

func makeSmallestPalindrome(s string) string {
	p1, p2 := 0, len(s)-1
	buf := []byte(s)

	for p1 < p2 {
		if buf[p1] != buf[p2] {
			if buf[p1] < buf[p2] {
				buf[p2] = buf[p1]
			} else {
				buf[p1] = buf[p2]
			}
		}
		p1++
		p2--
	}

	return string(buf)
}
