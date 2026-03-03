package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	p := []byte(strs[0])

	for i := 1; i < len(strs); i++ {
		j := 0
		for ; j < len(strs[i]) && j < len(p); j++ {
			if p[j] == strs[i][j] {
				continue
			}

			break
		}
		p = p[:j]
	}

	return string(p)
}
