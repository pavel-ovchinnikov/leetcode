package mergestringsalternately

func mergeAlternately(word1 string, word2 string) string {

	res := make([]byte, 0, len(word1)+len(word2))
	i := 0
	for i < len(word1) && i < len(word2) {
		res = append(res, word1[i])
		res = append(res, word2[i])
		i++
	}
	if i < len(word1) {
		res = append(res, word1[i:]...)
	}
	if i < len(word2) {
		res = append(res, word2[i:]...)
	}

	return string(res)
}
