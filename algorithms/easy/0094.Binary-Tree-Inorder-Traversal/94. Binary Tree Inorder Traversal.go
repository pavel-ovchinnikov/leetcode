package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type TreeNode = structures.TreeNode

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}

func main() {
	root := structures.Ints2TreeNode([]int{1, 1, 2, 3})
	fmt.Println(inorderTraversal(root))
}
