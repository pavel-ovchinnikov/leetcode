package main

import "fmt"

func numberOfBeams(bank []string) int {
	if len(bank) <= 1 {
		return 0
	}

	res := 0
	lastLasers := 0

	for i := 0; i < len(bank); i++ {
		count := countLaser(bank[i])
		if count == 0 {
			continue
		}

		if lastLasers == 0 {
			lastLasers = count
			continue
		}

		res += count * lastLasers
		lastLasers = count
	}

	return res
}

func countLaser(s string) (res int) {
	for _, c := range s {
		if c == '1' {
			res++
		}
	}
	return
}

func main() {
	fmt.Println(numberOfBeams([]string{"011001", "000000", "010100", "001000"}))
	fmt.Println(numberOfBeams([]string{"000", "111", "000"}))
}
