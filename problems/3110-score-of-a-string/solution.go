package scoreofastring

func scoreOfString(s string) int {
	res := 0

	for i := 0; i < len(s)-1; i++ {
		res += abs(int(s[i]) - int(s[i+1]))
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
