package reversevowelsofastring

func reverseVowels(s string) string {
	l, r := 0, len(s)-1
	res := []byte(s)

	for l < r {
		if !isVowels(res[l]) {
			l++
			continue
		}

		if !isVowels(res[r]) {
			r--
			continue
		}

		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return string(res)
}

func isVowels(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	case 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}
