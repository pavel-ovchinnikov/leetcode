package main

import "fmt"

func romanToInt(s string) int {
	result := 0
	tmp := 0
	hashmap := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := len(s) - 1; i >= 0; i-- {
		val := hashmap[s[i]]
		if val < tmp {
			result -= val
		} else {
			result += val
		}
		tmp = val
	}

	return result
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
