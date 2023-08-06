package main

import "fmt"

func destCity(paths [][]string) string {
	hash := map[string]string{}

	for _, val := range paths {
		hash[val[0]] = val[1]
	}

	curCity := paths[0][0]

	for {
		city, ok := hash[curCity]
		if !ok {
			return curCity
		}
		curCity = city
	}
}

func main() {
	fmt.Println(destCity([][]string{
		{"London", "New York"},
		{"New York", "Lima"},
		{"Lima", "Sao Paulo"},
	}))
}
