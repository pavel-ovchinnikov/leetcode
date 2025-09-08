package distringmatch

func diStringMatch(s string) []int {
	res := make([]int, len(s)+1)
	lo, hi := 0, len(s)

	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			res[i] = lo
			lo++
		} else {
			res[i] = hi
			hi--
		}
	}
	res[len(s)] = lo

	return res
}
