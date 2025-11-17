package reverseonlyletters

func reverseOnlyLetters(s string) string {
	l, r := 0, len(s)-1
	buf := []byte(s)

	for l < r {
		if !isLetters(buf[l]) {
			l++
			continue
		}
		if !isLetters(buf[r]) {
			r--
			continue
		}
		buf[l], buf[r] = buf[r], buf[l]
		l++
		r--
	}

	return string(buf)
}

func isLetters(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
