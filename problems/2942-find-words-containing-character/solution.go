package findwordscontainingcharacter

func findWordsContaining(words []string, x byte) []int {
	res := make([]int, 0, len(words))

	for i, w := range words {
		for j := 0; j < len(w); j++ {
			if w[j] == x {
				res = append(res, i)
				break
			}
		}
	}

	return res
}
