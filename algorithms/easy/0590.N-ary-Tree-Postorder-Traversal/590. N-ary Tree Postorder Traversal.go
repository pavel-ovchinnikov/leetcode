package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	if len(root.Children) == 0 {
		return []int{root.Val}
	}

	for _, node := range root.Children {
		res = append(res, postorder(node)...)
	}
	res = append(res, root.Val)

	return res
}

func main() {
	root := &Node{
		Val: 10,
		Children: []*Node{
			{Val: 5, Children: []*Node{}},
			{Val: 6, Children: []*Node{}},
			{Val: 1, Children: []*Node{}},
		},
	}
	fmt.Println(postorder(root))
}
