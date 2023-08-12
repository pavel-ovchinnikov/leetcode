package main

import "github.com/pavel-ovchinnikov/LeetCode/structures"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = structures.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func main() {
	root := structures.Ints2TreeNode([]int{4, 2, 7, 1, 3, 6, 9})
	invertTree(root)
}
