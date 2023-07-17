package main

import "fmt"

func interpret(command string) string {
	var res string
	i := 0
	size := len(command)
	for i < size {

		if command[i] == 'G' {
			res += "G"
		}
		if command[i] == '(' {
			if command[i+1] == ')' {
				res += "o"
			} else {
				res += "al"
			}
			i++
		}
		i++
	}
	return res
}

func main() {
	fmt.Println(interpret("G()(al)"))
}
