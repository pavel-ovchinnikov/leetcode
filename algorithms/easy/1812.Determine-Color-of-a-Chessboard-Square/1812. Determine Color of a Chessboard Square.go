package main

import "fmt"

func squareIsWhite(coordinates string) bool {
	return (int(coordinates[0]-'a'+1)+int(coordinates[1]-'1'+1))%2 == 1
}

func main() {
	fmt.Println(squareIsWhite("a1"))
}
