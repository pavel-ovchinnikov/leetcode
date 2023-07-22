package main

import "fmt"

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	result := 0
	col := 0
	switch ruleKey {
	case "type":
		col = 0
	case "color":
		col = 1
	case "name":
		col = 2
	}

	for _, v := range items {
		if v[col] == ruleValue {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(countMatches([][]string{
		{"phone", "blue", "pixel"},
		{"computer", "silver", "lenovo"},
		{"phone", "gold", "iphone"}},
		"color",
		"silver"))
}
