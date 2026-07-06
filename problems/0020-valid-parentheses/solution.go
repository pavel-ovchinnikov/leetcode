package tasks0020

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	stack := make([]rune, 0, len(s))
	for _, ch := range s {
		switch ch {
		case '[', '{', '(':
			stack = append(stack, ch)
		case ']', '}', ')':
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if (ch == ']' && top != '[') ||
				(ch == '}' && top != '{') ||
				(ch == ')' && top != '(') {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
