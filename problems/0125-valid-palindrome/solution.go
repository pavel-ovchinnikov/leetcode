package validpalindrome

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if !isAlphaNum(s[l]) {
			l++
			continue
		}
		if !isAlphaNum(s[r]) {
			r--
			continue
		}
		if toLower(s[l]) != toLower(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNum(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}
