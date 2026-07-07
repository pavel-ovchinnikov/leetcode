package tasks0028

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		f := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				f = false
				break
			}
		}
		if f {
			return i
		}
	}

	return -1
}
