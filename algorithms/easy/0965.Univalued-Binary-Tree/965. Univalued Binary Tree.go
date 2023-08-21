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

func isUnivalTree(root *TreeNode) bool {
	return helper(root, root.Val)
}

func helper(root *TreeNode, target int) bool {
	if root == nil {
		return true
	}
	if root.Val != target {
		return false
	}

	return helper(root.Left, target) && helper(root.Right, target)
}

func main() {
	root := structures.Ints2TreeNode([]int{2, 2, 2, 5, 2})
	fmt.Println(isUnivalTree(root))
}
