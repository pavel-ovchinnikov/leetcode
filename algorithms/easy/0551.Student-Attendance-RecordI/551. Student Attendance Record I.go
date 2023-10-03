package main

import "fmt"

// func checkRecord(s string) bool {
// 	return strings.Count(s, "A") < 2 && strings.Index(s, "LLL") == -1
// }

func checkRecord(s string) bool {
	absent := 0
	late := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			absent++
			if absent == 2 {
				return false
			}
		}

		if s[i] == 'L' {
			late++
			if late == 3 {
				return false
			}
		} else {
			late = 0
		}
	}

	return true
}

func main() {
	fmt.Println(checkRecord("PPALLP"))
}
