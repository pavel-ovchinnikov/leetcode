package defanginganipaddress

func defangIPaddr(address string) string {
	res := make([]rune, 0, len(address)+6)

	for _, r := range address {
		if r == '.' {
			res = append(res, '[', '.', ']')
			continue
		}

		res = append(res, r)
	}

	return string(res)
}
