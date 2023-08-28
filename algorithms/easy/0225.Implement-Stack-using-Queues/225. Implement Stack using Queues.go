package main

import "fmt"

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{[]int{}}
}

func (st *MyStack) Push(x int) {
	st.queue = append(st.queue[:0], append([]int{x}, st.queue[0:]...)...) // prepend
}

func (st *MyStack) Pop() int {
	temp := st.queue[0]
	st.queue = st.queue[1:]
	return temp
}

func (st *MyStack) Top() int {
	return st.queue[0]
}

func (st *MyStack) Empty() bool {
	return len(st.queue) == 0
}

func main() {

	obj := Constructor()
	obj.Push(10)
	obj.Push(20)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
