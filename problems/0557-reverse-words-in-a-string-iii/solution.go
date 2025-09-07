package reversewordsinastringiii

func reverseWords(s string) string {
	buf := []byte(s)

	l, r := 0, 1
	for l != r && l < len(buf) && r < len(buf) {
		for l < len(s) && s[l] == ' ' {
			l++
		}
		r = l
		for r < len(s) && s[r] != ' ' {
			r++
		}
		reverse(buf[l:r])
		l = r + 1
	}

	return string(buf)
}

func reverse(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
