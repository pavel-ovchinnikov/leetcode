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

func isBalanced(root *TreeNode) bool {
	result, _ := depth(root)
	return result
}

func depth(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	isLeftBalanced, leftDepth := depth(root.Left)
	isRightBalanced, rightDepth := depth(root.Right)
	diff := abs(leftDepth - rightDepth)
	if isLeftBalanced &&
		isRightBalanced &&
		diff <= 1 {
		return true, 1 + max(leftDepth, rightDepth)
	}
	return false, -1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := structures.Ints2TreeNode([]int{1, 2, 2, 3, 3, 1, 1, 4, 4})
	fmt.Println(isBalanced(root))
}
