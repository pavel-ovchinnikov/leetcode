package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type TreeNode = structures.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	root := structures.Ints2TreeNode([]int{3, 9, 20, 0, 0, 15, 7})
	fmt.Println(maxDepth(root))
}
