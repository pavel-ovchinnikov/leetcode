package checkifstringisaprefixofarray

func isPrefixString(s string, words []string) bool {
	k := 0
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			if k >= len(s) || s[k] != word[i] {
				return false
			}
			k++
		}
		if k == len(s) {
			return true
		}
	}

	return false
}
