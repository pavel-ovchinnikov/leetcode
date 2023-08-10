package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := []int{root.Val}
	if isLeaf(root) {
		return res
	}
	for _, node := range root.Children {
		res = append(res, preorder(node)...)
	}
	return res
}

func isLeaf(node *Node) bool {
	for _, node := range node.Children {
		if node != nil {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(preorder(&Node{Val: 10, Children: []*Node{}}))
}
