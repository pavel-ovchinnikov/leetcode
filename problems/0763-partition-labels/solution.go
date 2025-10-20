package partitionlabels

func partitionLabels(s string) []int {
	n := len(s)
	res := make([]int, 0, n)
	stopLetter := make([]int, 26)

	for i := 0; i < n; i++ {
		stopLetter[s[i]-'a'] = max(i, stopLetter[s[i]-'a'])
	}

	i := 0
	for i < n {
		l := i
		r := stopLetter[s[i]-'a']
		for j := i; j < r && j < n; j++ {
			r = max(r, stopLetter[s[j]-'a'])
		}
		res = append(res, r-l+1)
		i = r + 1
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
