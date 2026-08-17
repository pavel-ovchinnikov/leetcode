package task0345

func reverseVowels(s string) string {
	runes := []rune(s)
	l, r := 0, len(runes)-1

	for l < r {
		for l < r && !isVowel(runes[l]) {
			l++
		}

		for l < r && !isVowel(runes[r]) {
			r--
		}

		if l >= r {
			break
		}

		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}

	return string(runes)
}

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}
