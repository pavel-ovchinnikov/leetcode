package findtheindexofthefirstoccurrenceinastring

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 {
		return 0
	}
	l1, l2 := len(haystack), len(needle)
	for i := 0; i <= l1-l2; i++ {
		if haystack[i:i+l2] == needle {
			return i
		}
	}

	return -1
}
