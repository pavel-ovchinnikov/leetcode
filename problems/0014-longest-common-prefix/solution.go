package tasks0014

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	firstWord := strs[0]
	words := strs[1:]
	k := 0

	for i := range firstWord {
		for _, word := range words {
			if i >= len(word) {
				return firstWord[:k]
			}
			if word[i] != firstWord[i] {
				return firstWord[:k]
			}
		}
		k++
	}

	return firstWord[:k]
}
