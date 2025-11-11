package reverseprefixofword

func reversePrefix(word string, ch byte) string {
	buf := []rune(word)
	i := 0
	for i < len(buf) && byte(buf[i]) != ch {
		i++
	}
	if i < len(buf) {
		reverse(buf[:i+1])
	}

	return string(buf)
}

func reverse(buf []rune) {
	l, r := 0, len(buf)-1
	for l < r {
		buf[l], buf[r] = buf[r], buf[l]
		l++
		r--
	}
}
