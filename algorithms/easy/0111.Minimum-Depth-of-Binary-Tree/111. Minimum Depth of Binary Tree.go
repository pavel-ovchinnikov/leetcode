package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = structures.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if left == 0 || right == 0 {
		return left + right + 1
	}
	return 1 + min(left, right)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	root := structures.Ints2TreeNode([]int{3, 9, 20, 0, 0, 15, 7})
	fmt.Println(minDepth(root))
}
