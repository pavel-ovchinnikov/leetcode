package main

import (
	"fmt"
	"strconv"
)

func reverseBits(num uint32) (res uint32) {
	n := 31
	for n >= 0 {
		val := num & 1
		res += (val << n)
		num >>= 1
		n--
	}

	return res
}

func main() {
	i, _ := strconv.ParseUint("00000010100101000001111010011100", 2, 10)
	fmt.Println(reverseBits(uint32(i)))
}
