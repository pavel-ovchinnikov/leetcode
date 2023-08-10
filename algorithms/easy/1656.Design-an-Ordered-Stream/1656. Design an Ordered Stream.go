package main

import "fmt"

type OrderedStream struct {
	pointer int
	stream  []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{0, make([]string, n)}
}

func (orderedStream *OrderedStream) Insert(idKey int, value string) []string {
	var res []string
	orderedStream.stream[idKey-1] = value
	for orderedStream.pointer < len(orderedStream.stream) && orderedStream.stream[orderedStream.pointer] != "" {
		res = append(res, orderedStream.stream[orderedStream.pointer])
		orderedStream.pointer++
	}

	return res
}

func main() {
	orderedStream := Constructor(5)
	fmt.Println(orderedStream.Insert(3, "ccccc"))
	fmt.Println(orderedStream.Insert(1, "aaaaa"))
	fmt.Println(orderedStream.Insert(2, "bbbbb"))
	fmt.Println(orderedStream.Insert(5, "eeeee"))
	fmt.Println(orderedStream.Insert(4, "ddddd"))
}
