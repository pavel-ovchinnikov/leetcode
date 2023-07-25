package main

import "fmt"

func countAsterisks(s string) int {
	bar := 0
	res := 0

	for _, c := range s {
		if c == '|' {
			bar++
		} else if c == '*' {
			if bar%2 == 0 {
				res++
			}
		}
	}

	return res
}

func main() {
	fmt.Println(countAsterisks("l|*e*et|c**o|*de|"))
}
