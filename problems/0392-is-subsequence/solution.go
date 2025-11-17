package issubsequence

func isSubsequence(s string, t string) bool {
	l := 0
	for i := 0; i < len(t) && l < len(s); i++ {
		if s[l] == t[i] {
			l++
		}
	}

	return l == len(s)
}
