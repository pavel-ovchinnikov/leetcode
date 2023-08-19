package main

import "fmt"

func numUniqueEmails(emails []string) int {
	hash := make(map[string]struct{})

	address := make([]byte, 0, 128)
	shouldSkip, dotAfterAt, plusAfterAt := false, false, false

	for _, email := range emails {
		address = address[:0]
		shouldSkip, dotAfterAt, plusAfterAt = false, false, false

		for i := 0; i < len(email); i++ {
			char := email[i]

			if char == '.' && !dotAfterAt {
				continue
			}

			if char == '+' && !plusAfterAt {
				shouldSkip = true
			}

			if char == '@' {
				shouldSkip = false
				dotAfterAt = true
				plusAfterAt = true
			}

			if !shouldSkip {
				address = append(address, char)
			}
		}

		hash[string(address)] = struct{}{}
	}

	return len(hash)
}

func main() {
	fmt.Println(numUniqueEmails([]string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com"}))
}
